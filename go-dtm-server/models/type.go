package models

import (
	"database/sql"
	"fmt"
)

const (
	TASK_TYPE   = "task_type"
	TASK_STATUS = "task_status"
)

type Type struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	_db    *sql.DB
	_table string
}

func GetType(db *sql.DB, id int, table string) (*Type, error) {
	activeType := &Type{
		_db:    db,
		_table: table,
	}

	sql := fmt.Sprintf("SELECT id, %s_name FROM %s WHERE id = $1", table)
	result := db.QueryRow(sql, id)

	err := result.Scan(&activeType.Id, &activeType.Name)
	return activeType, err
}

func (this *Type) Save() error {
	sql := fmt.Sprintf("UPDATE %s SET %s_name = $1 WHERE id = $2", this._table)
	_, err := this._db.Exec(sql, this.Name, this.Id)

	return err
}

func (this *Type) Delete() error {
	sql := fmt.Sprintf("DELETE FROM  %s WHERE id = $1", this._table)
	_, err := this._db.Exec(sql, this.Id)

	return err
}
