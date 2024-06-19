// Package cmd cmd/run.go

package cmd

import (
	"log"
	"sync/atomic"
	"time"

	"github.com/geekxflood/orion/internal/helpers"
	"github.com/geekxflood/orion/internal/httpclient"
	"github.com/geekxflood/orion/internal/localtypes"
	"github.com/spf13/cobra"
)

// Global variables
var (
	conf     atomic.Value // Use atomic.Value to store the configuration
	interval int          // Declare the 'interval' variable
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Start the application and run the selected modules",
	Long: `The run command starts the application and runs the selected modules.
	It initializes the configuration, sets the port and insecure options, and runs the HTTP client.`,
	Run: func(cmd *cobra.Command, args []string) {
		if cfgFile == "" {
			log.Println("No config file specified, using default config file in $HOME/.orion/config.yml")
			homePath, err := helpers.GetHomePath()
			if err != nil {
				log.Fatalf("Error getting the home path: %s", err)
			}
			cfgFile = homePath + "/.orion/config.yml"
		}

		initialConf, err := helpers.ReadConfig(cfgFile)
		if err != nil {
			log.Fatalf("Error reading config file: %s", err)
		}
		initialConf.Port = port
		initialConf.Insecure = insecure

		conf.Store(&initialConf)

		go func() {
			for {
				time.Sleep(time.Duration(interval) * time.Second)
				newConf, err := helpers.ReadConfig(cfgFile)
				if err != nil {
					log.Printf("Error refreshing config: %s", err)
					continue
				}
				conf.Store(&newConf)
				log.Printf("Configuration refreshed")
			}
		}()

		httpClient := httpclient.Client{Conf: &conf}

		switch initialConf.Modules {
		case "file":
			log.Println("File module is enabled")
			if initialConf.Endpoints == nil {
				initialConf.Endpoints = []localtypes.Endpoints{}
			}
			httpClient.RunClient(initialConf.Endpoints)

		case "http":
			log.Println("HTTP module is enabled")

		default:
			log.Fatalln("No modules defined")
		}
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}
