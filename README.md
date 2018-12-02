# webshare

`webshare` is a CLI tool that provides web-interface for your local files.

## How to use this app

Download binary for your system:

[Windows](https://github.com/mxssl/webshare/releases/download/v0.0.2/webshare-windows-amd64.exe)

[Linux](https://github.com/mxssl/webshare/releases/download/v0.0.2/webshare-linux-amd64)

[MacOS](https://github.com/mxssl/webshare/releases/download/v0.0.2/webshare-darwin-amd64)

Rename binary for convinient use:

Windows: `ren webshare-windows-amd64.exe webshare.exe`

Linux: `mv webshare-linux-amd64 webshare`

MacOS: `mv webshare-darwin-amd64 webshare`

Then just run:

Windows: `webshare.exe`

Linux: `./webshare`

MacOS: `./webshare`

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
  -a, --address string   IP address of needed interface (default "0.0.0.0")
  -d, --dir string       path to files that you want to share (default ".")
  -h, --help             help for serve
  -p, --port string      listen this port (default "8080")
```

*  Start webshare server: `webshare serve`

*  Open your browser http://your_ip:8080 and you will see files in directory that you picked

*  To stop this app use `ctrl + c` combination
