package localtypes

type Config struct {
	Targets []Target `yaml:"targets" json:"targets" xml:"targets" toml:"targets"`
}

type Target struct {
	Name    string            `yaml:"name" json:"name" xml:"name" toml:"name"`
	Address string            `yaml:"address" json:"address" xml:"address" toml:"address"`
	Labels  map[string]string `yaml:"labels" json:"labels" xml:"labels" toml:"labels"`
}
