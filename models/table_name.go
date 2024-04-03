package models

func (u *User) TableName() string {
	return "user"
}

func (p *Post) TableName() string {
	return "post"
}
