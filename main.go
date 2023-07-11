package main

import (
	"fmt"
	"html/template"
	"math/rand"
	"os"
	"time"

	"github.com/charmbracelet/log"
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/proto"
	"github.com/gofiber/fiber/v2"
)

type TemplatePlacehoders struct {
	Www      string
	ImageUrl string
	Score    string
	Title    string
}

const tmplFile = "index.template.html"

func main() {
	app := fiber.New()
	api := app.Group("api")
	v1 := api.Group("v1")

	v1.Get("/opengraph", openGraphHandler)

	log.Fatal(app.Listen(":8080"))
}

func openGraphHandler(ctx *fiber.Ctx) error {
	source := rand.NewSource(time.Now().Unix())
	random := rand.New(source)

	postImageUrl := ctx.Query("image", "")
	postTitle := ctx.Query("title", "")
	score := random.Intn(100)

	placeholders := TemplatePlacehoders{
		Www:      "/Users/amirhosseinazizafshari/dev/og-post-space-invaders/www",
		ImageUrl: postImageUrl,
		Score:    fmt.Sprint(score),
		Title:    postTitle,
	}

	log.Infof("generated template placeholders: %v", placeholders)

	tmpHtmlFile, err := os.CreateTemp("", "output-*.html")
	if err != nil {
		return err
	}
	defer os.Remove(tmpHtmlFile.Name())
	defer tmpHtmlFile.Close()

	err = generateTemplate(tmpHtmlFile, placeholders)
	if err != nil {
		return err
	}

	ogImageBytes, err := generateOgImage(tmpHtmlFile)
	if err != nil {
		return err
	}

	ctx.Type("png")

	if _, err := ctx.Write(ogImageBytes); err != nil {
		return err
	}
	return nil
}

func generateOgImage(file *os.File) ([]byte, error) {
	page := rod.New().MustConnect().MustPage(fmt.Sprintf("file://%s", file.Name()))
	err := page.WaitLoad()
	el := page.MustElement("#main")
	if err != nil {
		return nil, err
	}
	return el.Screenshot(proto.PageCaptureScreenshotFormatPng, 100)
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

	log.Infof("template executed for %s", file.Name())
	return nil
}
