package models

type ChatMessage struct {
	Message   string `json:"message"`
	UserName  string `json:"userName"`
	UserEmail string `json:"userEmail"`
	Time      string `json:"time"`
}
