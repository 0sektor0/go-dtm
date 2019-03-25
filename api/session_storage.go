package api

import (
	models "github.com/0sektor0/go-dtm/models"
)

type ISessionStorage interface {
	LogIn(login string, password string) *models.Session

	LogOut(token string) error

	Authentificate(token string) *models.Session

	OnUserInfoUpdate(userInfo *models.User)
}

type SessionStorage struct {
	_users    IUserStorage
	_sessions map[string]*models.Session
}

func NewSessionStorage(users IUserStorage) *SessionStorage {
	sessions := &SessionStorage{
		_users:    users,
		_sessions: make(map[string]*models.Session),
	}

	users.Subscribe(sessions.OnUserInfoUpdate)

	return sessions
}

func (this *SessionStorage) LogIn(login string, password string) *models.Session {
	return nil
}

func (this *SessionStorage) LogOut(token string) error {
	return nil
}

func (this *SessionStorage) Authentificate(token string) *models.Session {
	return nil
}

func (this *SessionStorage) OnUserInfoUpdate(userInfo *models.User) {

}
