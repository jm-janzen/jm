package main

import (
	_ "github.com/joho/godotenv/autoload"
	"jm/domains/jmjanzen.com/routes"
	"jm/internal/webserver"
)

func main() {
	webserver.Root("domains/jmjanzen.com")

	routes.Launch()
}
