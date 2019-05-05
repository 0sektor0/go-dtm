package api

import (
	"database/sql"

	"github.com/0sektor0/go-dtm/models"
)

type UserUpdateDelegate func(*models.User)

type IUserStorage interface {
	Create(login string, password string) (*models.User, error)

	Delete(id int) error

	FindByLogin(login string) (*models.User, error) 

	FindById(id int) (*models.User, error) 

	Subscribe(handler UserUpdateDelegate)
}

type UserStorage struct {
	_db       *sql.DB
	_handlers []UserUpdateDelegate
}

func NewUserStorage(db *sql.DB) *UserStorage {
	storage := &UserStorage{
		_handlers: make([]UserUpdateDelegate, 0),
		_db: db,
	}

	return storage
}

func (this *UserStorage) Create(login string, password string) (*models.User, error) {
	result := this._db.QueryRow(
		"INSERT INTO developer(login, password) VALUES ($1,$2) RETURNING id, login, password, picture, is_admin;",
		login, password,
	)

	user, err := ScanUser(result)
	return user, err
}

func (this *UserStorage) Delete(id int) error {
	return nil
}

func (this *UserStorage) FindByLogin(login string) (*models.User, error) {
	result := this._db.QueryRow(
		"SELECT id, login, password, picture, is_admin FROM developer WHERE login=$1;",
		login, 
	)

	user, err := ScanUser(result)
	return user, err
}

func (this *UserStorage) FindById(id int) (*models.User, error) {
	result := this._db.QueryRow(
		"SELECT id, login, password, picture, is_admin FROM developer WHERE id=$1;",
		id, 
	)

	user, err := ScanUser(result)
	return user, err
}

func (this *UserStorage) Update(user *models.User) error {
	result := this._db.QueryRow(
		"UPDATE developer SET login=$1, password=$2, picture=$3, is_admin=$4 WHERE id=$5 RETURNING id, login, password, picture, is_admin;",
		user.Login, user.Password, user.Picture, user.IsAdmin, user.Id, 
	)

	updatedUser, err := ScanUser(result)
	if err != nil {
		return err
	}

	this.OnNext(updatedUser)
	return nil
}

func (this *UserStorage) Subscribe(handler UserUpdateDelegate) {
	this._handlers = append(this._handlers, handler)
}

func (this *UserStorage) OnNext(userInfo *models.User) {
	for _, handler := range this._handlers {
		handler(userInfo)
	}
}