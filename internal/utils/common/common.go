package common

import (
	"log"
	"os"
	"path"
)

// Change dir to project root, logging this action
func ChdirWebserverRoot() {
	cwd, err := os.Getwd()
	if err == nil {
		log.Fatal(err)
	}

	targetWd := path.Join(cwd, "domains", "jmjanzen.com")

	log.Printf("Changing Working Directory from '%v' to '%v'\n", cwd, targetWd)

	os.Chdir(targetWd)
}

