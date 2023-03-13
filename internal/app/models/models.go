package models

type Post struct {
	ID      string `json:"id,omitempty"`
	Title   string `json:"title,omitempty"`
	Content string `json:"content,omitempty"`
}

type Posts []Post

type User struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Nick string `json:"nick,omitempty"`

	Posts   Posts `json:"posts,omitempty"`
	Friends Users `json:"friends,omitempty"`
}

type Users []User

type Product struct {
	ID    string  `json:"id,omitempty"`
	Name  string  `json:"name,omitempty"`
	Type  string  `json:"type,omitempty"`
	Price float64 `json:"price,omitempty"`
	Sail  int     `json:"sail,omitempty"`
}

type Products []Product
