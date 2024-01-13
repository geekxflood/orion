// main.go
package main

import (
	"log"

	"github.com/geekxflood/orion/cmd"
	"github.com/geekxflood/orion/internal/helpers"
)

func main() {
	// read local config
	config, err := helpers.ReadConfig()
	if err != nil {
		panic(err)
	}

	log.Println("Config read successfully")
	log.Printf("Config: %v\n", config)
	cmd.Execute()
}
