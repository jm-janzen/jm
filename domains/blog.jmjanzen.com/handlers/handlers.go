package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleBlog(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		c.Param("slug")+".tmpl",
		nil,
	)
}

func Default(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"index.tmpl",
		nil,
	)
}
