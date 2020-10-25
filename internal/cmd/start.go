package cmd

import (
	"github.com/jpiechowka/go-file-server/internal/config"
	"github.com/jpiechowka/go-file-server/internal/server"
	"github.com/spf13/cobra"
)

const (
	defaultServerAddr         = "0.0.0.0:13337"
	defaultServeDirectoryPath = "./files"
)

var (
	serverAddr         string
	serveDirectoryPath string
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
	startCommand.Flags().StringVarP(&serverAddr, "address", "a", defaultServerAddr, "Server address")
	startCommand.Flags().StringVarP(&serveDirectoryPath, "dir", "d", defaultServeDirectoryPath, "Path to directory with files to serve")
}

func startCmd() error {
	cfg := &config.ServerConfig{
		Address:            serverAddr,
		ServeDirectoryPath: serveDirectoryPath,
	}

	srv := server.NewFiberFileServer(cfg)
	return srv.ConfigureAndStart()
}
