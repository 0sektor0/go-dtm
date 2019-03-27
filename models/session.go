package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

const (
	TOKEN_DEFAULT_TTL = 86400
)

type Session struct {
	Token       string `json:"token"`
	User        *User  `json:"user"`
	ExpiredTime int64  `json:"expired"`
}

func (this *Session) IsAlive() bool {
	return this.ExpiredTime > time.Now().Unix()
}

func NewSession(user *User) *Session {
	token := uuid.NewV4().String()

	session := &Session{
		Token:       token,
		User:        user,
		ExpiredTime: time.Now().Unix() + TOKEN_DEFAULT_TTL,
	}

	return session
}
