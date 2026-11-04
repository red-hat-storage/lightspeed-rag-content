package download

import (
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func TestExtractChapters(t *testing.T) {
	html := `
		<a href="/documentation/red_hat_openshift_data_foundation/4.22/html/4.22_release_notes/overview">
		<a href="/documentation/red_hat_openshift_data_foundation/4.22/html/planning_your_deployment/overview">
		<a href="/documentation/red_hat_openshift_data_foundation/4.22/html/4.22_release_notes/overview">
	`

	got := ExtractChapters(html, "4.22")

	want := []string{
		"4.22_release_notes/overview",
		"planning_your_deployment/overview",
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("ExtractChapters() = %v, want %v", got, want)
	}
}

func TestExtractChaptersEmptyHTML(t *testing.T) {
	got := ExtractChapters("", "4.22")

	if len(got) != 0 {
		t.Errorf("Expected no chapters, got %v", got)
	}
}

func TestExtractChaptersDifferentVersion(t *testing.T) {
	html := `<a href="/documentation/red_hat_openshift_data_foundation/4.22/html/4.22_release_notes/overview">`

	got := ExtractChapters(html, "4.21")

	if len(got) != 0 {
		t.Errorf("Expected no chapters for version 4.21, got %v", got)
	}
}

func TestDownloadPDFSuccess(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/pdf")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte("%PDF-1.4 test content"))
		},
	))
	defer server.Close()

	filePath := filepath.Join(t.TempDir(), "sample.pdf")

	downloaded, err := DownloadPDF(
		server.Client(),
		server.URL,
		filePath,
		false,
	)

	if err != nil {
		t.Fatalf("DownloadPDF() returned error: %v", err)
	}
	if !downloaded {
		t.Fatal("expected PDF to be downloaded")
	}

	content, err := os.ReadFile(filePath)
	if err != nil {
		t.Fatalf("could not read downloaded file: %v", err)
	}
	if string(content) != "%PDF-1.4 test content" {
		t.Fatalf("unexpected file content: %q", content)
	}
}

func TestDownloadPDFHTTPError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "Forbidden", http.StatusForbidden)
		},
	))
	defer server.Close()

	filePath := filepath.Join(t.TempDir(), "sample.pdf")

	downloaded, err := DownloadPDF(
		server.Client(),
		server.URL,
		filePath,
		false,
	)

	if err == nil {
		t.Fatal("expected an HTTP error")
	}
	if downloaded {
		t.Fatal("expected download to fail")
	}
}

func TestDownloadPDFSkipsExistingFile(t *testing.T) {
	filePath := filepath.Join(t.TempDir(), "sample.pdf")

	if err := os.WriteFile(filePath, []byte("existing PDF"), 0600); err != nil {
		t.Fatal(err)
	}

	server := httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			t.Error("server should not be called for an existing file")
		},
	))
	defer server.Close()

	downloaded, err := DownloadPDF(
		server.Client(),
		server.URL,
		filePath,
		false,
	)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if downloaded {
		t.Fatal("expected existing file to be skipped")
	}
}

func TestDownloadPDFForceReplacesExistingFile(t *testing.T) {
	filePath := filepath.Join(t.TempDir(), "sample.pdf")

	if err := os.WriteFile(filePath, []byte("old content"), 0600); err != nil {
		t.Fatal(err)
	}

	server := httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/pdf")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte("%PDF-1.4 new content"))
		},
	))
	defer server.Close()

	downloaded, err := DownloadPDF(
		server.Client(),
		server.URL,
		filePath,
		true,
	)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !downloaded {
		t.Fatal("expected PDF to be downloaded in force mode")
	}

	content, err := os.ReadFile(filePath)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(content, []byte("%PDF-1.4 new content")) {
		t.Fatalf("unexpected file content: %q", content)
	}
}

func TestDownloadPDFRejectsNonPDFContent(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "text/html")
			w.WriteHeader(http.StatusOK)
			_, _ = io.WriteString(w, "<html>Not a PDF</html>")
		},
	))
	defer server.Close()

	filePath := filepath.Join(t.TempDir(), "sample.pdf")

	downloaded, err := DownloadPDF(
		server.Client(),
		server.URL,
		filePath,
		false,
	)

	if err == nil {
		t.Fatal("expected non-PDF response to be rejected")
	}
	if downloaded {
		t.Fatal("expected download to fail")
	}
}
