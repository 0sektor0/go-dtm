package api

import (
	"database/sql"

	"github.com/0sektor0/go-dtm/models"
)

type ICommentStorage interface {
	Add(taskId int, text string, userId int) error
	
	GetByTaskId(taskId int) ([]*models.Comment, error)
	
	Edit(id int, text string) error
	
	Delete(id int) error
	
	CheckPermision(user *models.User, id int) bool
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
	_, err := this._db.Exec(`INSERT INTO comment(developer_id, task_id, comment_text) VALUES($1, $2, $3)`,
		userId, taskId, text)

	return err
}

func (this *CommentStorage) GetByTaskId(taskId int) ([]*models.Comment, error) {
	rows, err := this._db.Query( `WITH task_comments AS (
		SELECT * 
		FROM comment
		WHERE task_id = $1
	)
	
	SELECT c.id, c.developer_id, c.comment_text, d.id, d.login, d.password, d.picture, d.is_admin
	FROM task_comments AS c
	JOIN developer AS d ON c.developer_id = d.id 
	`,
	taskId,
	)
	
	if err != nil {
		return nil, err
	}

	return ScanComments(rows)
}

func (this *CommentStorage) Edit(id int, text string) error {
	_, err := this._db.Exec(`UPDATE comment SET comment_text=$1`, text)
	return err
}

func (this *CommentStorage) CheckPermision(user *models.User, id int) bool {
	if user.IsAdmin {
		return true
	}

	result := this._db.QueryRow("SELECT COUNT(id) FROM comment WHERE developer_id=$1 AND id=$2;", user.Id, id)
	count, err := ScanCount(result)
	if err != nil || count == 0 {
		return false
	}

	return true
}

func (this *CommentStorage) Delete(id int) error {
	_, err := this._db.Exec("DELETE FROM comment WHERE id=$1", id)
	return err
}
