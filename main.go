package main

import (
	"os"

	"github.com/jpiechowka/go-file-server/internal/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
