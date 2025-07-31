package handlers

import (
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func Default(c *gin.Context) {
	contents, err := os.ReadFile(
        "./domains/jmjanzen.cv/static/html/jmjanzen-resume.html",
    )
	if err != nil {
		log.Println("Error:", err)
	}

	data := struct {
		Resume template.HTML
	}{
        template.HTML(contents),
	}
	c.HTML(
		http.StatusOK,
		"home.tmpl",
		data,
	)
}
