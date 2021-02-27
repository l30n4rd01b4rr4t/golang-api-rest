package domain

type User struct {
	Id       string `json:"id"`
	Nickname string `json:"nickname"`
	UserType string `json:"user_type"`
}
