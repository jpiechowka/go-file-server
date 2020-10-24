package cmd

import (
	"fmt"
	"github.com/jpiechowka/go-file-server/internal/config"
	"github.com/jpiechowka/go-file-server/internal/server"
	"github.com/spf13/cobra"
)

const serverDefaultAddr = "0.0.0.0:13337"

var (
	serverAddr string
	//enableTls         bool
	//enableBasicAuth   bool
	//basicAuthUser     string
	//basicAuthPassword string

	rootCmd = &cobra.Command{
		Version: "0.1.0",
		Short:   "A file server built in Go using Fiber",
	}
)

// Execute executes the root command.
func Execute() error {
	err := rootCmd.Execute()
	if err != nil {
		return err
	}

	cfg := &config.ServerConfig{
		Address: serverAddr,
	}

	srv := server.NewFiberFileServer(cfg)
	return srv.ConfigureAndStart()
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&serverAddr, "address", "a", serverDefaultAddr, fmt.Sprintf("Server address (default is %s)", serverDefaultAddr))
}
