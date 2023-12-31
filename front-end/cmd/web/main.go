package main

import (
	"embed"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"runtime"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		render(w, "test.page.gohtml")
	})

	os := runtime.GOOS

	switch os {
	case "windows":
		fmt.Println("Starting front end service on port 1514")
		err := http.ListenAndServe(":1514", nil)
		if err != nil {
			log.Panic(err)
		}
	case "linux": // WSL2 Env
		fmt.Println("Starting front end service on port 1514")
		err := http.ListenAndServe(":1514", nil)
		if err != nil {
			log.Panic(err)
		}
	case "darwin":
		fmt.Println("Starting front end service on port 8081")
		err := http.ListenAndServe(":8081", nil)
		if err != nil {
			log.Panic(err)
		}
	}

}

//go:embed templates
var templatesFS embed.FS

func render(w http.ResponseWriter, t string) {

	partials := []string{
		"templates/base.layout.gohtml",
		"templates/header.partial.gohtml",
		"templates/footer.partial.gohtml",
	}

	var templateSlice []string
	templateSlice = append(templateSlice, fmt.Sprintf("./templates/%s", t))

	for _, x := range partials {
		templateSlice = append(templateSlice, x)
	}

	tmpl, err := template.ParseFS(templatesFS, templateSlice...)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
