package routes

import (
	"os"

	"jm/domains/blog.jmjanzen.com/handlers"

	"github.com/gin-gonic/gin"
)

func Launch() {
	router := gin.Default()

	router.SetTrustedProxies([]string{os.Getenv("TRUSTED_PROXY")})

	router.LoadHTMLGlob("./domains/blog.jmjanzen.com/blogs/*")

	router.GET("/", handlers.Default)
	router.GET("/:slug", handlers.HandleBlog)

	router.Run(":" + os.Getenv("BLOG_PORT"))
}
