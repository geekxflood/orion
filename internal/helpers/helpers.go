package helpers

import (
	"errors"
	"log"
	"os"

	"github.com/geekxflood/orion/internal/localtypes"
	"github.com/geekxflood/orion/internal/unmarshaller"
)

// ReadConfig reads a config file and returns a Config struct
func ReadConfig() (localtypes.Config, error) {

	path, format, err := LookupConfig()
	if err != nil {
		return localtypes.Config{}, err
	}

	conf, err := unmarshaller.UnmarshalConfig(path, format)
	if err != nil {
		return localtypes.Config{}, err
	}

	return conf, nil
}

// LookupConfig looks for a config file in the current directory
func LookupConfig() (string, string, error) {
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

	if _, err := os.Stat("config.xml"); err == nil {
		log.Println("Found config.xml")
		return "config.xml", "xml", nil
	}

	if _, err := os.Stat("config.toml"); err == nil {
		log.Println("Found config.toml")
		return "config.toml", "toml", nil
	}

	return "", "", errors.New("no config file found")
}
