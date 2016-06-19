package main

import (
	"log"
	"os"

	"github.com/progrium/go-basher"
)

func main() {
	os.Setenv("WAFFLES_DIR", "bash")
	bash, _ := basher.NewContext("/bin/bash", false)
	bash.CopyEnv()
	bash.Source("bash/init.sh", Asset)
	status, err := bash.Run("source \"${1:-/dev/stdin}\"", os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	os.Exit(status)
}
