package handlers

import (
	"html/template"
	"jm/internal/word"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yosssi/ace"
)

type pageData struct {
	Title string
	Slug  string
}

const templatesPath = "templates/bodies/"

func fetchTemplate(slug string) (*template.Template, error) {
	tpl, err := ace.Load(templatesPath+slug, "", nil)
	if err != nil {

		// If 404 fails for some reason, just quit
		if tpl, err = ace.Load(templatesPath+"404", "", nil); err != nil {
			log.Println("Error:", err.Error())
			return nil, err
		}
	}

	return tpl, err
}

func RenderAce(c *gin.Context, slug string) {

	tpl, err := fetchTemplate(slug)
	if err != nil {
		tpl, err = fetchTemplate("404")
		if err != nil {
			log.Println("Error:", err.Error())
			return
		}
		return
	}

	data := pageData{
		Title: "JM Janzen's " + word.Capitalise(slug),
		Slug: slug,
	}

	if err := tpl.Execute(c.Writer, data); err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
		log.Println("Error:", err.Error())
		return
	}
}
