package api

import (
	"errors"
	"sync"
	"time"

	"github.com/0sektor0/go-dtm/models"
	uuid "github.com/satori/go.uuid"
)

const (
	TOKEN_DEFAULT_TTL = 86400
)

type ISessionStorage interface {
	LogIn(login string, password string) (*models.Session, error)

	LogOut(token string) error

	Authentificate(token string) (*models.Session, error)

	OnUserInfoUpdate(userInfo *models.User)
}

type SessionStorage struct {
	sync.Mutex
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

func (this *SessionStorage) LogIn(login string, password string) (*models.Session, error) {
	user, err := this._users.FindByLogin(login)
	if err != nil {
		return nil, err
	}

	if user.Password != password {
		err = errors.New("user dosent exist")
		return nil, err
	}

	token := uuid.NewV4().String()
	session := &models.Session{
		Token:       token,
		User:        user,
		ExpiredTime: time.Now().Unix() + TOKEN_DEFAULT_TTL,
	}

	this.Lock()
	defer this.Unlock()

	this._sessions[token] = session
	return session, nil
}

func (this *SessionStorage) LogOut(token string) error {
	_, ok := this._sessions[token]
	if !ok {
		return errors.New("session dosent exist")
	}

	this.Lock()
	defer this.Unlock()
	
	delete(this._sessions, token);
	return nil
}

func (this *SessionStorage) Authentificate(token string) (*models.Session, error) {
	session, ok := this._sessions[token]
	if !ok {
		return nil, errors.New("session dosent exist")
	}

	return session, nil
}

func (this *SessionStorage) OnUserInfoUpdate(userInfo *models.User) {

}
