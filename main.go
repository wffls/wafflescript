package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/progrium/go-basher"
)

func main() {
	os.Setenv("WAFFLES_DIR", "bash")
	os.Setenv("WAFFLES_NO_HELP", "true")
	bash, _ := basher.NewContext("/bin/bash", false)
	bash.CopyEnv()

	for file, _ := range _bindata {
		if file == "bash/init.sh" {
			continue
		}
		bash.Source(file, Asset)
	}

	// Read in the first argument, which is the wafflescript.
	// This will need modified in the future to support flags.
	wscript, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	cmd := fmt.Sprintf("echo \"%s\" | source", wscript)
	status, err := bash.Run(cmd, os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	os.Exit(status)
}
