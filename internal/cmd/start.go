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

	startCommand = &cobra.Command{
		Use:   "start",
		Short: "Start the server",
		Long:  `Start command starts the builtin Fiber server to serve static files`,
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := startCmd(); err != nil {
				return err
			}
			return nil
		},
	}
)

func init() {
	rootCmd.LocalNonPersistentFlags().StringVarP(&serverAddr, "address", "a", serverDefaultAddr, fmt.Sprintf("Server address (default is %s)", serverDefaultAddr))
}

func startCmd() error {
	cfg := &config.ServerConfig{
		Address: serverAddr,
	}

	srv := server.NewFiberFileServer(cfg)
	return srv.ConfigureAndStart()
}
