package main

import (
	"jm/domains/jmjanzen.com/routes"
	webserver "jm/internal/utils/common"
)


func main() {
	webserver.Root("domains/jmjanzen.com")

	routes.Launch()
}
