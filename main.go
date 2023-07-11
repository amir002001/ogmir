package main

import (
	"fmt"
	"html/template"
	"math/rand"
	"os"
	"time"

	"github.com/charmbracelet/log"
	"github.com/gofiber/fiber/v2"
)

type TemplatePlacehoders struct {
	Www      string
	ImageUrl string
	Score    string
	Title    string
}

const tmplFile = "index.template.html"

func openGraphHandler(ctx *fiber.Ctx) error {
	source := rand.NewSource(time.Now().Unix())
	random := rand.New(source)

	postImageUrl := ctx.Params("image", "")
	postTitle := ctx.Params("title", "")
	score := random.Intn(100)

	placeholders := TemplatePlacehoders{
		Www:      "/Users/amirhosseinazizafshari/dev/og-post-space-invaders/www",
		ImageUrl: postImageUrl,
		Score:    fmt.Sprint(score),
		Title:    postTitle,
	}

	ctx.Type("image/png")

	tmpHtmlFile, err := os.CreateTemp("", "output-*.html")
	if err != nil {
		return err
	}
	defer os.Remove(tmpHtmlFile.Name())

	err = generateTemplate(tmpHtmlFile, placeholders)
	if err != nil {
		return err
	}

	ogImageBytes, err := generateOgImage(tmpHtmlFile)
	if err != nil {
		return err
	}

	if _, err := ctx.Write(ogImageBytes); err != nil {
		return err
	}
	return nil
}

func generateOgImage(*os.File) ([]byte, error) {
	return nil, nil
}

func generateTemplate(file *os.File, placeholders TemplatePlacehoders) error {
	tmpl, err := template.New(tmplFile).ParseFiles(tmplFile)
	if err != nil {
		return err
	}

	err = tmpl.Execute(file, placeholders)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	app := fiber.New()
	api := app.Group("api")
	v1 := api.Group("v1")

	v1.Get("/opengraph", openGraphHandler)

	log.Fatal(app.Listen(":8080"))
}
