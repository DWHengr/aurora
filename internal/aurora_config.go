package internal

type AuroraConfig struct {
	PrometheusRulePath string `yaml:"prometheusRulePath"`
	PrometheusUrl      string `yaml:"prometheusUrl"`
}
