package routes

import (
	"net/http"
	"os"

	docs "jm/domains/api.jmjanzen.com/docs"
	"jm/domains/api.jmjanzen.com/handlers/interests"
	"jm/domains/api.jmjanzen.com/handlers/me"
	"jm/domains/api.jmjanzen.com/handlers/tils"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Launch() {
	router := gin.Default()
	router.SetTrustedProxies([]string{os.Getenv("TRUSTED_PROXY")})

	router.GET("/", func(c *gin.Context) {
		url := c.Request.URL.Scheme + c.Request.Host
		c.JSON(http.StatusOK, gin.H{
			"meUrl":        url + "/me{/id}",
			"interestsUrl": url + "/interests{/slug}",
			"tilsUrl":      url + "/tilsUrl{/slug}",
			"swaggerUrl":   url + "/docs",
		})
	})
	router.GET("/me", me.GetMe)
	router.GET("/me/:id", me.GetMe)

	router.GET("/interests", interests.GetInterests)
	router.GET("/interests/:slug", interests.GetInterest)

	router.GET("/tils", tils.GetTils)
	router.GET("/tils/:slug", tils.GetTil)

	docs.SwaggerInfo.BasePath = "/"
	router.GET("/docs", func(ctx *gin.Context) {
		ctx.Redirect(http.StatusPermanentRedirect, "/swagger/index.html")
	})
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	router.Run(":" + os.Getenv("API_PORT"))
}
