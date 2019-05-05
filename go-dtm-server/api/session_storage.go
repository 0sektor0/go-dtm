package api

import (
	"errors"
	"sync"

	"github.com/0sektor0/go-dtm/models"
)

type ISessionStorage interface {
	LogIn(login string, password string) (*models.Session, error)

	LogOut(token string) error

	Authentificate(token string) (*models.Session, error)

	OnUserInfoUpdate(userInfo *models.User)
}

type SessionStorage struct {
	sync.Mutex
	_users          IUserStorage
	_tokenSessions  map[string]*models.Session
	_userIdSessions map[int]*models.Session
}

func NewSessionStorage(users IUserStorage) *SessionStorage {
	sessions := &SessionStorage{
		_users:          users,
		_tokenSessions:  make(map[string]*models.Session),
		_userIdSessions: make(map[int]*models.Session),
	}

	users.Subscribe(sessions.OnUserInfoUpdate)

	return sessions
}

func (this *SessionStorage) LogIn(login string, password string) (*models.Session, error) {
	user, err := this._users.FindByLogin(login)
	if err != nil {
		return nil, err
	}

	if user.Password != password {
		err = errors.New("user dosent exist")
		return nil, err
	}

	session := models.NewSession(user)

	this.Lock()
	defer this.Unlock()

	this._tokenSessions[session.Token] = session
	this._userIdSessions[user.Id] = session

	return session, nil
}

func (this *SessionStorage) LogOut(token string) error {
	session, ok := this._tokenSessions[token]
	if !ok {
		return errors.New("session dosent exist")
	}

	this.Lock()
	defer this.Unlock()

	delete(this._tokenSessions, token)
	delete(this._userIdSessions, session.User.Id)

	return nil
}

func (this *SessionStorage) Authentificate(token string) (*models.Session, error) {
	session, ok := this._tokenSessions[token]
	if !ok {
		return nil, errors.New("session dosent exist")
	}

	if !session.IsAlive() {
		return nil, errors.New("session have expired")
	}

	return session, nil
}

func (this *SessionStorage) OnUserInfoUpdate(userInfo *models.User) {
	session, ok := this._userIdSessions[userInfo.Id]
	if !ok {
		return
	}

	session.User = userInfo
}
