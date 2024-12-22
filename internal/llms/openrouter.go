package llms

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

var (
	openRouterURL = "https://openrouter.ai/api/v1"
)

type OpenRouter struct {
	model    string
	apiKey   string
	messages []Message
}

var _ Provider = (*OpenRouter)(nil)

type openRouterCompletionsResponse struct {
	Choices []struct {
		Message Message `json:"message"`
	} `json:"choices"`
	Error *struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
}

func NewOpenRouter(model, apiKey string) *OpenRouter {
	return &OpenRouter{
		model:  model,
		apiKey: apiKey,
	}
}

func (o *OpenRouter) Complete(request string) (*Command, error) {
	o.messages = append(o.messages, Message{Role: "user", Content: request})
	payload := map[string]interface{}{
		"model":    o.model,
		"messages": o.messages,
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("cannot encode payload: %w", err)
	}

	url := openRouterURL + "/chat/completions"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("cannot create request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+o.apiKey)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Title", "Sidekick")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("cannot call openrouter: %w", err)
	}
	defer resp.Body.Close()

	var result openRouterCompletionsResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("cannot decode response: %w", err)
	}

	if result.Error != nil {
		return nil, fmt.Errorf("cannot call open router: %s", result.Error.Message)
	}

	content := result.Choices[0].Message.Content
	o.messages = append(o.messages, Message{Role: "assistant", Content: content})

	var cmd Command
	if err = json.Unmarshal([]byte(content), &cmd); err != nil {
		return nil, fmt.Errorf("cannot decode command %s: %w", content, err)
	}

	fmt.Printf("%#v\n", cmd)

	return &cmd, nil
}
