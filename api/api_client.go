package api

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type ApiClient struct {
	Users       IUserStorage
	Sessions    ISessionStorage
	Tasks       ITaskStorage
	Comments    ICommentStorage
	Attachments IAttachmentStorage
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

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	users := NewUserStorage(db)
	sessions := NewSessionStorage(users)
	attachments := NewAttachmentStorage(db)
	comments := NewCommentStorage(db)
	tasks := NewTaskStorage(db, comments, attachments)

	client := &ApiClient{
		Users:       users,
		Sessions:    sessions,
		Tasks:       tasks,
		Comments:    comments,
		Attachments: attachments,
	}

	return client, nil
}
