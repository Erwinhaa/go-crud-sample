package models

type User struct {
	ID       int    `uri:"id" json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserCreateInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserUpdateInput struct {
	ID       int    `uri:"id" json:"id"`
	Username string `json:"username"`
}

type UserReturn struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
}

type UserWithPostReturn struct {
	UserID   int           `json:"user_id"`
	Username string        `json:"username"`
	Post     []*PostReturn `json:"post"`
}

type UserPublicReturn struct {
	Status bool `json:"status"`
	Data   *UserReturn
}

type UsersPublicReturn struct {
	Status bool `json:"status"`
	Data   []*UserReturn
}

type UserLoginInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserTokenReturn struct {
	Status bool   `json:"status"`
	Token  string `json:"token"`
}
