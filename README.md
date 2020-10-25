# Go File Server
A simple web server written in Go that can be used to serve files and transfer them (I use it to transfer files to my mobile devices on LAN).

## Features
* Ability to configure serve directory
* Automatically add security and cache headers to the responses
* Built-in request logging
* Built-in and configurable request rate limiter
* Automated generation of self-signed TLS certs (TODO)
* Ability to add custom TLS certs (TODO)
* Configurable compression (TODO: configuration)
* Configurable Basic Auth

## Usage
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

Start command can be used to start the server and configure options:
```
Start command starts the builtin Fiber server to serve static files

Usage:
  go-file-server start [flags]

Flags:
  -a, --address string      server address (default "0.0.0.0:13337")
  -b, --basic-auth string   enables Basic Auth. Credentials should be provided as username:password
  -d, --dir string          path to directory with files to serve (default "./files")
  -h, --help                help for start
  -r, --rate-limit uint     configure max requests per minute (default 60)
```
