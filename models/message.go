package models

type Message struct {
	ID      int    `json:"id"`
	Content string `json:"content"`
	Image   string `json:"image"`
}
