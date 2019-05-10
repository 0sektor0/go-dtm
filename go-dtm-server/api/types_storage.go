package api

import (
	"github.com/0sektor0/go-dtm/go-dtm-server/models"
	"database/sql"
	"errors"
	"fmt"
)

const (
	TASK_TYPE 		= "task_type"
	TASK_STATUS 	= "task_status"
	ATTACHMENT_TYPE = "attachment_type"
)

type ITypeStorage interface {
	GetTypes() *models.Types
	GetType(id int) (*models.Type, bool)
	Update(id int, name string) error
	Create(name string) error
	Delete(id int) error
	CheckPermision(user *models.User, id int) bool
}

type TypeStorage struct {
	_db 		*sql.DB
	_types		[]*models.Type
	_cache 		map[int]*models.Type
	_tableName 	string
}

func NewTypeStorage(typeTable string, db *sql.DB) *TypeStorage {
	storage := &TypeStorage{
		_db: db,
		_types: make([]*models.Type, 0),
		_cache: make(map[int]*models.Type),
		_tableName: typeTable,
	}

	storage.LoadDataToCache()
	return storage
}

func (this *TypeStorage) LoadDataToCache() {
	query := fmt.Sprintf("SELECT * FROM %s", this._tableName)
	rows, err := this._db.Query(query)

	if err != nil {
		panic(err)
	}

	for rows.Next() {
		dataType, err := ScanType(rows)		
		if err != nil {
			panic(err)	
		} else {
			this._cache[dataType.Id] = dataType
		}
	}

	this._types = make([]*models.Type, 0)
	for _, dataType := range(this._cache) {
		this._types = append(this._types, dataType)
	}
}

func (this *TypeStorage) GetTypes() *models.Types {
	types := &models.Types {
		Types: this._types,
	}
	
	return types
}

func (this *TypeStorage) GetType(id int) (*models.Type, bool) {
	dataType, ok := this._cache[id]
	return dataType, ok
}

func (this *TypeStorage) Update(id int, name string) error {
	query := fmt.Sprintf("UPDATE $1 SET %s_text=$2 WHERE id=$3", this._tableName)
	_, err := this._db.Exec(query, name, id)

	if err != nil {
		return err
	}

	dataType, ok := this._cache[id]
	if !ok {
		return errors.New("cache miss")
	}

	dataType.Name = name
	return nil
}

func (this *TypeStorage) Create(name string) error {
	query := fmt.Sprintf("INSERT INTO %s(%s_text) VALUES($1) RETURNING (id, %s_text)", this._tableName, this._tableName, this._tableName)
	row := this._db.QueryRow(query, name)

	dataType, err := ScanType(row)
	if err != nil {
		return err
	}

	this._cache[dataType.Id] = dataType
	this._types = append(this._types, dataType)
	return nil
}

func (this *TypeStorage) Delete(id int) error {	
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1", this._tableName)
	_, err := this._db.Exec(query, id)

	return err
}

func  (this *TypeStorage) CheckPermision(user *models.User, id int) bool {
	return user.IsAdmin
}