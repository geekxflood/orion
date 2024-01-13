// main.go
package main

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"log"
	"os"

	"github.com/BurntSushi/toml"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Targets []Target `yaml:"targets" json:"targets" xml:"targets" toml:"targets"`
}

type Target struct {
	Name    string            `yaml:"name" json:"name" xml:"name" toml:"name"`
	Address string            `yaml:"address" json:"address" xml:"address" toml:"address"`
	Labels  map[string]string `yaml:"labels" json:"labels" xml:"labels" toml:"labels"`
}

func main() {
	// read local config
	config, err := ReadConfig()
	if err != nil {
		panic(err)
	}

	log.Println("Config read successfully")
	log.Printf("Config: %v\n", config)

}

func ReadConfig() (Config, error) {

	path, format, err := LookupConfig()
	if err != nil {
		return Config{}, err
	}

	conf, err := UnmarshalConfig(path, format)
	if err != nil {
		return Config{}, err
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

// UnmarshalConfig unmarshals a config file into a Config struct
func UnmarshalConfig(path string, format string) (Config, error) {
	var conf Config

	data, err := os.ReadFile(path)
	if err != nil {
		return conf, err
	}

	if format == "yaml" {
		err := yaml.Unmarshal([]byte(data), &conf)
		if err != nil {
			return conf, err
		}
	}

	if format == "json" {
		err := json.Unmarshal([]byte(data), &conf)
		if err != nil {
			return conf, err
		}
	}

	if format == "xml" {
		err := xml.Unmarshal([]byte(data), &conf)
		if err != nil {
			return conf, err
		}
	}

	if format == "toml" {
		err := toml.Unmarshal([]byte(data), &conf)
		if err != nil {
			return conf, err
		}
	}
	return conf, nil
}
