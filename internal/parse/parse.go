package parse

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

const docsBaseURL = "https://docs.redhat.com/en/documentation/red_hat_openshift_data_foundation"

// ParsePDFBytes converts PDF bytes into markdown text using the liteparse CLI.
// It pipes the bytes via stdin to avoid temporary files.
func ParsePDFBytes(pdfBytes []byte) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 45*time.Second)
	defer cancel()

	cmd := exec.CommandContext(ctx, "lit", "parse", "--format", "markdown", "--no-ocr", "-")

	var stdout, stderr bytes.Buffer
	cmd.Stdin = bytes.NewReader(pdfBytes)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("liteparse error: %w, stderr: %s", err, stderr.String())
	}

	return stdout.String(), nil
}

// metadata returns YAML front matter for a parsed document.
func metadata(name, version string) string {
	return fmt.Sprintf(`---
source_file: %s.pdf
version: "%s"
chapter: %s
source_url: %s/%s/pdf/%s/%s.pdf
---

`, name, version, name, docsBaseURL, version, name, name)
}

// Run parses all PDFs in downloads/<version>/ and writes
// markdown with metadata to parsed/<version>/.
func Run(version string) error {
	inputDir := filepath.Join("downloads", version)
	outputDir := filepath.Join("parsed", version)

	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return fmt.Errorf("creating output dir: %w", err)
	}

	pdfs, err := filepath.Glob(filepath.Join(inputDir, "*.pdf"))
	if err != nil {
		return fmt.Errorf("listing PDFs: %w", err)
	}

	if len(pdfs) == 0 {
		return fmt.Errorf("no PDFs found in %s", inputDir)
	}

	fmt.Printf("Parsing %d PDFs from %s\n\n", len(pdfs), inputDir)

	parsed := 0
	failed := 0

	for i, pdf := range pdfs {
		name := strings.TrimSuffix(filepath.Base(pdf), ".pdf")
		outPath := filepath.Join(outputDir, name+".md")

		fmt.Printf("[%d/%d] %s\n", i+1, len(pdfs), name)

		pdfBytes, err := os.ReadFile(pdf)
		if err != nil {
			fmt.Printf("  Failed to read: %v\n", err)
			failed++
			continue
		}

		text, err := ParsePDFBytes(pdfBytes)
		if err != nil {
			fmt.Printf("  Failed to parse: %v\n", err)
			failed++
			continue
		}

		content := metadata(name, version) + text

		if err := os.WriteFile(outPath, []byte(content), 0644); err != nil {
			fmt.Printf("  Failed to write: %v\n", err)
			failed++
			continue
		}

		fmt.Printf("  → %s (%d bytes)\n", outPath, len(content))
		parsed++
	}

	fmt.Println("\n========== Parse Summary ==========")
	fmt.Println("Version:", version)
	fmt.Println("Parsed:", parsed)
	fmt.Println("Failed:", failed)
	fmt.Println("Output:", outputDir)

	if failed > 0 {
		return fmt.Errorf("%d files failed to parse", failed)
	}

	return nil
}
