package handlers

import (
	"fmt"
	"math/rand"
	"og-post-space-invaders/types"
	"og-post-space-invaders/utils"
	"os"
	"path/filepath"
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
