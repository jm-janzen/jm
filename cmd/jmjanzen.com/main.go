package main

import (
	"jmjanzen/domains/jmjanzen.com/routes"
	"jmjanzen/internal/utils/common"
)


func main() {
	// Change dir to project root, if not already there
	common.ChdirWebserverRoot()

	routes.Launch()
}
