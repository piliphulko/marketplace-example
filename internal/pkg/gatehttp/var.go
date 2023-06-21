package gatehttp

import (
	"html/template"
	"io/fs"
	"path/filepath"
	"strings"
)

var TempHTML = template.New("html").Funcs(template.FuncMap{
	"addFloatFloat": func(a, b float64) float64 {
		return a + b
	},
	"mulFloatInt": func(a float64, b int) float64 {
		return float64(a * float64(b))
	},
})

func FillTempHTMLfromDir(pathDir string) error {
	err := filepath.WalkDir(pathDir, func(path string, d fs.DirEntry, err error) error {
		if strings.Contains(path, ".html") {
			_, err := TempHTML.ParseFiles(path)
			if err != nil {
				return err
			}
		}
		return err
	})
	if err != nil {
		return err
	}
	return nil
}
