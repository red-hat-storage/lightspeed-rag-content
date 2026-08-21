"""
Process recipes using the existing Markdown chunking logic.
"""

import json
import logging
import sys
from pathlib import Path
from typing import Any, Callable

from llama_index.core import Settings, SimpleDirectoryReader
from llama_index.core.schema import TextNode
from llama_index.readers.file.flat.base import FlatReader


def recipe_file_metadata_func(file_path: str) -> dict[str, Any]:
    """
    Populate metadata for a recipe page.

    Args:
        file_path: File path string

    Returns:
        Dictionary with metadata
    """
    path_obj = Path(file_path)

    title = ""
    try:
        with open(file_path, "r", encoding="utf-8") as f:
            first_line = f.readline().strip()
            if first_line.startswith("#"):
                title = first_line.lstrip("# ").strip()
            else:
                title = path_obj.stem.replace("_", " ").replace("-", " ").title()
    except Exception:
        title = path_obj.stem.replace("_", " ").replace("-", " ").title()

    RECIPES_ROOT_URL = "https://github.com/red-hat-storage/lightspeed-rag-content/blob/main"
    relative_path = file_path
    if "/recipes/" in file_path:
        relative_path = file_path.split("/recipes/")[1]
    elif "\\recipes\\" in file_path:
        relative_path = file_path.split("\\recipes\\")[1]

    docs_url = f"{RECIPES_ROOT_URL}/odf/dr/recipes/{relative_path}"

    return {
        "docs_url": docs_url,
        "title": title,
        "doc_type": "recipe",
        "source_file": str(path_obj.name),
    }


def process_recipes(
    recipes_dir: Path, output_dir: Path, max_token_limit: int = 380
) -> bool:
    """
    Process recipes and save as chunks.

    Args:
        recipes_dir: Directory containing recipes
        output_dir: Directory to save processed chunks
        max_token_limit: Maximum token limit (used for text splitter settings)

    Returns:
        True if processing was successful
    """
    logger = logging.getLogger(__name__)

    if not recipes_dir.exists():
        logger.error("Recipes directory does not exist: %s", recipes_dir)
        return False

    md_files = list(recipes_dir.rglob("*.md"))
    if not md_files:
        logger.warning("No markdown files found in %s", recipes_dir)
        return True

    logger.info("Found %s recipe files to process", len(md_files))

    try:
        output_dir.mkdir(parents=True, exist_ok=True)

        Settings.chunk_size = max_token_limit
        Settings.chunk_overlap = 20

        logger.info("Loading recipe documents...")
        recipe_documents = SimpleDirectoryReader(
            str(recipes_dir),
            recursive=True,
            required_exts=[".md"],
            file_extractor={".md": FlatReader()},
            file_metadata=recipe_file_metadata_func,
        ).load_data()

        if not recipe_documents:
            logger.warning("No recipe documents loaded")
            return True

        logger.info("Loaded %s recipe documents", len(recipe_documents))

        logger.info("Creating chunks from recipe documents...")
        recipe_nodes = Settings.text_splitter.get_nodes_from_documents(
            recipe_documents
        )

        good_nodes = []
        for node in recipe_nodes:
            if isinstance(node, TextNode) and has_whitespace(node.text):
                good_nodes.append(node)
            else:
                logger.debug("Skipping node without whitespace: %s...", node.text[:50])

        logger.info("Generated %s recipe chunks", len(good_nodes))

        chunk_count = 0
        for i, node in enumerate(good_nodes):
            chunk_data = {
                "id": f"recipe_chunk_{i:04d}",
                "content": node.text,
                "metadata": {
                    **node.metadata,
                    "chunk_index": i,
                    "total_chunks": len(good_nodes),
                    "doc_type": "recipe",
                },
            }

            chunk_filename = f"recipe_chunk_{i:04d}.json"
            chunk_file_path = output_dir / chunk_filename

            with open(chunk_file_path, "w", encoding="utf-8") as f:
                json.dump(chunk_data, f, indent=2, ensure_ascii=False)

            chunk_count += 1

        summary = {
            "total_documents": len(recipe_documents),
            "total_chunks": chunk_count,
            "chunking_params": {
                "chunk_size": Settings.chunk_size,
                "chunk_overlap": Settings.chunk_overlap,
            },
            "processing_type": "markdown_recipes",
        }

        with open(output_dir / "recipes_summary.json", "w", encoding="utf-8") as f:
            json.dump(summary, f, indent=2)

        logger.info(
            "Successfully processed %s recipes into %s chunks",
            len(recipe_documents),
            chunk_count,
        )
        return True

    except Exception as e:
        logger.error("Recipes processing failed: %s", e)
        return False


def has_whitespace(text: str) -> bool:
    """
    Check if text contains whitespace (indicates meaningful content).

    Args:
        text: Text to check

    Returns:
        True if text contains whitespace
    """
    if not text:
        return False

    for char in text:
        if char.isspace():
            return True
    return False


def validate_recipe_chunks(output_dir: Path) -> dict[str, Any]:
    """
    Validate generated recipe chunks.

    Args:
        output_dir: Directory containing chunks

    Returns:
        Validation results
    """
    chunk_files = list(output_dir.glob("recipe_chunk_*.json"))

    if not chunk_files:
        return {"valid": False, "error": "No recipe chunk files found"}

    validation_results = {
        "total_chunks": len(chunk_files),
        "valid_chunks": 0,
        "empty_chunks": 0,
        "invalid_json": 0,
        "missing_metadata": 0,
    }

    for chunk_file in chunk_files:
        try:
            with open(chunk_file, "r", encoding="utf-8") as f:
                chunk_data = json.load(f)

            content = chunk_data.get("content", "")
            metadata = chunk_data.get("metadata", {})

            if not content.strip():
                validation_results["empty_chunks"] += 1
                continue

            required_metadata = ["docs_url", "title", "doc_type"]
            if not all(key in metadata for key in required_metadata):
                validation_results["missing_metadata"] += 1
                continue

            validation_results["valid_chunks"] += 1

        except json.JSONDecodeError:
            validation_results["invalid_json"] += 1
        except Exception:
            continue

    validation_results["valid"] = (
        validation_results["valid_chunks"] > 0
        and validation_results["invalid_json"] == 0
        and validation_results["empty_chunks"] < len(chunk_files) * 0.1
    )

    return validation_results


def get_recipes_stats(output_dir: Path) -> dict[str, Any]:
    """
    Get statistics about processed recipes.

    Args:
        output_dir: Directory containing chunks

    Returns:
        Statistics dictionary
    """
    chunk_files = list(output_dir.glob("recipe_chunk_*.json"))

    if not chunk_files:
        return {"total_chunks": 0}

    source_files = set()
    chunk_sizes = []

    for chunk_file in chunk_files:
        try:
            with open(chunk_file, "r", encoding="utf-8") as f:
                chunk_data = json.load(f)

            metadata = chunk_data.get("metadata", {})
            content = chunk_data.get("content", "")

            source_file = metadata.get("source_file", "unknown")
            source_files.add(source_file)
            chunk_sizes.append(len(content))

        except Exception:
            continue

    stats = {
        "total_chunks": len(chunk_files),
        "source_files": len(source_files),
        "avg_chunk_size": sum(chunk_sizes) / len(chunk_sizes) if chunk_sizes else 0,
        "min_chunk_size": min(chunk_sizes) if chunk_sizes else 0,
        "max_chunk_size": max(chunk_sizes) if chunk_sizes else 0,
    }

    return stats


def main():
    """Command line interface for standalone usage."""
    import argparse

    parser = argparse.ArgumentParser(description="Process recipes for embeddings")
    parser.add_argument(
        "--recipes-dir", "-r", required=True, help="Recipes directory"
    )
    parser.add_argument("--output-dir", "-o", required=True, help="Output directory")
    parser.add_argument(
        "--max-token-limit", "-t", type=int, default=380, help="Max token limit"
    )
    parser.add_argument("--verbose", action="store_true", help="Verbose logging")
    parser.add_argument(
        "--validate", action="store_true", help="Validate chunks after generation"
    )
    parser.add_argument(
        "--stats", action="store_true", help="Show processing statistics"
    )

    args = parser.parse_args()

    level = logging.DEBUG if args.verbose else logging.INFO
    logging.basicConfig(level=level, format="%(asctime)s - %(levelname)s - %(message)s")

    recipes_dir = Path(args.recipes_dir)
    output_dir = Path(args.output_dir)

    from llama_index.core.llms.utils import resolve_llm

    Settings.llm = resolve_llm(None)
    Settings.chunk_size = args.max_token_limit
    Settings.chunk_overlap = 20

    success = process_recipes(
        recipes_dir=recipes_dir,
        output_dir=output_dir,
        max_token_limit=args.max_token_limit,
    )

    if args.validate and success:
        print("\nValidating recipe chunks...")
        validation = validate_recipe_chunks(output_dir)
        print(f"Validation: {'PASSED' if validation['valid'] else 'FAILED'}")
        print(f"  Total chunks: {validation['total_chunks']}")
        print(f"  Valid chunks: {validation['valid_chunks']}")
        print(f"  Empty chunks: {validation['empty_chunks']}")
        print(f"  Invalid JSON: {validation['invalid_json']}")
        print(f"  Missing metadata: {validation['missing_metadata']}")

    if args.stats and success:
        print("\nRecipes Statistics:")
        stats = get_recipes_stats(output_dir)
        print(f"  Total chunks: {stats['total_chunks']}")
        print(f"  Source files: {stats['source_files']}")
        print(f"  Avg chunk size: {stats['avg_chunk_size']:.0f} chars")
        print(
            f"  Size range: {stats['min_chunk_size']}-{stats['max_chunk_size']} chars"
        )

    sys.exit(0 if success else 1)


if __name__ == "__main__":
    main()
