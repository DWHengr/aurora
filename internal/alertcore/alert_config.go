package alertcore

type AlertConfig struct {
	Thread int `yaml:"thread"`
	Buffer int `yaml:"buffer"`
}
