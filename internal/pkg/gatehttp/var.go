package gatehttp

import (
	"html/template"

	"github.com/piliphulko/marketplace-example/internal/pkg/accountaut"
)

var TempHTML = template.New("html").Funcs(template.FuncMap{
	"addFloatFloat": func(a, b float64) float64 {
		return a + b
	},
	"mulFloatInt": func(a float64, b int) float64 {
		return float64(a * float64(b))
	},
})

var (
	ConnServerAA accountaut.ConnAccountAut
)

const (
	grpcAA = iota
)
