package routes

import (
	"net/http"
	"os"

	"jm/domains/api.jmjanzen.com/handlers/me"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)

func Launch() {
	router := gin.Default()
	router.SetTrustedProxies([]string{os.Getenv("API_TRUSTED_PROXY")})

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"meUrl": "https://api.jmjanzen.com/me{/id}",
		})
	})
	router.GET("/me", me.GetMe)
	router.GET("/me/:id", me.GetMe)

	router.Run(os.Getenv("API_HOST"))
}
