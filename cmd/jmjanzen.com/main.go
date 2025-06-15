package main

import (
	"jm/domains/jmjanzen.com/routes"
	webserver "jm/internal/utils/common"
	_ "github.com/joho/godotenv/autoload"
)


func main() {
	webserver.Root("domains/jmjanzen.com")

	routes.Launch()
}
