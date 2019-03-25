package api

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type ApiClient struct {
	Users    IUserStorage
	Sessions ISessionStorage
}

func NewApiClient() (*ApiClient, error) {
	settings, err := GetSettings()
	if err != nil {
		return nil, err
	}

	db, err := sql.Open(settings.Connector, settings.ConnectionString)
	if err != nil {
		return nil, err
	}

	users := NewUserStorage(db)
	sessions := NewSessionStorage(users)

	client := &ApiClient{
		Users:    users,
		Sessions: sessions,
	}

	return client, nil
}
