package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command
var rootCmd = &cobra.Command{
	Use:   "webshare",
	Short: "Webshare - fast way to share your files",
	Long: `
Webshare is a CLI tool that provides web-interface for your local files.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
