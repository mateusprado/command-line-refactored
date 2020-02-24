package main

import (
	"log"
	"os"

	"github.com/mateusprado/command-line-refactored/cli"
)

func main() {

	app := cli.Builder()

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}
