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

// Render whatever template is present at "/templates/bodies/{resource}.ace",
// else write verbose error to w
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

// Renders given template by name (string)
// FIXME handle errors better once dev settles
func RenderTpl(w http.ResponseWriter, r *http.Request, template string, pageTitle string) {
	templatesPath := "templates/bodies/"

	// Load given template by name
	tpl, err := ace.Load(templatesPath+template, "", nil)
	if err != nil {

		// Invalid resource - hardcode to redirect to 404 page
		log.Println("Error:", err.Error(), "trying 404 instead")
		pageTitle, template = "not found", "404"

		// If this fails for some reason, just quit
		if tpl, err = ace.Load(templatesPath+"404", "", nil); err != nil {
			log.Println("Error:", err.Error())
			return
		}
	}

	// Print IP, URL, requested path; path to template file
	log.Println("Serving template:", templatesPath+template)

	// Load our Data obj
	data := Data{Title: "jm - " + pageTitle, Slug: template}

	// Apply parsed template to w, passing in our Data obj
	if err := tpl.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println("Error:", err.Error())
		return
	}
}

