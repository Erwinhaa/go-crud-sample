package models

type Post struct {
	ID     int    `uri:"id" json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
	UserId int    `json:"user_id"`
	User   User
}

type PostReturn struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
	UserId int    `json:"user_id"`
}

type PostInput struct {
	ID    string `uri:"id" json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

type PostsPublicReturn struct {
	Status bool          `json:"status"`
	Data   []*PostReturn `json:"data"`
}

type PostPublicReturn struct {
	Status bool        `json:"status"`
	Data   *PostReturn `json:"data"`
}
