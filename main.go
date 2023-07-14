package main

import (
	"fmt"
	"og-post-space-invaders/handlers"
	"os"
	"time"

	"github.com/charmbracelet/log"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/storage/sqlite3"
)

const (
	GIGA uint = 1024 * 1024 * 1024
	Dev       = "DEVELOPMENT"
	Prod      = "PRODUCTION"
)

var Environment = os.Getenv("ENVIRONMENT")

func main() {
	app := fiber.New()

	cacheStore := sqlite3.New(sqlite3.Config{
		Table: "cache",
	})
	limiterStore := sqlite3.New(sqlite3.Config{
		Table: "ratelimit",
	})

	app.Use(cors.New(cors.Config{AllowOrigins: "https://amir.day, *.amir002001.pages.dev"}))

	app.Use(limiter.New(limiter.Config{
		Max:        10,
		Expiration: 1 * time.Minute,
		Storage:    limiterStore,
	}))

	app.Use(cache.New(cache.Config{
		Expiration:   30 * time.Hour * 24 * 30,
		CacheControl: true,
		KeyGenerator: func(c *fiber.Ctx) string {
			return fmt.Sprintf("%s👾%s", c.Query("title"), c.Query("image"))
		},
		Storage:  cacheStore,
		MaxBytes: 3 * GIGA * 1,
	}))

	api := app.Group("api")
	v1 := api.Group("v1")

	v1.Get("/opengraph", handlers.OpenGraphHandler)

	log.Fatal(app.Listen(":8080"))
}
