package cmd

import (
	"context"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/spf13/cobra"
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

// Path to files
var Path string

// IP address
var IP string

// Port for tcp socket
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
		"0.0.0.0",
		"IP address of needed interface")
	serveCmd.Flags().StringVarP(&Port,
		"port",
		"p",
		"8080",
		"listen this port")
}

func webshareServer(Port string, IP string, Path string) {
	ipAddressPort := IP + ":" + Port

	log.Printf("webshare server started on IP: %v, Port: %v, Path to files: %v\n", IP, Port, Path)

	localIP, err := getLocalIP()
	if err != nil {
		log.Printf("cannot obtain local ip. Error: %v\n", err)
	}

	log.Printf("local url: http://%s:%s\n", localIP, Port)

	globalIP, err := getGlobalIP()
	if err == nil {
		log.Printf("global url: http://%s:%s\n", globalIP, Port)
		log.Println("you need to have public static ip or NAT configured to use global url")
	} else {
		log.Printf("cannot abtain global ip. Error: %v\n", err)
	}

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

func getLocalIP() (string, error) {
	conn, err := net.Dial("udp", "8.8.8.8:53")
	if err != nil {
		log.Println(err)
		return "", err
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	ip := localAddr.IP.String()

	return ip, nil
}

func getGlobalIP() (string, error) {

	res, err := http.Get("http://ntwrk.cf")
	if err != nil {
		log.Println(err)
		return "", err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
		return "", err
	}

	ip := string(body)

	return strings.TrimSpace(ip), nil
}
