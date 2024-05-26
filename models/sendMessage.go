package models

type SendMessage struct {
	StatusCode int    `json:"status"`
	Message    string `json:"messaage"`
}
