package handlers

import (
	"log"
	"net/http"

	"github.com/yosssi/ace"
)

type Data struct {
	Title string
	Slug string
}

func Default(w http.ResponseWriter, r *http.Request) {
	var requestedPath = r.URL.Path[1:] // Trim leading `/'

	// Default to rendering home template
	if requestedPath == "" {
		requestedPath = "home"
	}

	var pageTitle string = requestedPath

	RenderTpl(w, r, requestedPath, pageTitle)
}

func HandleFile (w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/img/fav/favicon.ico")
}

func RenderTpl(w http.ResponseWriter, r *http.Request, template string, pageTitle string) {
	templatesPath := "templates/bodies/"

	tpl, err := ace.Load(templatesPath+template, "", nil)
	if err != nil {

		log.Println("Error:", err.Error(), "trying 404 instead")
		pageTitle, template = "not found", "404"

		// If 404 fails for some reason, just quit
		if tpl, err = ace.Load(templatesPath+"404", "", nil); err != nil {
			log.Println("Error:", err.Error())
			return
		}
	}

	log.Println("Serving template:", templatesPath+template)

	data := Data{Title: "jm - " + pageTitle, Slug: template}

	if err := tpl.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println("Error:", err.Error())
		return
	}
}

