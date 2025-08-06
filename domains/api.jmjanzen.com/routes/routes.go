package routes

import (
	"os"

	"jm/domains/api.jmjanzen.com/handlers/interests"
	"jm/domains/api.jmjanzen.com/handlers/modes"
	"jm/domains/api.jmjanzen.com/handlers/tils"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Launch() {
	router := gin.Default()
	router.SetTrustedProxies([]string{os.Getenv("TRUSTED_PROXY")})

	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	router.GET("/me", modes.GetModes)
	router.GET("/me/:id", modes.GetModes)
	router.GET("/modes", modes.GetModes)
	router.GET("/modes/:id", modes.GetModes)

	router.GET("/interests", interests.GetInterests)
	router.GET("/interests/:slug", interests.GetInterest)

	router.GET("/tils", tils.GetTils)
	router.GET("/tils/:slug", tils.GetTil)

	router.Run(":" + os.Getenv("API_PORT"))
}
