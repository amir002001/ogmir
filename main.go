package main

import (
	"image"
	"os"

	"github.com/charmbracelet/log"
	"github.com/fogleman/gg"
	"golang.org/x/image/webp"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	reader, err := os.Open("Game.png")
	if err != nil {
		log.Fatal(err)
	}
	defer reader.Close()

	backgroundImage, _, err := image.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}
	reader2, err := os.Open("image.webp")
	if err != nil {
		log.Fatal(err)
	}
	defer reader2.Close()

	actualImage, err := webp.Decode(reader2)
	if err != nil {
		log.Fatal(err)
	}

	dc := gg.NewContextForImage(backgroundImage)
	dc.DrawImageAnchored(actualImage, 1600, 1800, 0.5, 0.5)

	err = dc.LoadFontFace("font.ttf", 144.0)
	if err != nil {
		log.Fatal(err)
	}

	dc.SetRGB(1, 171/255.0, 186/255.0)
	dc.DrawStringAnchored("Docuchat", 1600, 400, 0.5, 0.5)

	err = dc.SavePNG("out.png")
	if err != nil {
		log.Fatalf("error saving image: %v", err)
	}
}
