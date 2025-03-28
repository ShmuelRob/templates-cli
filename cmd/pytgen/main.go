package main

import (
	"fmt"
	"os"

	"github.com/ShmuelRob/python-template-generator/internal/cli"
)

func main() {
	// Create CLI app
	app := cli.NewApp()

	// Run the CLI app
	if err := app.Run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(1)
	}
}
