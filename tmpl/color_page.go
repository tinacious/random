package tmpl

import (
	"html/template"
	"io"
)

type ColourPage struct {
	Colours []string
}

func (c ColourPage) Render(w io.Writer) error {
	tmplFile := "tmpl/colours.html.tmpl"
	tmpl := template.Must(template.ParseFiles(tmplFile))
	err := tmpl.Execute(w, c)
	return err
}
