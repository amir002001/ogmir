package utils

import (
	"fmt"
	"os"

	"github.com/charmbracelet/log"
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/go-rod/rod/lib/proto"
)

var launcherUrl string

func init() {
	browser := launcher.NewBrowser()
	path, err := browser.Get()
	if err != nil {
		log.Fatal("could not get browser binary")
	}
	log.Infof("browser path: %s", path)

	launcherUrl = launcher.New().Bin(path).MustLaunch()
}

func ScreenshotMainElement(file *os.File) ([]byte, error) {
	page := rod.New().
		ControlURL(launcherUrl).
		MustConnect().
		MustPage(fmt.Sprintf("file://%s", file.Name()))
	err := page.WaitLoad()
	el := page.MustElement("#main")
	if err != nil {
		return nil, err
	}
	return el.Screenshot(proto.PageCaptureScreenshotFormatPng, 100)
}
