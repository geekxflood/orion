// Package cmd cmd/init.go

package cmd

import (
	"fmt"
	"os"

	"github.com/geekxflood/orion/internal/helpers"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize the configuration file",

	Run: func(cmd *cobra.Command, args []string) {
		homePath, err := helpers.GetHomePath()

		if err != nil {
			fmt.Println("Error getting the home path")
			os.Exit(1)
		}

		if _, err := os.Stat(homePath + "/.orion/config.yaml"); err == nil {
			fmt.Println("Configuration file already exists. Do you want to overwrite it? (yes/No)")
			var response string
			fmt.Scanln(&response)
			if (response == "yes") || (response == "y") {
				fmt.Println("Overwriting the configuration file")
			} else {
				fmt.Println("Exiting the command")
				os.Exit(0)
			}
		}

		fmt.Printf("Creating the configuration file at %v/.orion/config.yaml\n", homePath)

		if _, err := os.Stat(homePath + "/.orion"); err != nil {
			err := os.Mkdir(homePath+"/.orion", 0755)
			if err != nil {
				fmt.Println("Error creating the .orion directory")
				panic(err)
			}
		}
		file, err := os.Create(homePath + "/.orion/config.yaml")
		if err != nil {
			panic(err)
		}
		defer file.Close()

		// write the default configuration to the file
		_, err = file.WriteString(`# Default configuration, update values accordingly
---
module: module_name
port: 9981
insecure: false
interval: 5
targets: []
`)
		if err != nil {
			fmt.Println("Error writing to the configuration file")
			os.Exit(1)
		}

		fmt.Println("Configuration file created successfully")

	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
