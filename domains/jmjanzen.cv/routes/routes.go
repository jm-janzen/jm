package routes

import (
	"os"

	"jm/domains/jmjanzen.cv/handlers"

	"github.com/gin-gonic/gin"
)

func Launch() {
	router := gin.Default()

	router.SetTrustedProxies([]string{os.Getenv("TRUSTED_PROXY")})
	router.LoadHTMLGlob("./domains/jmjanzen.cv/templates/*")
	router.Static("/static", "./domains/jmjanzen.cv/static")
	router.StaticFile("/favicon.ico", "./domains/jmjanzen.cv/static/img/favicon.ico")

	router.GET("/", handlers.Default)

	router.Run(":" + os.Getenv("CV_PORT"))
}
