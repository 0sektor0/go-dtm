package models

type User struct {
	Id       int    `json: "id"`
	Login    string `json: "login"`
	Picture  string `json: "picture"`
	Password string
	IsAdmin  bool
}
