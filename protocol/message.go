package protocol

type Message struct {
	Type string `json:"type"`
	Name string `json:"name"`
	IP   string `json:"ip"`
	Port int    `json:"port"`
}
