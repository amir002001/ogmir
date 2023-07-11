package main

import (
	"og-post-space-invaders/handlers"

	"github.com/charmbracelet/log"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	api := app.Group("api")
	v1 := api.Group("v1")

	v1.Get("/opengraph", handlers.OpenGraphHandler)

	log.Fatal(app.Listen(":8080"))
}
