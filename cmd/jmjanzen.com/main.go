package main

import (
	"jm/domains/jmjanzen.com/routes"
	"jm/internal/utils/common"
)


func main() {
	// Change dir to project root, if not already there
	common.ChdirWebserverRoot()

	routes.Launch()
}
