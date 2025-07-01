package routes

import (
	"net/http"
	"os"

	docs "jm/domains/api.jmjanzen.com/docs"
	"jm/domains/api.jmjanzen.com/handlers/interests"
	"jm/domains/api.jmjanzen.com/handlers/me"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Launch() {
	router := gin.Default()
	router.SetTrustedProxies([]string{os.Getenv("TRUSTED_PROXY")})

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"meUrl":        c.Request.Host + "/me{/id}",
			"interestsUrl": c.Request.Host + "/interests{/slug}",
			"swaggerUrl":   c.Request.Host + "/docs",
		})
	})
	router.GET("/me", me.GetMe)
	router.GET("/me/:id", me.GetMe)

	router.GET("/interests", interests.GetInterests)
	router.GET("/interests/:slug", interests.GetInterest)

	docs.SwaggerInfo.BasePath = "/"
	router.GET("/docs", func(ctx *gin.Context) {
		ctx.Redirect(http.StatusPermanentRedirect, "/swagger/index.html")
	})
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	router.Run(os.Getenv("API_HOST"))
}
