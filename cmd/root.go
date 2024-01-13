// cmd/root.go
package cmd

import (
	"log"
	"os"

	"github.com/geekxflood/orion/internal/helpers"
	"github.com/geekxflood/orion/internal/localtypes"
	"github.com/spf13/cobra"
)

var cfgFile string
var insecure bool

var rootCmd = &cobra.Command{
	Use:   "orion",
	Short: "Orion is a tool for serving prometheus target configurations",
	Long: `Orion is a tool for serving prometheus target configurations
		Complete documentation is available at https://github.com/geekxflood/orion`,
	Example: `orion -c config.yml`,
	Version: "dev",

	Run: func(cmd *cobra.Command, args []string) {
		var conf localtypes.Config
		var err error

		conf, err = helpers.ReadConfig(cfgFile)
		if err != nil {
			log.Panicf("Error reading default config file: %s", err)
		}

		log.Printf("Config: %+v", conf)
	},
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
}
