package utils_test

import (
	"html/template"
	"og-post-space-invaders/types"
	"os"
	"path/filepath"
	"testing"
)

func TestGenerateTemplate(t *testing.T) {
	file, err := os.CreateTemp("", "template.*.html")
	if err != nil {
		t.Fatalf("Failed to create temporary file: %v", err)
	}
	defer os.Remove(file.Name())

	tmpl, err := template.New("index.template.html").ParseFiles("../index.template.html")
	if err != nil {
		t.Fatalf("Failed to create template from file %v", err)
	}

	cwd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get cwd %v", err)
	}

	wwwDir := filepath.Join(cwd, "www")
	wwwAbs, err := filepath.Abs(wwwDir)
	if err != nil {
		t.Fatalf("Failed to get abs path for www folder %v", err)
	}

	placeholders := types.TemplatePlacehoders{
		Www:      wwwAbs,
		ImageUrl: "https://placehold.co/600x400",
		Date:     "September 23 2023",
		Title:    "Amir",
	}

	err = tmpl.Execute(file, placeholders)
	if err != nil {
		t.Fatalf("Failed to generate template: %v", err)
	}
}
