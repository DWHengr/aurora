package internal

type AuroraConfig struct {
	PrometheusHostIp          string "prometheusHostIp"
	PrometheusHostSshUsername string "prometheusHostSshUsername"
	PrometheusHostSshPassword string "prometheusHostSshPassword"
	PrometheusHostSshPort     int    "prometheusHostSshPort"
	PrometheusRulePath        string `yaml:"prometheusRulePath"`
	PrometheusUrl             string `yaml:"prometheusUrl"`
	Username                  string `yaml:"username"`
	Password                  string `yaml:"password"`
}
