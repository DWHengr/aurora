package alert

type AlertMessage struct {
	Annotations map[string]interface{} `json:"annotations"`
	Labels      map[string]interface{} `json:"labels"`
}
