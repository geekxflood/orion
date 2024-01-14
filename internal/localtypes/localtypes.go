package localtypes

type Config struct {
	Modules  string   `yaml:"modules" json:"modules" toml:"modules"`
	Port     string   `yaml:"port" json:"port" toml:"port"`
	Insecure bool     `yaml:"insecure" json:"insecure" toml:"insecure"`
	Targets  []Target `yaml:"targets" json:"targets" toml:"targets"`
}

type Target struct {
	Targets []string          `yaml:"name" json:"name" toml:"name"`
	Labels  map[string]string `yaml:"labels" json:"labels" toml:"labels"`
}
