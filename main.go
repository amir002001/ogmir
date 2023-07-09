package main

import (
	"image"
	"os"

	"github.com/charmbracelet/log"
	"github.com/fogleman/gg"

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

	img, _, err := image.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}

	dc := gg.NewContextForImage(img)
	dc.DrawCircle(500, 500, 400)
	dc.SetRGB(0, 0, 0)
	dc.Fill()
	err = dc.SavePNG("out.png")
	if err != nil {
		log.Fatalf("error saving image: %v", err)
	}
}
