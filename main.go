package main

import (
	"og-post-space-invaders/handlers"

	"github.com/charmbracelet/log"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()

	app.Use(cors.New(cors.Config{AllowOrigins: "https://amir.day, *.amir002001.pages.dev"}))

	api := app.Group("api")
	v1 := api.Group("v1")

	v1.Get("/opengraph", handlers.OpenGraphHandler)

	log.Fatal(app.Listen(":8080"))
}
