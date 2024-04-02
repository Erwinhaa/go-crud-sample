package models

type Post struct {
	ID    int64  `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

type PostReturn struct {
	ID    int64  `json:"id"`
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
