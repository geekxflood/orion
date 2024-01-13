package helpers

import (
	"errors"
	"log"
	"os"

	"github.com/geekxflood/orion/internal/localtypes"
	"github.com/geekxflood/orion/internal/unmarshaller"
)

// ReadConfig reads a config file and returns a Config struct
func ReadConfig(file string) (localtypes.Config, error) {

	// LookupConfig looks for a config file in the current directory
	path, format, err := LookupConfig(file)
	if err != nil {
		return localtypes.Config{}, err
	}

	// UnmarshalConfig unmarshals a config file into a Config struct
	conf, err := unmarshaller.UnmarshalConfig(path, format)
	if err != nil {
		return localtypes.Config{}, err
	}

	return conf, nil
}

// LookupConfig looks for a config file in the current directory
func LookupConfig(file string) (string, string, error) {
	// if a config file is specified, use it
	if file != "" {
		log.Printf("Using config file: %s", file)
		if _, err := os.Stat(file); err == nil {

			if file[len(file)-4:] == "json" {
				log.Println("Found json format")
				return file, "json", nil
			}

			if file[len(file)-3:] == "yml" || file[len(file)-4:] == "yaml" {
				log.Println("Found yaml format")
				return file, "yaml", nil
			}

			if file[len(file)-4:] == "toml" {
				log.Println("Found toml format")
				return file, "toml", nil
			}
		}
		// if no config file is specified, use the default
	} else {
		log.Println("No config file specified, looking for default config file")
		if _, err := os.Stat("config.yml"); err == nil {
			log.Println("Found config.yml")
			return "config.yml", "yaml", nil
		}

		if _, err := os.Stat("config.yaml"); err == nil {
			log.Println("Found config.yaml")
			return "config.yaml", "yaml", nil
		}

		if _, err := os.Stat("config.json"); err == nil {
			log.Println("Found config.json")
			return "config.json", "json", nil
		}

		if _, err := os.Stat("config.toml"); err == nil {
			log.Println("Found config.toml")
			return "config.toml", "toml", nil
		}
	}
	// No config has been found
	return "", "", errors.New("no config file found")
}
