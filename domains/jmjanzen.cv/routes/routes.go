package routes

import (
	"os"

	"jm/domains/jmjanzen.cv/handlers"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)

func Launch() {
	router := gin.Default()

	router.SetTrustedProxies([]string{os.Getenv("TRUSTED_PROXY")})
	router.LoadHTMLGlob("./domains/jmjanzen.cv/static/*")

	router.GET("/", handlers.Default)

	router.Run(os.Getenv("CV_HOST"))
}

