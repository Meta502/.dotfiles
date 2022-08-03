package gorofiytm

type Line struct {
	Text      string `json:"text"`
	Urgent    bool   `json:"urgent"`
	Highlight bool   `json:"highlight"`
	Markup    bool   `json:"markup"`
	Icon      string `json:"icon"`
	Data      string `json:"data"`
}

type Output struct {
	Message     string `json:"message,omitempty"`
	Overlay     string `json:"overlay,omitempty"`
	Prompt      string `json:"prompt,omitempty"`
	Input       string `json:"input,omitempty"`
	InputAction string `json:"input action,omitempty"`
	EventFormat string `json:"event format,omitempty"`
	ActiveEntry int    `json:"active entry,omitempty"`
	Lines       []Line `json:"lines"`
	Value string `json:"value"`
}

type Input struct {
	Name  string `json:"name"`
	Value string `json:"value"`
	Data  string `json:"data"`
}
