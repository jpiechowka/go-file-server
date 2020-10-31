package config

type ServerConfig struct {
	Address                 string
	ServeDirectoryPath      string
	RateLimitPerMinute      uint
	EnableBasicAuth         bool
	BasicAuthUser           string
	BasicAuthPassword       string
	CompressionLevel        int
	EnableTls               bool
	DisableDirectoryListing bool
	CertFilePath            string
	KeyFilePath             string
}
