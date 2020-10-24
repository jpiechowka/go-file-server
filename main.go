package main

import (
	"github.com/jpiechowka/go-file-server/internal/cmd"
	"os"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
