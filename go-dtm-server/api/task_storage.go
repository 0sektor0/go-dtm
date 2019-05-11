package api

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/0sektor0/go-dtm/go-dtm-server/models"
)

const (
	DEFAULT_TASK_TYPE        = 1
	TASK_FILTER_DEFAULT_TYPE = "default"
	TASK_FILTER_ASIGNEE_TYPE = "asignee"
	TASK_FILTER_CREATOR_TYPE = "creator"
	TASK_FILTER_STATUS_TYPE  = "status"
	TASK_FILTER_END_TYPE     = "end"
	TASK_FILTER_START_TYPE   = "start"
)

type ITaskStorage interface {
	FindById(id int) (*models.Task, error)

	GetList(offset int, limit int, filterType string, filterParam int) (*models.Tasks, error)

	Create(creatorId int, taskTypeId int, title string, text string) (*models.Task, error)

	Change(taskId int, task *models.Task) error

	Delete(id int) error

	CheckPermision(user *models.User, id int) bool
}

type TaskStorage struct {
	_db          *sql.DB
	_comments    ICommentStorage
	_attachments IAttachmentStorage
}

func NewTaskStorage(db *sql.DB, comments ICommentStorage, attachments IAttachmentStorage) *TaskStorage {
	storage := &TaskStorage{
		_db:          db,
		_comments:    comments,
		_attachments: attachments,
	}

	return storage
}

func (this *TaskStorage) FindById(id int) (*models.Task, error) {
	result := this._db.QueryRow(`SELECT t.id, t.title, t.task_text, t.creation_date, t.start_date, t.end_date, t.update_date,
	d.id, d.login, d.password, d.picture, d.is_admin, 
	s.id, s.login, s.password, s.picture, s.is_admin, 
	tt.id, tt.task_type_name, 
	ts.id, ts.task_status_name
	FROM task AS t 
	JOIN developer AS d ON t.creator_id = d.id
	JOIN developer AS s ON t.asignee_id = s.id
	JOIN task_type AS tt ON t.task_type_id = tt.id
	JOIN task_status AS ts ON t.task_status_id = ts.id
	WHERE t.id = $1;`,
		id,
	)

	task, err := ScanTask(result)
	if err != nil {
		return nil, err
	}

	err = this.GetTaskLinkedInfo(task)
	return task, err
}

func (this *TaskStorage) GetTaskLinkedInfo(task *models.Task) error {
	comments, err := this._comments.GetByTaskId(task.Id)
	task.Comments = comments
	if err != nil {
		return err
	}

	task.Attachments, err = this._attachments.GetByTaskId(task.Id)
	if err != nil {
		return err
	}

	return nil
}

func (this *TaskStorage) GetList(offset int, limit int, filterType string, filterParam int) (*models.Tasks, error) {
	where := ""

	switch filterType {
	case TASK_FILTER_DEFAULT_TYPE:
		break

	case TASK_FILTER_ASIGNEE_TYPE:
		where = `WHERE t.asignee_id=$3`

	case TASK_FILTER_CREATOR_TYPE:
		where = `WHERE t.creator_id=$3`

	case TASK_FILTER_END_TYPE:
		where = `WHERE t.end_date>=$3`

	case TASK_FILTER_START_TYPE:
		where = `WHERE t.start_date>=$3`

	case TASK_FILTER_STATUS_TYPE:
		where = `WHERE ts.id=$3`

	default:
		return nil, errors.New("unknown filter type")
	}

	sqlRequest := fmt.Sprintf(`SELECT t.id, t.title, t.task_text, t.creation_date, t.start_date, t.end_date, t.update_date,
	d.id, d.login, d.password, d.picture, d.is_admin, 
	s.id, s.login, s.password, s.picture, s.is_admin, 
	tt.id, tt.task_type_name, 
	ts.id, ts.task_status_name
	FROM task AS t 
	JOIN developer AS d ON t.creator_id = d.id
	JOIN developer AS s ON t.asignee_id = s.id
	JOIN task_type AS tt ON t.task_type_id = tt.id
	JOIN task_status AS ts ON t.task_status_id = ts.id
	%s ORDER BY t.id OFFSET $1 LIMIT $2;`, where)

	var rows *sql.Rows
	var err error 

	if where == "" {
		rows, err = this._db.Query(sqlRequest, offset, limit)
	} else {
		rows, err = this._db.Query(sqlRequest, offset, limit, filterParam)
	}

	if err != nil {
		return nil, err
	}

	tasks, err := ScanTasks(rows)
	if err != nil {
		return nil, err
	}

	for _, task := range tasks.Tasks {
		err = this.GetTaskLinkedInfo(task)
		if err != nil {
			return nil, err
		}
	}

	return tasks, nil
}

func (this *TaskStorage) Create(creatorId int, taskTypeId int, title string, text string) (*models.Task, error) {
	result := this._db.QueryRow(
		`WITH new_task AS (
			INSERT INTO task(creator_id, task_type_id, title, task_text, creation_date) 
			VALUES($1, $2, $3, $4, $5) 
			RETURNING id, creator_id, asignee_id, task_type_id, task_status_id, title, task_text, creation_date, start_date, end_date, update_date
		)
		
		SELECT t.id, t.title, t.task_text, t.creation_date, t.start_date, t.end_date, t.update_date,
		d.id, d.login, d.password, d.picture, d.is_admin, 
		s.id, s.login, s.password, s.picture, s.is_admin, 
		tt.id, tt.task_type_name, 
		ts.id, ts.task_status_name
		FROM new_task AS t 
		JOIN developer AS d ON t.creator_id = d.id
		JOIN developer AS s ON t.asignee_id = s.id
		JOIN task_type AS tt ON t.task_type_id = tt.id
		JOIN task_status AS ts ON t.task_status_id = ts.id;`,
		creatorId, taskTypeId, title, text, time.Now().Unix(),
	)

	return ScanTask(result)
}

func (this *TaskStorage) Change(taskId int, task *models.Task) error {
	now := time.Now().Unix()

	_, err := this._db.Exec(`UPDATE task SET 
	task_status_id=$1, task_type_id=$2, asignee_id=$3, task_text=$4, title=$5, start_date=$6, end_date=$7, update_date=$8 WHERE id = $9`,
		task.TaskStatus.Id, task.TaskType.Id, task.Asignee.Id, task.Text, task.Title, task.StartDate, task.EndDate, now, taskId)

	return err
}

func (this *TaskStorage) CheckPermision(user *models.User, id int) bool {
	if user.IsAdmin {
		return true
	}

	result := this._db.QueryRow("SELECT COUNT(id) FROM task WHERE creator_id=$1 AND id=$2;", user.Id, id)
	count, err := ScanCount(result)
	if err != nil || count == 0 {
		return false
	}

	return true
}

func (this *TaskStorage) Delete(id int) error {
	_, err := this._db.Exec("DELETE FROM  task WHERE id=$1", id)
	return err
}
