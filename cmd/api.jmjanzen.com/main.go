package main

import (

	"jm/domains/api.jmjanzen.com/routes"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	routes.Launch()
}
