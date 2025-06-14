package webserver

import (
	"log"
	"os"
	"path"
)

// Change dir to given path, logging action
func Root(targetPath string) {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	targetWd := path.Join(cwd, targetPath)

	log.Printf("Changing Working Directory from '%v' to '%v'\n", cwd, targetWd)

	os.Chdir(targetWd)
}

