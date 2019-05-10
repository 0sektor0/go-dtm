package api

import (
	"database/sql"
	"time"

	"github.com/0sektor0/go-dtm/go-dtm-server/models"
)

type IAttachmentStorage interface {
	Add(taskId int, path string) error
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

func (this *AttachmentStorage) Add(taskId int, path string) error {
	now := time.Now().Unix()
	_, err := this._db.Exec(`INSERT INTO attachment(task_id, attachment_path, creation_date) VALUES($1, $2, $3)`,
	taskId, path, now)

	return err
}

func (this *AttachmentStorage) GetByTaskId(taskId int) ([]*models.Attachment, error) {
	rows, err := this._db.Query( `SELECT attachment_path, creation_date
		FROM attachment
		WHERE task_id=$1
		`,
		taskId,
	)
	
	if err != nil {
		return nil, err
	}

	return ScanAttachments(rows)
}

func (this *AttachmentStorage) Delete(id int) error {
	_, err := this._db.Exec("DELETE FROM attachment WHERE id=$1", id)
	return err
}
