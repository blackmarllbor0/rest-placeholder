package models

type Post struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type Posts []Post

type User struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Nick   string `json:"nick"`
	Avatar string `json:"avatar"`

	Posts   Posts `json:"posts"`
	Friends Users `json:"friends"`
}

type Users []User

type Product struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Type  string  `json:"type"`
	Price float64 `json:"price"`
	Sail  int     `json:"sail"`
}

type Products []Product
