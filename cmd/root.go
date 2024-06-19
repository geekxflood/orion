// cmd/root.go
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var (
	cfgFile  string
	insecure bool
	port     string
)

var rootCmd = &cobra.Command{
	Use:     "orion",
	Short:   "Prometheus target configuration server",
	Long:    "Orion is a tool for serving prometheus target configurations.",
	Version: "dev",
}

// Execute runs the root command
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file (default is $HOME/.orion/config.yml)")
	rootCmd.PersistentFlags().BoolVarP(&insecure, "insecure", "i", false, "disable TLS verification")
	rootCmd.PersistentFlags().StringVarP(&port, "port", "p", "9981", "port to listen on")
}
