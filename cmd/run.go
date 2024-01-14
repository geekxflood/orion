// cmd/run.go
package cmd

import (
	"log"

	"github.com/geekxflood/orion/internal/helpers"
	"github.com/geekxflood/orion/internal/httpclient"
	"github.com/spf13/cobra"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:     "run",
	Short:   "Start the orion application",
	Example: `orion run -c config.yml`,
	Run: func(cmd *cobra.Command, args []string) {

		conf.Port = port
		conf.Insecure = insecure

		conf, err = helpers.ReadConfig(cfgFile)

		switch conf.Modules {
		case "file":
			log.Println("File module is enabled")
		default:
			log.Fatalln("No modules defined")
		}

		if err != nil {
			log.Panicf("Error reading default config file: %s", err)
		}
		httpclient.RunClient(conf)

	},
}

func init() {
	rootCmd.AddCommand(runCmd)

}
