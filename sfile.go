package main

import (
	"bytes"
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"strings"

	"github.com/go-martini/martini"
)

func main() {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("can not know current path")
		return
	}

	app := martini.Classic()

	app.Get("/", func() string {
		fileInfos, err := ioutil.ReadDir(dir)
		if err != nil {
			return err.Error()
		}

		fileNames := make([]string, 0, len(fileInfos))
		for _, fileInfo := range fileInfos {
			if !fileInfo.IsDir() && !strings.HasPrefix(fileInfo.Name(), ".") {
				fileNames = append(fileNames, fileInfo.Name())
			}
		}

		b := bytes.NewBufferString("")
		tmpl := template.New("index")
		tmpl, err = tmpl.Parse("{{range .fileNames}}<a href={{.}}>{{.}}<a/><br/>{{end}}")
		if err != nil {
			return err.Error()
		}

		tmpl.Execute(b, map[string]interface{}{"fileNames": fileNames})

		return b.String()
	})

	app.Use(martini.Static(dir))
	app.Run()
}
