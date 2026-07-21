package tmpl

import (
	_ "embed"
	"html/template"
	"io"
)

type ColourPage struct {
	Colours []string
}

//go:embed colours.html.tmpl
var colourPageTemplate string

func (c ColourPage) Render(w io.Writer) error {
	tmpl, err := template.New("colours").Parse(colourPageTemplate)
	if err != nil {
		return err
	}

	err = tmpl.Execute(w, c)
	return err
}
