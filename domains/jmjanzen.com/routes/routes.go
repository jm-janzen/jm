package routes

import (
	"os"

	"jm/domains/jmjanzen.com/handlers"
	"jm/internal/webserver"

	"github.com/gin-gonic/gin"
)

type Data struct {
	Title string
	Slug  string
}

func Launch() {
	webserver.Root("domains/jmjanzen.com")

	router := gin.Default()
	router.Static("/static/", "./static/")

	router.SetTrustedProxies([]string{os.Getenv("TRUSTED_PROXY")})

	router.StaticFile("/favicon.ico", "static/img/fav/favicon.ico")

	router.GET("/", func(c *gin.Context) {
		handlers.RenderAce(c, "home")
	})

	router.GET("/:slug", func(c *gin.Context) {
		handlers.RenderAce(c, c.Param("slug"))
	})

	router.Run(os.Getenv("COM_HOST"))
}
