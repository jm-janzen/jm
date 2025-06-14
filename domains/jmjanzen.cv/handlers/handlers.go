package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Default(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"resume.html",
		nil,
	)
}
