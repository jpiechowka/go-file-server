package config

type ServerConfig struct {
	Address            string
	ServeDirectoryPath string
	EnableBasicAuth    bool
	BasicAuthUser      string
	BasicAuthPassword  string
}
