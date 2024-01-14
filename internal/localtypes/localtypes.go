package localtypes

// Config represents the configuration for the application.
type Config struct {
	Modules   string      `yaml:"module" json:"module" toml:"module"`          // Modules represents the modules to be loaded.
	Port      string      `yaml:"port" json:"port" toml:"port"`                // Port represents the port number.
	Insecure  bool        `yaml:"insecure" json:"insecure" toml:"insecure"`    // Insecure indicates whether the connection is secure or not.
	Interval  string      `yaml:"interval" json:"interval" toml:"interval"`    // Interval represents the time interval.
	Endpoints []Endpoints `yaml:"endpoints" json:"endpoints" toml:"endpoints"` // Endpoints represents the list of endpoints.
}

// Endpoints represents the endpoints structure.
type Endpoints struct {
	Targets []string          `yaml:"targets" json:"targets" toml:"targets"` // Targets represents the list of target addresses.
	Labels  map[string]string `yaml:"labels" json:"labels" toml:"labels"`    // Labels represents the key-value pairs for labels.
}
