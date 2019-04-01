package api

import (
	"database/sql"

	"github.com/0sektor0/go-dtm/models"
)

type IAttachmentStorage interface {
	Add(taskId int, text string, userId int) error
	GetByTaskId(taskId int) ([]*models.Attachment, error)
	Delete(id int) error
}

type AttachmentStorage struct {
	_db *sql.DB
}

func NewAttachmentStorage(db *sql.DB) *AttachmentStorage {
	storage := &AttachmentStorage{
		_db: db,
	}

	return storage
}

func (this *AttachmentStorage) Add(taskId int, text string, userId int) error {
	return nil
}

func (this *AttachmentStorage) GetByTaskId(taskId int) ([]*models.Attachment, error) {
	return nil, nil
}

func (this *AttachmentStorage) Delete(id int) error {
	return nil
}
