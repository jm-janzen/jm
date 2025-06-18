package main

import (
	_ "github.com/joho/godotenv/autoload"
	"jm/domains/jmjanzen.com/routes"
)

func main() {
	routes.Launch()
}
