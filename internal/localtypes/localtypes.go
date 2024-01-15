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

// GlpiConfig represents the configuration for the GLPI module.
type GlpiConfig struct {
	Url      string `yaml:"url" json:"url" toml:"url"`
	Username string `yaml:"username" json:"username" toml:"username"`
	Password string `yaml:"password" json:"password" toml:"password"`
	ApiKey   string `yaml:"apikey" json:"apikey" toml:"apikey"`
}

// GlpiResponse represents the response from the GLPI API.
type GlpiResponse struct {
	TotalCount int `json:"totalcount"`
	Entities   []struct {
		ID          int    `json:"id"`
		Name        string `json:"name"`
		IPAddresses []struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"ipaddresses"`
	} `json:"entities"`
}

// PhpIPAMConfig represents the configuration for the phpIPAM module.
type PhpIPAMConfig struct {
	Url      string `yaml:"url" json:"url" toml:"url"`
	Username string `yaml:"username" json:"username" toml:"username"`
	Password string `yaml:"password" json:"password" toml:"password"`
}

// PhpIPAMResponse represents the response from the phpIPAM API.
type PhpIPAMResponse struct {
	Success string `json:"success"`
	Data    []struct {
		ID          string `json:"id"`
		Hostname    string `json:"hostname"`
		Description string `json:"description"`
		IP          string `json:"ip"`
	} `json:"data"`
}
