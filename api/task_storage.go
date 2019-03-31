package api

import (
	"database/sql"
	"time"

	"github.com/0sektor0/go-dtm/models"
)

const (
	DEFAULT_TASK_TYPE = 1
)

type ITaskStorage interface {
	FindById(id int) (*models.Task, error)
	GetList(offset int, limit int) (*models.Tasks, error)
	Create(creatorId int, taskTypeId int, title string, text string) (*models.Task, error)
	ChangeTask(taskId int, task *models.Task) error
	CanUserEditTask(user *models.User, taskId int) bool
	Delete(id int) error
}

type TaskStorage struct {
	_db *sql.DB
}

func NewTaskStorage(db *sql.DB) *TaskStorage {
	storage := &TaskStorage{
		_db: db,
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

	return ScanTask(result)
}

func (this *TaskStorage) GetList(offset int, limit int) (*models.Tasks, error) {
	rows, err := this._db.Query(`SELECT t.id, t.title, t.task_text, t.creation_date, t.start_date, t.end_date, t.update_date,
	d.id, d.login, d.password, d.picture, d.is_admin, 
	s.id, s.login, s.password, s.picture, s.is_admin, 
	tt.id, tt.task_type_name, 
	ts.id, ts.task_status_name
	FROM task AS t 
	JOIN developer AS d ON t.creator_id = d.id
	JOIN developer AS s ON t.asignee_id = s.id
	JOIN task_type AS tt ON t.task_type_id = tt.id
	JOIN task_status AS ts ON t.task_status_id = ts.id
	OFFSET $1
	LIMIT $2;`,
		offset, limit,
	)

	if err != nil {
		return nil, err
	}

	tasks := &models.Tasks{
		Tasks: make([]*models.Task, 0),
	}

	for rows.Next() {
		task, err := ScanTask(rows)
		if err != nil {
			return nil, err
		}

		tasks.Tasks = append(tasks.Tasks, task)
	}

	return tasks, err
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

func (this *TaskStorage) ChangeTask(taskId int, task *models.Task) error {
	now := time.Now().Unix()

	_, err := this._db.Exec(`UPDATE task SET 
	task_status_id=$1 task_type_id=$2 asignee_id=$3 task_text=$4 title=$5 start_date=$6 end_date=$7 update_date=$8`,
		task.TaskStatusId, task.TaskTypeId, task.AsigneeId, task.Text, task.Title, task.StartDate, task.EndDate, now)

	return err
}

func (this *TaskStorage) CanUserEditTask(user *models.User, taskId int) bool {
	if user.IsAdmin {
		return true
	}

	_, err := this._db.Exec("SELECT id FROM task WHERE creator_id=$1 AND id=$2;", user.Id, taskId)
	if err != nil {
		return false
	}

	return true
}

func (this *TaskStorage) Delete(id int) error {
	_, err := this._db.Exec("DELETE task WHERE id=$1", id)
	return err
}
