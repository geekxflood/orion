package localtypes

// Config represents the configuration for the application.
type Config struct {
	Port      string     `yaml:"port" json:"port" toml:"port"`                // Port represents the port number.
	Insecure  bool       `yaml:"insecure" json:"insecure" toml:"insecure"`    // Insecure indicates whether the connection is secure or not.
	Interval  string     `yaml:"interval" json:"interval" toml:"interval"`    // Interval represents the time interval.
	Endpoints []Endpoint `yaml:"endpoints" json:"endpoints" toml:"endpoints"` // Endpoints represents the list of endpoints.
}

// Endpoint represents the endpoints structure.
type Endpoint struct {
	Type          string            `yaml:"type" json:"type" toml:"type"`                                                             // Type represents the type of endpoint (e.g., rest, file).
	URL           string            `yaml:"url,omitempty" json:"url,omitempty" toml:"url,omitempty"`                                  // URL represents the URL for rest endpoints.
	Method        string            `yaml:"method,omitempty" json:"method,omitempty" toml:"method,omitempty"`                         // Method represents the HTTP method for rest endpoints.
	Headers       map[string]string `yaml:"headers,omitempty" json:"headers,omitempty" toml:"headers,omitempty"`                      // Headers represents the HTTP headers for rest endpoints.
	ParserRules   map[string]string `yaml:"parser_rules,omitempty" json:"parser_rules,omitempty" toml:"parser_rules,omitempty"`       // ParserRules represents parsing rules.
	Timeout       int               `yaml:"timeout,omitempty" json:"timeout,omitempty" toml:"timeout,omitempty"`                      // Timeout represents the timeout for the endpoint.
	RetryCount    int               `yaml:"retry_count,omitempty" json:"retry_count,omitempty" toml:"retry_count,omitempty"`          // RetryCount represents the number of retries for the endpoint.
	RetryInterval int               `yaml:"retry_interval,omitempty" json:"retry_interval,omitempty" toml:"retry_interval,omitempty"` // RetryInterval represents the interval between retries.
	FilePath      string            `yaml:"file_path,omitempty" json:"file_path,omitempty" toml:"file_path,omitempty"`                // FilePath represents the file path for file endpoints.
	FileFormat    string            `yaml:"file_format,omitempty" json:"file_format,omitempty" toml:"file_format,omitempty"`          // FileFormat represents the file format for file endpoints.
}
