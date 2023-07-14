package utils_test

import (
	"og-post-space-invaders/utils"
	"os"
	"testing"
)

func TestScreenshotMainElement(t *testing.T) {
	// Create a temporary file for testing
	file, err := os.CreateTemp("", "testfile.*.html")
	if err != nil {
		t.Fatalf("Failed to create temporary file: %v", err)
	}
	defer os.Remove(file.Name()) // Clean up the temporary file

	// Write some HTML content to the file
	_, err = file.WriteString(`<html><body><div id="main">Hello, world!</div></body></html>`)
	if err != nil {
		t.Fatalf("Failed to write HTML content to the file: %v", err)
	}

	imageBytes, err := utils.ScreenshotMainElement(file)
	if err != nil {
		t.Fatalf("Failed to capture screenshot: %v", err)
	}

	if len(imageBytes) == 0 {
		t.Error("Expected non-empty screenshot image, got empty")
	}
}
