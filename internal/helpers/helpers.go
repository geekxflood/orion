package helpers

import (
	"encoding/base64"
	"errors"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/geekxflood/orion/internal/localtypes"
	"github.com/geekxflood/orion/internal/unmarshaller"
)

// RunRESTClient runs the REST client with the provided endpoint.

func RunRESTClient(*localtypes.Module) {
	// Placeholder function for running REST client
}

func RunFileClient(*localtypes.Module) {
	// Placeholder function for running file client
}

// ReadConfig reads a config file and returns a Config struct.
// It looks for the config file in the current directory and unmarshals it into a Config struct.
// If the config file is specified, it checks if the file exists and determines its format based on the file extension.
// If no config file is specified, it looks for default config files in the current directory.
// Supported file formats are JSON, YAML, and TOML.
func ReadConfig(file string) (localtypes.Config, error) {
	path, format, err := LookupConfig(file)
	if err != nil {
		return localtypes.Config{}, err
	}

	conf, err := unmarshaller.UnmarshalConfig(path, format)
	if err != nil {
		return localtypes.Config{}, err
	}

	return conf, nil
}

// LookupConfig looks for a config file in the current directory.
// If a config file is specified, it checks if the file exists and determines its format based on the file extension.
// If no config file is specified, it looks for default config files in the current directory.
// Supported file formats are JSON, YAML, and TOML.
func LookupConfig(file string) (string, string, error) {
	if file != "" {
		log.Printf("Using config file: %s", file)
		if _, err := os.Stat(file); err == nil {
			ext := strings.ToLower(filepath.Ext(file))
			switch ext {
			case ".json":
				log.Println("Found JSON format")
				return file, "json", nil
			case ".yml", ".yaml":
				log.Println("Found YAML format")
				return file, "yaml", nil
			case ".toml":
				log.Println("Found TOML format")
				return file, "toml", nil
			}
		}
	} else {
		log.Println("No config file specified, looking for default config file")
		defaultFiles := []string{"config.yml", "config.yaml", "config.json", "config.toml"}
		for _, defaultFile := range defaultFiles {
			if _, err := os.Stat(defaultFile); err == nil {
				log.Printf("Found %s", defaultFile)
				ext := strings.ToLower(filepath.Ext(defaultFile))
				switch ext {
				case ".yml", ".yaml":
					return defaultFile, "yaml", nil
				case ".json":
					return defaultFile, "json", nil
				case ".toml":
					return defaultFile, "toml", nil
				}
			}
		}
	}

	return "", "", errors.New("no config file found")
}

// EncodeBase64 encodes a string to Base64.
func EncodeBase64(str string) string {
	encoded := base64.StdEncoding.EncodeToString([]byte(str))
	return encoded
}

// GetHomePath returns the home directory path.
func GetHomePath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return home, nil
}
