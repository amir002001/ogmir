package handlers

import (
	"og-post-space-invaders/types"
	"og-post-space-invaders/utils"
	"os"
	"path/filepath"

	"github.com/charmbracelet/log"
	"github.com/gofiber/fiber/v2"
)

func OpenGraphHandler(ctx *fiber.Ctx) error {
	postImageUrl := ctx.Query("image", "")
	postTitle := ctx.Query("title", "")
	postDate := ctx.Query("date", "")

	dir, err := os.Getwd()
	if err != nil {
		return err
	}

	wwwDir := filepath.Join(dir, "www")
	wwwAbs, err := filepath.Abs(wwwDir)
	if err != nil {
		return err
	}

	placeholders := types.TemplatePlacehoders{
		Www:      wwwAbs,
		ImageUrl: postImageUrl,
		Title:    postTitle,
		Date:     postDate,
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
