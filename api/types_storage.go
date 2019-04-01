package api

import (
	"github.com/0sektor0/go-dtm/models"
)

type ITypeStorage interface {
	GetTypes() []*models.Type
	GetStatuses() []*models.Type
	GetType(id int) *models.Type
	GetStatuse(id int) *models.Type
}
