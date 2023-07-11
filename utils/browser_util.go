package utils

import (
	"fmt"
	"os"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/go-rod/rod/lib/proto"
)

var launcherUrl string

func init() {
	launcherUrl = launcher.New().Bin("chromium-browser").NoSandbox(true).MustLaunch()
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
