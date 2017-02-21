package models

type Book struct {
	ID      string  `json:"id,omitempty"`
	Title   string  `json:"title,omitempty"`
	Content string  `json:"content,omitempty"`
	Author  *Author `json:"author,omitempty"`
}
