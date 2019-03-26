package models

type User struct {
	Id       int    `json:"id"`
	Login    string `json:"login"`
	Picture  string `json:"picture"`
	IsAdmin  bool   `json:"isAdmin"`
	Password string `json:"-"`
}
