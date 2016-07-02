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

	// If an argument is given, treat it as the input file.
	// If not, check if there is anything being piped in.
	// This will need modified in the future to support flags.
	var wscript string
	if len(os.Args) > 1 {
		wscript = fmt.Sprintf("source %s", os.Args[1])
	} else {
		w, err := readStdinPipe()
		if err != nil {
			log.Fatal(err)
		}
		wscript = w
	}

	status, err := bash.Run(wscript, []string{})
	if err != nil {
		log.Fatal(err)
	}
	os.Exit(status)
}

func readStdinPipe() (string, error) {
	var bytes []byte
	stat, _ := os.Stdin.Stat()
	if stat.Mode()&os.ModeNamedPipe != 0 {
		b, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			return "", err
		}
		bytes = b
	}

	if len(bytes) == 0 {
		e := fmt.Errorf("No input received.")
		return "", e
	}

	return string(bytes), nil
}
