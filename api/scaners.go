package api

import (
	"github.com/0sektor0/go-dtm/models"
)

type IRow interface {
	Scan(dest ...interface{}) error
}

func ScanUser(row IRow) (*models.User, error) {
	user := new(models.User)
	err := row.Scan(&user.Id, &user.Login, &user.Password, &user.Picture, &user.IsAdmin)

	return user, err
}

func ScanType(row IRow) (*models.Type, error) {
	scannedType := new(models.Type)
	err := row.Scan(&scannedType.Id, &scannedType.Name)

	return scannedType, err
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
