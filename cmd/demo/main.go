package main

import (
	"fmt"
	"html/template"
	"os"
)

var templateDemo = `
{{ define "a" }}
Template A includes Template B.
{{block "b" .}} Block {{end}}
Template A ends.
{{ end }}

{{define "b"}}
Template B.
{{end}}
`

func main() {
	var err error

	t := template.New("templateActionDemo")

	t, err = t.Parse(templateDemo)
	if err != nil {
		fmt.Printf("parsing failed: %s\n", err)
	}
	err = t.ExecuteTemplate(os.Stdout, "b", nil)
	if err != nil {
		fmt.Printf("execution failed: %s\n", err)
	}
//	err = t.ExecuteTemplate(os.Stdout, "b", nil)
//	if err != nil {
//		fmt.Printf("execution failed: %s\n", err)
//	}
}
