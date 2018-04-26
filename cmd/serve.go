package cmd

import (
	"context"
	"github.com/spf13/cobra"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start webshare server",
	Long: `
Webshare is a CLI tool that provides web-interface for your local files`,
	Run: func(cmd *cobra.Command, args []string) {
		webshareServer(Port, IP, Path)
	},
}

// variables for command's flags
var Path string
var IP string
var Port string

func init() {
	rootCmd.AddCommand(serveCmd)

	// local flags for "webshare serve" command
	serveCmd.Flags().StringVarP(&Path,
		"dir",
		"d",
		".",
		"path to files that you want to share")
	serveCmd.Flags().StringVarP(&IP,
		"address",
		"a",
		"127.0.0.1",
		"IP address of needed interface")
	serveCmd.Flags().StringVarP(&Port,
		"port",
		"p",
		"8080",
		"listen this port")
}

func webshareServer(Port string, IP string, Path string) {
	log.Printf("webshare server started on IP: %v, Port: %v, Path to files: %v", IP, Port, Path)
	ipAddressPort := IP + ":" + Port

	// handler for path
	fs := http.FileServer(http.Dir(Path))
	http.Handle("/", fs)

	// run a webserver
	Server := http.Server{Addr: ipAddressPort}
	go func() {
		log.Fatal(Server.ListenAndServe())
	}()

	// graceful shutdown
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	<-signalChan
	log.Println("Shutdown signal received, exiting...")
	Server.Shutdown(context.Background())
}
