package render

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"time"

	"invoice-api/env"
	"invoice-api/invoice"
)

var chrome_executable string

func init() {
	chrome_executable = env.GetDefault("CHROME_EXECUTABLE", "google-chrome")
}

func PDF(invoice *invoice.Invoice) ([]byte, error) {
	// Render Invoice HTML to Temp File
	htmlFile, err := os.CreateTemp(os.TempDir(), "invoice.*.html")
	if err != nil {
		return nil, err
	}
	defer os.Remove(htmlFile.Name())

	html, err := renderInvoiceHTML(invoice)
	if err != nil {
		return nil, err
	}

	_, err = htmlFile.Write(html)
	if err != nil {
		return nil, err
	}

	err = htmlFile.Sync()
	if err != nil {
		return nil, err
	}

	htmlFile.Close()

	// Invoke headless chromium to print to PDF file
	pdfFile, err := os.CreateTemp(os.TempDir(), "invoice.*.pdf")
	if err != nil {
		return nil, err
	}
	defer os.Remove(pdfFile.Name())
	pdfFile.Close()

	cmdc, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	cmd := exec.CommandContext(
		cmdc,
		chrome_executable,
		"--no-sandbox",
		"--disable-gpu",
		"--timeout=5000",
		"--headless",
		"--no-margins",
		"--run-all-compositor-stages-before-draw",
		"--print-to-pdf-no-header",
		fmt.Sprintf("--print-to-pdf=%s", pdfFile.Name()),
		fmt.Sprintf("file://%s", htmlFile.Name()),
	)

	err = cmd.Run()
	if err != nil {
		return nil, err
	}

	// Read PDF into memory
	pdfData, err := os.ReadFile(pdfFile.Name())
	if err != nil {
		return nil, err
	}

	return pdfData, nil
}
