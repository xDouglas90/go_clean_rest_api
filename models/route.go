package models

type Post struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

var Posts []Post

func init() {
	Posts = []Post{
		{ID: "1", Title: "Post 1", Content: "Content 1"},
		{ID: "2", Title: "Post 2", Content: "Content 2"},
		{ID: "3", Title: "Post 3", Content: "Content 3"},
	}
}
