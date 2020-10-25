package main

import (
	"github.com/jpiechowka/go-file-server/internal/cmd"
	"os"
)

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
