package handlers

import (
	"fmt"
	"math/rand"
	"og-post-space-invaders/types"
	"og-post-space-invaders/utils"
	"os"
	"time"

	"github.com/charmbracelet/log"
	"github.com/gofiber/fiber/v2"
)

func OpenGraphHandler(ctx *fiber.Ctx) error {
	source := rand.NewSource(time.Now().Unix())
	random := rand.New(source)

	postImageUrl := ctx.Query("image", "")
	postTitle := ctx.Query("title", "")
	score := random.Intn(100)

	placeholders := types.TemplatePlacehoders{
		Www:      "/Users/amirhosseinazizafshari/dev/og-post-space-invaders/www",
		ImageUrl: postImageUrl,
		Score:    fmt.Sprint(score),
		Title:    postTitle,
	}

	log.Infof("generated template placeholders: %v", placeholders)

	tmpFile, err := os.CreateTemp("", "output-*.html")
	if err != nil {
		return err
	}
	defer os.Remove(tmpFile.Name())
	defer tmpFile.Close()

	err = utils.GenerateTemplate(tmpFile, placeholders)
	if err != nil {
		return err
	}

	ogImageBytes, err := utils.ScreenshotMainElement(tmpFile)
	if err != nil {
		return err
	}

	ctx.Type("png")

	if _, err := ctx.Write(ogImageBytes); err != nil {
		return err
	}
	return nil
}
