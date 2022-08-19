package alertcore

type AlertMessage struct {
	Name      string                 `json:"name"`
	UniqueId  string                 `json:"uniqueid"`
	Value     string                 `json:"value"`
	Summary   string                 `json:"Summary"`
	Attribute map[string]interface{} `json:"attribute"`
}
