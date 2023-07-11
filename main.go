package main

import (
	"html/template"
	"os"

	"github.com/charmbracelet/log"
)

type TemplatePlacehoders struct {
	Www      string
	ImageUrl string
	Score    string
	Title    string
}

func main() {
	tmplFile := "index.template.html"
	tmpl, err := template.New(tmplFile).ParseFiles(tmplFile)
	if err != nil {
		log.Fatalf("nyooo \n%v", err)
	}

	f, err := os.CreateTemp("", "output-*.html")
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("file created : %s", f.Name())
	//	defer os.Remove(f.Name()) // clean up

	placeholders := TemplatePlacehoders{
		Www:      "/Users/amirhosseinazizafshari/dev/og-post-space-invaders/www",
		ImageUrl: "https://res.cloudinary.com/df3h8ffly/image/upload/v1684094745/portfolio/gallery_wp18zu.webp",
		Score:    "12",
		Title:    "FoodBox",
	}
	err = tmpl.Execute(f, placeholders)
	if err != nil {
		log.Fatalf("nyooo \n%v", err)
	}
}
