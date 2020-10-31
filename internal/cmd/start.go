package cmd

import (
	"errors"
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	"github.com/jpiechowka/go-file-server/internal/config"
	"github.com/jpiechowka/go-file-server/internal/server"
	"github.com/jpiechowka/go-file-server/internal/tls"
)

const (
	defaultServerAddr              = "0.0.0.0:13337"
	defaultServeDirectoryPath      = "./files"
	defaultRateLimitPerMinute      = 60
	defaultCompressionLevel        = 2
	defaultGenerateSelfSignedCerts = false
	defaultEnableTls               = false
	defaultTlsCertificateHosts     = "localhost"
	defaultDisableDirListing       = false

	certFilePath = "cert.pem"
	keyFilePath  = "key.pem"
)

var (
	serverAddr              string
	serveDirectoryPath      string
	basicAuthCredentials    string
	rateLimitPerMinute      uint
	compressionLevel        int
	generateSelfSignedCerts bool
	enableTls               bool
	tlsCertificateHosts     string
	disableDirListing       bool

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
	startCommand.Flags().IntVarP(&compressionLevel, "compression", "c", defaultCompressionLevel, "configure compression level. -1 to disable, 0 for default level, 1 for best speed, 2 for best compression")
	startCommand.Flags().BoolVarP(&generateSelfSignedCerts, "generate-cert", "g", defaultGenerateSelfSignedCerts, fmt.Sprintf("enable TLS and generate self-signed certs for the server. Outputs to '%s' and '%s' and will overwrite existing files", certFilePath, keyFilePath))
	startCommand.Flags().BoolVarP(&enableTls, "tls", "t", defaultEnableTls, fmt.Sprintf("enables TLS. Files should be saved as '%s' and '%s'", certFilePath, keyFilePath))
	startCommand.Flags().StringVar(&tlsCertificateHosts, "cert-hosts", defaultTlsCertificateHosts, "comma separated list of DNS names (Subject Alt Names extension). Used only when generating self-signed certs. Example values: example1.com,example2.com")
	startCommand.Flags().BoolVarP(&disableDirListing, "disable-dir-listing", "l", defaultDisableDirListing, "disables directory listing which is turned on by default")
}

func startCmd() error {
	cfg := &config.ServerConfig{
		Address:                 serverAddr,
		ServeDirectoryPath:      serveDirectoryPath,
		RateLimitPerMinute:      rateLimitPerMinute,
		DisableDirectoryListing: disableDirListing,
		CertFilePath:            certFilePath,
		KeyFilePath:             keyFilePath,
	}

	// Basic Auth
	if basicAuthCredentials != "" {
		credentials := strings.Split(basicAuthCredentials, ":")

		if len(credentials) != 2 {
			return errors.New("provided Basic Auth credentials are invalid. Expected format is username:password")
		}

		if credentials[0] == "" || credentials[1] == "" {
			return errors.New("provided Basic Auth credentials are invalid. Password and username cannot be empty")
		}

		cfg.EnableBasicAuth = true
		cfg.BasicAuthUser = credentials[0]
		cfg.BasicAuthPassword = credentials[1]
	}

	// Compression
	if compressionLevel < -1 || compressionLevel > 2 {
		return errors.New("provided compression level is invalid. Valid values are -1, 0, 1 and 2")
	}
	cfg.CompressionLevel = compressionLevel

	// TLS
	if generateSelfSignedCerts {
		if err := tls.GenerateSelfSignedCertAndKey(certFilePath, keyFilePath, tlsCertificateHosts); err != nil {
			return err
		}

		enableTls = true
	}
	cfg.EnableTls = enableTls

	srv := server.NewFiberFileServer(cfg)
	return srv.ConfigureAndStart()
}
