package internal

type AuroraConfig struct {
	PrometheusRulePath string `yaml:"prometheusRulePath"`
	PrometheusUrl      string `yaml:"prometheusUrl"`
	Username           string `yaml:"username"`
	Password           string `yaml:"password"`
}
