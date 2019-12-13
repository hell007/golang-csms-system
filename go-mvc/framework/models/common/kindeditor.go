package common

type Kindeditor struct {
	Error     int    `json:"error"`
	Url       string `json:"url,omitempty"`
	Message   string `json:"message,omitempty"`
}
