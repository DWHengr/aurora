package internal

type AuroraConfig struct {
	Remote             *Remote `yaml:"remote"`
	PrometheusRulePath string  `yaml:"prometheusRulePath"`
	PrometheusUrl      string  `yaml:"prometheusUrl"`
	Username           string  `yaml:"username"`
	Password           string  `yaml:"password"`
}

type Remote struct {
	PrometheusHostIp          string "prometheusHostIp"
	PrometheusHostSshUsername string "prometheusHostSshUsername"
	PrometheusHostSshPassword string "prometheusHostSshPassword"
	PrometheusHostSshPort     int    "prometheusHostSshPort"
}
