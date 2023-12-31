package main

import (
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
	case "darwin":
		fmt.Println("Starting front end service on port 80")
		err := http.ListenAndServe(":80", nil)
		if err != nil {
			log.Panic(err)
		}
	}

}

func render(w http.ResponseWriter, t string) {

	partials := []string{
		"./cmd/web/templates/base.layout.gohtml",
		"./cmd/web/templates/header.partial.gohtml",
		"./cmd/web/templates/footer.partial.gohtml",
	}

	var templateSlice []string
	templateSlice = append(templateSlice, fmt.Sprintf("./cmd/web/templates/%s", t))

	for _, x := range partials {
		templateSlice = append(templateSlice, x)
	}

	tmpl, err := template.ParseFiles(templateSlice...)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
