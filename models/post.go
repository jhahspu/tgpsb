package models

type Post struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Intro     string `json:"intro"`
	Stack     string `json:"stack"`
	Content   string `json:"content"`
	UserName  string `json:"user_name" db:"user_name"`
	UpdatedAt string `json:"updated_at" db:"updated_at"`
}

type PostPayload struct {
	Title   string `json:"title"`
	Intro   string `json:"intro"`
	Stack   string `json:"stack"`
	Content string `json:"content"`
}
