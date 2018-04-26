# Webshare

[![Build Status](https://travis-ci.org/mxssl/webshare.svg?branch=master)](https://travis-ci.org/mxssl/webshare)

`Webshare` is a CLI tool that provides web-interface for your local files.

## How to use this programm

Download binary for your system:

Windows: https://github.com/mxssl/webshare/releases/download/v1.0.0/webshare-1.0.0-windows-amd64.exe

Linux: https://github.com/mxssl/webshare/releases/download/v1.0.0/webshare-1.0.0-linux-amd64

Rename binary for convinient use:

Windows: `ren webshare-1.0.0-windows-amd64.exe webshare.exe`

Linux: `mv webshare-1.0.0-linux-amd64 webshare`

Then just run:

Windows: `webshare.exe`

Linux: `webshare`

```
  webshare [command]

Available Commands:
  help        Help about any command
  serve       Start webshare server
  version     Print the version number of Webshare
```

`webshare serve` - main command that you need

```
  webshare serve [flags]

Flags:
  -a, --address string   IP address of needed interface (default "127.0.0.1")
  -d, --dir string       path to files that you want to share (default ".")
  -h, --help             help for serve
  -p, --port string      listen this port (default "8080")
```

*  Determine your local IP address

**Windows:** `ipconfig`

**Linux:** `ip addr`

*  Start webshare server: `webshare serve -a 192.168.88.10 -d . -p 8080`

*  Open your browser http://192.168.88.10:8080 and you will see files in directory that you picked
