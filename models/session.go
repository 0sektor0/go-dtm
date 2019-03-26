package models

type Session struct {
	Token       string `json:"token"`
	User        *User  `json:"user"`
	ExpiredTime int64  `json:"expired"`
}
