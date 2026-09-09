# Spec health report

Last evaluated: 2026-09-09
Trigger: RHOAI wheel-build and retired-pipeline alignment
Layout: software (`.ai/spec/`)

## Current state

- The `lsc/` library, its pipelines, and its integration test have been removed. OCP product documentation is served by OKP via the RHOKP sidecar.
- The active Konflux pipelines build only `lightspeed-rag-tool` from `byok/Containerfile.tool`.
- CPU dependencies are represented by split lockfiles: RHOAI wheels, PyPI source distributions, PEP 517 build dependencies, and hermetic bootstrap tools.
- The embedding model is committed as chunked archives under `embeddings_model/` and reassembled in the container build; it is not a Cachi2 generic artifact.

## Follow-up

Regenerate the Konflux dependency lockfiles after changing `pyproject.toml` or `requirements.overrides.txt`. The generation command must complete `pybuild-deps` successfully; a failed build-dependency resolution is not safe to commit.
