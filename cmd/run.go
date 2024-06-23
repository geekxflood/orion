// Package cmd cmd/run.go

package cmd

import (
	"log"
	"sync/atomic"
	"time"

	"github.com/geekxflood/orion/internal/helpers"
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
		// If the configuration path file is not specified, use the default one
		if cfgFilePath == "" {
			log.Println("No config file specified, using default config file in $HOME/.orion/config.yml")
			homePath, err := helpers.GetHomePath()
			if err != nil {
				log.Fatalf("Error getting the home path: %s", err)
			}
			cfgFilePath = homePath + "/.orion/config.yml"
		}

		// Read the initial configuration
		initialConf, err := helpers.ReadConfig(cfgFilePath)
		if err != nil {
			log.Fatalf("Error reading config file: %s", err)
		}

		// Override the configuration with the provided flags values
		initialConf.Port = port
		initialConf.Insecure = insecure

		// Store the configuration
		conf.Store(&initialConf)

		// Refresh the configuration at the specified interval
		go func() {
			for {
				time.Sleep(time.Duration(interval) * time.Second)
				newConf, err := helpers.ReadConfig(cfgFilePath)
				if err != nil {
					log.Printf("Error refreshing config: %s", err)
					continue
				}
				conf.Store(&newConf)
				log.Printf("Configuration refreshed")
			}
		}()

		// For each type of module run there respective client in a goroutine
		for _, module := range initialConf.Modules {
			switch module.Type {
			case "rest":
				go helpers.RunRESTClient(&module)
			case "file":
				go helpers.RunFileClient(&module)
			default:
				log.Fatalf("Unknown endpoint type: %s", module.Type)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}
