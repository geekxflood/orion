// cmd/root.go
package cmd

import (
	"os"

	"github.com/geekxflood/orion/internal/localtypes"
	"github.com/spf13/cobra"
)

var conf localtypes.Config
var err error

var cfgFile string
var insecure bool
var port string

var rootCmd = &cobra.Command{
	Use:   "orion",
	Short: "Orion is a tool for serving prometheus target configurations",
	Long: `Orion is a tool for serving prometheus target configurations
		Complete documentation is available at https://github.com/geekxflood/orion
		Default port is 9981 and default config file is config.yml`,
	Version: "dev",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file (default is config.yml)")
	rootCmd.PersistentFlags().BoolVarP(&insecure, "insecure", "i", false, "disable TLS verification")
	rootCmd.PersistentFlags().StringVarP(&port, "port", "p", "9981", "port to listen on")
}
