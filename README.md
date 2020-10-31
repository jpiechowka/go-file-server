# Go File Server
A simple web server written in Go that can be used to serve files and transfer them (I use it to transfer files to my devices on LAN).

## Features
* Built using https://github.com/gofiber/fiber
* Ability to configure serve directory
* Directory listing / browsing is enabled by default. It can be switched off by providing a correct cli flag
* Automatically add security and cache headers to the responses
* Built-in request logging
* Built-in and configurable request rate limiter
* Automated generation of self-signed TLS certs with configurable DNS names (Subject Alt Names extension)
* Ability to add custom / trusted TLS certs by saving them with correct names and providing one command line flag
* Configurable compression and ability to disable it
* Configurable Basic Auth

## Usage
Basic usage with default options: ```go-file-server start```. It will start server without TLS and serve files from ```./files``` directory.

```
A file server built in Go using Fiber

Usage:
  go-file-server [command]

Available Commands:
  help        Help about any command
  start       Start the server

Flags:
  -h, --help      help for go-file-server
  -v, --version   version for go-file-server

Use "go-file-server [command] --help" for more information about a command.
```

```start``` command can be used to start the server and configure additional options:
```
Start command starts the builtin Fiber server to serve static files

Usage:
  go-file-server start [flags]

Flags:
  -a, --address string        server address (default "0.0.0.0:13337")
  -b, --basic-auth string     enables Basic Auth. Credentials should be provided as username:password
      --cert-hosts string     comma separated list of DNS names (Subject Alt Names extension). Used only when generating self-signed certs. Example values: example1.com,example2.com (default "localhost")
  -c, --compression int       configure compression level. -1 to disable, 0 for default level, 1 for best speed, 2 for best compression (default 2)
  -d, --dir string            path to directory with files to serve (default "./files")
  -l, --disable-dir-listing   disables directory listing which is turned on by default
  -g, --generate-cert         enable TLS and generate self-signed certs for the server. Outputs to 'cert.pem' and 'key.pem' and will overwrite existing files
  -h, --help                  help for start
  -r, --rate-limit uint       configure max requests per minute (default 60)
  -t, --tls                   enables TLS. Files should be saved as 'cert.pem' and 'key.pem'
```
