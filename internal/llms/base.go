package llms

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type Command struct {
	Command string `json:"command"`
	Why     string `json:"why"`
	Risk    int    `json:"risk"`
	Done    bool   `json:"done"`
}

type Provider interface {
	Complete(request string) (*Command, error)
}
