package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yosssi/ace"
)

type Data struct {
	Title string
	Slug string
}

func RenderAce(c *gin.Context, template string) {
	templatesPath := "templates/bodies/"

	tpl, err := ace.Load(templatesPath+template, "", nil)
	if err != nil {

		log.Println("Error:", err.Error(), "trying 404 instead")
		template = "404"

		// If 404 fails for some reason, just quit
		if tpl, err = ace.Load(templatesPath+template, "", nil); err != nil {
			log.Println("Error:", err.Error())
			return
		}
	}

	log.Println("Serving template:", templatesPath+template)

	data := Data{Title: "jm", Slug: template}

	if err := tpl.Execute(c.Writer, data); err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
		log.Println("Error:", err.Error())
		return
	}
}
