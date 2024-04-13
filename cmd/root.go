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
	Use:   "orion",
	Short: "Orion is a tool for serving prometheus target configurations",
	Long: `Orion is a tool for serving prometheus target configurations. 
	It provides a simple and efficient way to manage and serve target configurations for Prometheus.
	With Orion, you can easily update and distribute target configurations to multiple Prometheus instances.
	For more information and usage examples, please visit the official Orion GitHub repository: https://github.com/geekxflood/orion.
	Default port is 9981 and default config file is config.yml.`,
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
