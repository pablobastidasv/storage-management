package layouts

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/oxtoacart/bpool"
)

var templates map[string]*template.Template
var bufpool *bpool.BufferPool
var mainTmpl = `{{define "main"}} {{template "base" .}} {{end}}`

func init() {
	bufpool = bpool.NewBufferPool(64)
	log.Println("buffer allocation successful")
}

func Init(layoutPath, templatesPath string) error {
	if templates == nil {
		templates = make(map[string]*template.Template)
	}

	layoutFiles, err := filepath.Glob(layoutPath + "*.html")
	if err != nil {
		return err
	}

	templateFiles, err := filepath.Glob(templatesPath + "*.html")
	if err != nil {
		return err
	}

	mainTemplate := template.New("main")

	mainTemplate, err = mainTemplate.Parse(mainTmpl)
	if err != nil {
		return err
	}

	log.Printf("processing %d layouts.\n", len(layoutFiles))
	log.Printf("processing %d templates.\n", len(templateFiles))

	for _, file := range templateFiles {
		filename := filepath.Base(file)
		files := append(layoutFiles, file)
		templates[filename], err = mainTemplate.Clone()
		if err != nil {
			return err
		}

		templates[filename] = template.Must(templates[filename].ParseFiles(files...))
	}

	log.Println("templates loading successful")

	return nil
}

func Render(wr http.ResponseWriter, fileName string, data interface{}) error {
	tmpl, ok := templates[fileName]
	if !ok {
		err := fmt.Errorf("template %s doesn't exist", fileName)
		return err
	}

	buf := bufpool.Get()
	defer bufpool.Put(buf)

	err := tmpl.Execute(wr, data)
	if err != nil {
		log.Println("Error executing the template: " + err.Error())
		return err
	}

	wr.Header().Set("Content-Type", "text/html; charset=utf-8")
	buf.WriteTo(wr)
	return nil
}
