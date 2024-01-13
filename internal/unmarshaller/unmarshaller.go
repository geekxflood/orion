package unmarshaller

import (
	"encoding/json"
	"encoding/xml"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/geekxflood/orion/internal/localtypes"
	"gopkg.in/yaml.v2"
)

// UnmarshalConfig unmarshals a config file into a Config struct
func UnmarshalConfig(path string, format string) (localtypes.Config, error) {
	var conf localtypes.Config

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
