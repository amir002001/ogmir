package utils

import (
	"og-post-space-invaders/types"
	"os"
	"text/template"

	"github.com/charmbracelet/log"
)

const tmplFile = "index.template.html"

func GenerateTemplate(file *os.File, placeholders types.TemplatePlacehoders) error {
	tmpl, err := template.New(tmplFile).ParseFiles(tmplFile)
	if err != nil {
		return err
	}

	err = tmpl.Execute(file, placeholders)
	if err != nil {
		return err
	}

	log.Infof("template executed for %s", file.Name())
	return nil
}
