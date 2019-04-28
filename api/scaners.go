package api

import (
	"github.com/0sektor0/go-dtm/models"
	"database/sql"
)

type IRow interface {
	Scan(dest ...interface{}) error
}

func ScanUser(row IRow) (*models.User, error) {
	user := new(models.User)
	err := row.Scan(&user.Id, &user.Login, &user.Password, &user.Picture, &user.IsAdmin)

	return user, err
}

func ScanTask(row IRow) (*models.Task, error) {
	task := new(models.Task)
	task.Creator = &models.User{}
	task.Asignee = &models.User{}
	task.TaskStatus = &models.Type{}
	task.TaskType = &models.Type{}

	err := row.Scan(&task.Id, &task.Title, &task.Text, &task.CreationDate, &task.StartDate, &task.EndDate, &task.UpdateDate,
		&task.Creator.Id, &task.Creator.Login, &task.Creator.Password, &task.Creator.Picture, &task.Creator.IsAdmin,
		&task.Asignee.Id, &task.Asignee.Login, &task.Asignee.Password, &task.Asignee.Picture, &task.Asignee.IsAdmin,
		&task.TaskType.Id, &task.TaskType.Name,
		&task.TaskStatus.Id, &task.TaskStatus.Name,
	)

	return task, err
}

func ScanTasks(rows *sql.Rows) (*models.Tasks, error) {
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

	return tasks, nil
}

func ScanComment(row IRow) (*models.Comment, error) {
	comment := new(models.Comment)
	comment.Developer = &models.User{}

	err := row.Scan(&comment.Id, &comment.DeveloperId, &comment.Text, 
		&comment.Developer.Id, &comment.Developer.Login, &comment.Developer.Password, &comment.Developer.Picture, &comment.Developer.IsAdmin)

	return comment, err
}

func ScanComments(rows *sql.Rows) ([]*models.Comment, error) {
	comments :=  make([]*models.Comment, 0)

	for rows.Next() {
		comment, err := ScanComment(rows)
		if err != nil {
			return nil, err
		}

		comments = append(comments, comment)
	}

	return comments, nil
}

func ScanAttachment(row IRow) (*models.Attachment, error) {
	attachment := &models.Attachment{}
	err := row.Scan(&attachment.Path, attachment.CreationDate)
	return attachment, err
}

func ScanAttachments(rows *sql.Rows) ([]*models.Attachment, error) {
	attachments :=  make([]*models.Attachment, 0)

	for rows.Next() {
		attachment, err := ScanAttachment(rows)
		if err != nil {
			return nil, err
		}

		attachments = append(attachments, attachment)
	}

	return attachments, nil
}

func ScanType(row IRow) (*models.Type, error) {
	dataType := &models.Type{}
	err := row.Scan(&dataType.Id, &dataType.Name)

	return dataType, err
}