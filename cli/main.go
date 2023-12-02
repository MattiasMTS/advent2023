package main

import (
	"log"
	"os"

	"github.com/MattiasMTS/advent2023/cli/cmd"
)

func main() {
	if err := cmd.App.Run(os.Args); err != nil {
		log.Println(err)
	}

	// cleanup
	os.Unsetenv("_SUBMIT")
	os.Exit(0)
}
