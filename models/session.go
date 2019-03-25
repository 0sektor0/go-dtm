package models

type Session struct {
	Token      string `json: "token"`
	TimeToLive int    `json: "ttl"`
	User       *User  `json: "user"`
}
