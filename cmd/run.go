// cmd/run.go
package cmd

import (
	"log"
	"strconv"
	"sync/atomic"
	"time"

	"github.com/geekxflood/orion/internal/helpers"
	"github.com/geekxflood/orion/internal/httpclient"
	"github.com/geekxflood/orion/internal/localtypes"
	module_glpi "github.com/geekxflood/orion/internal/modules/glpi"
	module_phpipam "github.com/geekxflood/orion/internal/modules/phpipam"
	"github.com/spf13/cobra"
)

// Global variables
var (
	conf atomic.Value // Use atomic.Value to store the configuration
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Start the application and run the selected modules",
	Long: `The run command starts the application and runs the selected modules.
	It initializes the configuration by reading from a config file, sets the port and insecure options,
	and converts the interval to an integer value.
	Then, it stores the initial configuration in an atomic.Value and starts a goroutine to periodically
	refresh the configuration by reading from the config file.
	After that, it creates an HTTP client and passes the atomic.Value to it.
	Finally, it selects the modules to run based on the configuration and runs the HTTP client.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Initialize the config
		initialConf, err := helpers.ReadConfig(cfgFile)
		if err != nil {
			log.Fatalf("Error reading config file: %s", err)
		}
		initialConf.Port = port
		initialConf.Insecure = insecure

		// convert initialConf.Interval to int and store it in interval
		interval, err := strconv.Atoi(initialConf.Interval)
		if err != nil {
			log.Printf("Error converting interval to int: %s", err)
			log.Println("Using default interval of 5 seconds")
			interval = 5
		}

		conf.Store(&initialConf) // Store the initial configuration in an atomic.Value

		// Start the goroutine to refresh the config
		go func() {
			for {
				time.Sleep(time.Duration(interval) * time.Second)
				newConf, err := helpers.ReadConfig(cfgFile)
				if err != nil {
					log.Printf("Error refreshing config: %s", err)
					continue
				}
				conf.Store(&newConf) // Store the updated configuration in an atomic.Value
				log.Printf("Configuration refreshed")
			}
		}()

		// Create the HTTP client and pass the atomic.Value
		httpClient := httpclient.Client{Conf: &conf}

		// Select the modules to run
		switch initialConf.Modules {
		case "file":
			log.Println("File module is enabled")

			// if conf.endpoint is empty, use empty array
			if initialConf.Endpoints == nil {
				initialConf.Endpoints = []localtypes.Endpoints{}
			}

			// Run the http client
			httpClient.RunClient(initialConf.Endpoints)

		case "glpi":
			log.Println("GLPI module is enabled")

			// run the glpi module
			module_glpi.GetGLPI()

			// Run the http client
			httpClient.RunClient(initialConf.Endpoints)


		case "phpipam":
			log.Println("phpipam module is enabled")

			// run the phpipam module
			module_phpipam.GetphpIPAM()

			// Run the http client
			httpClient.RunClient(initialConf.Endpoints)

		default:
			log.Fatalln("No modules defined")
		}
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}
