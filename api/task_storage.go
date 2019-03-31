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
	FindTaskById(id int) (*models.Task, error)
	FindTasks(offset int, limit int) (*models.Tasks, error)
	CreateTask(creatorId int, taskTypeId int, title string, text string) (*models.Task, error)
	UpdateTask(asigneeId int, taskTypeId int, title string, text string, startDate int32, endDate int32) (*models.Task, error)
	CanUserEditTask(user *models.User) bool
	DeleteTask(id int) error
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

func (this *TaskStorage) FindTaskById(id int) (*models.Task, error) {
	result := this._db.QueryRow(`SELECT t.id, t.title, t.task_text, t.creation_date, 
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

func (this *TaskStorage) CreateTask(creatorId int, taskTypeId int, title string, text string) (*models.Task, error) {
	result := this._db.QueryRow(
		`WITH new_task AS (
			INSERT INTO task(creator_id, task_type_id, title, task_text, creation_date) 
			VALUES($1, $2, $3, $4, $5) 
			RETURNING id, creator_id, asignee_id, task_type_id, task_status_id, title, task_text, creation_date
		)
		
		SELECT t.id, t.title, t.task_text, t.creation_date, 
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

func (this *TaskStorage) FindTasks(offset int, limit int) (*models.Tasks, error) {
	rows, err := this._db.Query(`SELECT t.id, t.title, t.task_text, t.creation_date, 
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

	tasks := &models.Tasks {
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

func (this *TaskStorage) UpdateTask(asigneeId int, taskTypeId int, title string, text string, startDate int32, endDate int32) (*models.Task, error) {
	return nil, nil
}

func (this *TaskStorage) CanUserEditTask(user *models.User) bool {
	return false
}

func (this *TaskStorage) DeleteTask(id int) error {
	return nil
}
