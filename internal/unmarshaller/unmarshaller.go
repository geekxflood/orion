package unmarshaller

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/geekxflood/orion/internal/localtypes"
	"gopkg.in/yaml.v2"
)

// UnmarshalConfig unmarshals a config file into a Config struct
// UnmarshalConfig reads a configuration file from the specified path and unmarshals it into a localtypes.Config struct based on the specified format.
// The supported formats are "yaml", "json", and "toml".
//
// Parameters:
//   - path: The path to the configuration file.
//   - format: The format of the configuration file ("yaml", "json", or "toml").
//
// Returns:
//   - localtypes.Config: The unmarshaled configuration.
//   - error: An error if the unmarshaling process fails.
//
// Example usage:
//
//	conf, err := UnmarshalConfig("/path/to/config.yaml", "yaml")
//	if err != nil {
//	  log.Fatal(err)
//	}
func UnmarshalConfig(path string, format string) (localtypes.Config, error) {
	var conf localtypes.Config

	// Read the config file into a byte array
	data, err := os.ReadFile(path)
	if err != nil {
		return conf, errors.New("failed to read config file: " + err.Error())
	}

	// Unmarshal the config file into a Config struct based on the format
	switch format {
	case "yaml":
		err := yaml.Unmarshal(data, &conf)
		if err != nil {
			return conf, errors.New("failed to unmarshal YAML: " + err.Error())
		}
	case "json":
		err := json.Unmarshal(data, &conf)
		if err != nil {
			return conf, errors.New("failed to unmarshal JSON: " + err.Error())
		}
	case "toml":
		err := toml.Unmarshal(data, &conf)
		if err != nil {
			return conf, errors.New("failed to unmarshal TOML: " + err.Error())
		}
	default:
		return conf, errors.New("unsupported format: " + format)
	}

	return conf, nil
}
