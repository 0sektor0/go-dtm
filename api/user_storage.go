package api

import (
	"database/sql"

	models "github.com/0sektor0/go-dtm/models"
)

type UserUpdateDelegate func(*models.User)

type IUserStorage interface {
	Create(login string, password string) error

	Delete(id int) error

	FindByLogin(login string) *models.User

	FindById(id int) *models.User

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

func (this *UserStorage) Create(login string, password string) error {
	return nil
}

func (this *UserStorage) Delete(id int) error {
	return nil
}

func (this *UserStorage) FindByLogin(login string) *models.User {
	return nil
}

func (this *UserStorage) FindById(id int) *models.User {
	return nil
}

func (this *UserStorage) Subscribe(handler UserUpdateDelegate) {

}
