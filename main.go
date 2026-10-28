package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

const docsBaseURL = "https://docs.redhat.com/en/documentation/red_hat_openshift_data_foundation"

// extractChapters extracts unique chapter identifiers from the TOC HTML.
func extractChapters(html string, version string) []string {
	pattern := regexp.MustCompile(
		`/documentation/red_hat_openshift_data_foundation/` +
			regexp.QuoteMeta(version) +
			`/html/([^"?#]+)`,
	)

	matches := pattern.FindAllStringSubmatch(html, -1)

	seen := make(map[string]bool)
	var chapters []string

	for _, match := range matches {
		if len(match) < 2 {
			continue
		}

		chapter := strings.TrimSuffix(match[1], "/")

		if chapter != "" && !seen[chapter] {
			seen[chapter] = true
			chapters = append(chapters, chapter)
		}
	}

	return chapters
}

func downloadPDF(
	client *http.Client,
	pdfURL string,
	filePath string,
	force bool,
) (bool, error) {
	if !force {
		info, err := os.Stat(filePath)
		if err == nil && info.Size() > 0 {
			return false, nil
		}
		if err != nil && !os.IsNotExist(err) {
			return false, fmt.Errorf("checking existing file: %w", err)
		}
	}

	req, err := http.NewRequest(http.MethodGet, pdfURL, nil)
	if err != nil {
		return false, fmt.Errorf("creating request: %w", err)
	}

	resp, err := client.Do(req)
	if err != nil {
		return false, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return false, fmt.Errorf("HTTP error: %s", resp.Status)
	}

	if !strings.Contains(
		strings.ToLower(resp.Header.Get("Content-Type")),
		"application/pdf",
	) {
		return false, fmt.Errorf("response is not identified as a PDF")
	}

	if err := os.MkdirAll(filepath.Dir(filePath), 0755); err != nil {
		return false, fmt.Errorf("creating output directory: %w", err)
	}

	tempFile := filePath + ".part"
	out, err := os.Create(tempFile)
	if err != nil {
		return false, fmt.Errorf("creating temporary file: %w", err)
	}

	_, copyErr := io.Copy(out, resp.Body)
	closeErr := out.Close()

	if copyErr != nil || closeErr != nil {
		os.Remove(tempFile)
		return false, fmt.Errorf(
			"saving PDF: copy error: %v, close error: %v",
			copyErr,
			closeErr,
		)
	}

	info, err := os.Stat(tempFile)
	if err != nil || info.Size() == 0 {
		os.Remove(tempFile)
		return false, fmt.Errorf("downloaded PDF is empty or unreadable")
	}

	if err := os.Rename(tempFile, filePath); err != nil {
		os.Remove(tempFile)
		return false, fmt.Errorf("finalizing PDF: %w", err)
	}

	return true, nil
}

func main() {
	version := flag.String("version", "", "ODF documentation version")
	force := flag.Bool("force", false, "Redownload existing files")
	flag.Parse()

	if *version == "" {
		fmt.Println("Usage: go run . --version 4.22 [--force]")
		os.Exit(1)
	}

	tocURL := docsBaseURL + "/" + *version
	fmt.Println("Fetching documentation TOC:", tocURL)

	client := &http.Client{
		Timeout: 2 * time.Minute,
	}

	resp, err := client.Get(tocURL)
	if err != nil {
		fmt.Println("Error fetching TOC:", err)
		os.Exit(1)
	}

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Unexpected TOC HTTP status:", resp.Status)
		resp.Body.Close()
		os.Exit(1)
	}

	body, err := io.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Println("Error reading TOC:", err)
		os.Exit(1)
	}

	html := string(body)
	fmt.Println("TOC fetched successfully. HTML size:", len(html), "bytes")

	chapters := extractChapters(html, *version)

	if len(chapters) == 0 {
		fmt.Println("No chapters found. Check the TOC URL or extraction pattern.")
		os.Exit(1)
	}

	outputDir := filepath.Join("downloads", *version)
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		fmt.Println("Error creating output directory:", err)
		os.Exit(1)
	}

	fmt.Println("Unique chapters found:", len(chapters))
	fmt.Println("Output directory:", outputDir)

	downloaded := 0
	skipped := 0
	failed := 0

	for i, chapter := range chapters {
		filename := chapter + ".pdf"
		filePath := filepath.Join(outputDir, filename)

		fmt.Printf("\n[%d/%d] %s\n", i+1, len(chapters), filename)

		pdfURL := fmt.Sprintf(
			"%s/%s/pdf/%s/%s.pdf",
			docsBaseURL,
			*version,
			chapter,
			chapter,
		)

		wasDownloaded, err := downloadPDF(
			client,
			pdfURL,
			filePath,
			*force,
		)

		if err != nil {
			fmt.Println("Failed:", err)
			failed++
			continue
		}

		if !wasDownloaded {
			fmt.Println("Skipped: file already exists.")
			skipped++
			continue
		}

		info, err := os.Stat(filePath)
		if err != nil {
			fmt.Println("Downloaded file exists but cannot be checked:", err)
			failed++
			continue
		}

		fmt.Printf("Downloaded successfully (%d bytes)\n", info.Size())
		downloaded++
	}

	fmt.Println("\n========== Download Summary ==========")
	fmt.Println("Version:", *version)
	fmt.Println("Downloaded:", downloaded)
	fmt.Println("Skipped:", skipped)
	fmt.Println("Failed:", failed)
	fmt.Println("Saved in:", outputDir)

	if failed > 0 {
		fmt.Println("Some files failed. Review the messages above before retrying.")
		os.Exit(1)
	}
}
