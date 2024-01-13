package unmarshaller

import (
	"encoding/json"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/geekxflood/orion/internal/localtypes"
	"gopkg.in/yaml.v2"
)

// UnmarshalConfig unmarshals a config file into a Config struct
func UnmarshalConfig(path string, format string) (localtypes.Config, error) {
	var conf localtypes.Config

	// read the config file into a byte array
	data, err := os.ReadFile(path)
	if err != nil {
		return conf, err
	}

	// unmarshal the config file into a Config struct based on the format
	switch format {
	case "yaml":
		err := yaml.Unmarshal([]byte(data), &conf)
		if err != nil {
			return conf, err
		}
	case "json":
		err := json.Unmarshal([]byte(data), &conf)
		if err != nil {
			return conf, err
		}
	case "toml":
		err := toml.Unmarshal([]byte(data), &conf)
		if err != nil {
			return conf, err
		}
	}

	return conf, nil
}
