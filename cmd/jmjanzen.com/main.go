package main

import (
	_ "github.com/joho/godotenv/autoload"
	"jm/domains/jmjanzen.com/routes"
	webserver "jm/internal/utils/common"
)

func main() {
	webserver.Root("domains/jmjanzen.com")

	routes.Launch()
}
