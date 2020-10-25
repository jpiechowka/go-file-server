package cmd

import (
	"errors"
	"github.com/jpiechowka/go-file-server/internal/config"
	"github.com/jpiechowka/go-file-server/internal/server"
	"github.com/spf13/cobra"
	"strings"
)

const (
	defaultServerAddr         = "0.0.0.0:13337"
	defaultServeDirectoryPath = "./files"
	defaultRateLimitPerMinute = 60
)

var (
	serverAddr           string
	serveDirectoryPath   string
	basicAuthCredentials string
	rateLimitPerMinute   uint

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
	startCommand.Flags().StringVarP(&serverAddr, "address", "a", defaultServerAddr, "server address")
	startCommand.Flags().StringVarP(&serveDirectoryPath, "dir", "d", defaultServeDirectoryPath, "path to directory with files to serve")
	startCommand.Flags().StringVarP(&basicAuthCredentials, "basic-auth", "b", "", "enables Basic Auth. Credentials should be provided as username:password")
	startCommand.Flags().UintVarP(&rateLimitPerMinute, "rate-limit", "r", defaultRateLimitPerMinute, "configure max requests per minute")
}

func startCmd() error {
	cfg := &config.ServerConfig{
		Address:            serverAddr,
		ServeDirectoryPath: serveDirectoryPath,
		RateLimitPerMinute: rateLimitPerMinute,
	}

	if basicAuthCredentials != "" {
		cfg.EnableBasicAuth = true

		credentials := strings.Split(basicAuthCredentials, ":")

		if len(credentials) != 2 {
			return errors.New("provided Basic Auth credentials are invalid. Expected format is username:password")
		}

		if credentials[0] == "" || credentials[1] == "" {
			return errors.New("provided Basic Auth credentials are invalid. Password and username cannot be empty")
		}

		cfg.BasicAuthUser = credentials[0]
		cfg.BasicAuthPassword = credentials[1]
	}

	srv := server.NewFiberFileServer(cfg)
	return srv.ConfigureAndStart()
}
