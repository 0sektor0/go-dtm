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