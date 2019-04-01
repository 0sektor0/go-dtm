package api

import (
	"database/sql"

	"github.com/0sektor0/go-dtm/models"
)

type ICommentStorage interface {
	Add(taskId int, text string, userId int) error
	GetByTaskId(taskId int) ([]*models.Comment, error)
	Edit(id int, comment *models.Comment) error
	CanUserEdit(user *models.User, commentId int) bool
	Delete(id int) error
}

type CommentStorage struct {
	_db *sql.DB
}

func NewCommentStorage(db *sql.DB) *CommentStorage {
	storage := &CommentStorage{
		_db: db,
	}

	return storage
}

func (this *CommentStorage) Add(taskId int, text string, userId int) error {
	return nil
}

func (this *CommentStorage) GetByTaskId(taskId int) ([]*models.Comment, error) {
	return nil, nil
}

func (this *CommentStorage) Edit(id int, comment *models.Comment) error {
	return nil
}

func (this *CommentStorage) CanUserEdit(user *models.User, commentId int) bool {
	return false
}

func (this *CommentStorage) Delete(id int) error {
	return nil
}
