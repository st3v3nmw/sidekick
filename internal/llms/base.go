package llms

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type Command struct {
	Command    string `json:"command"`
	Reasoning  string `json:"reasoning"`
	Risk       int    `json:"risk"`
	Assessment string `json:"assessment"`
	Done       bool   `json:"done"`
}

type Provider interface {
	Complete(request string) (*Command, error)
}
