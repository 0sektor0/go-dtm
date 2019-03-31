package models

type Task struct {
	Id           int           `json:"id"`
	Creator      *User         `json:"creator,omitempty"`
	Asignee      *User         `json:"asignee,omitempty"`
	TaskType     *Type         `json:"type"`
	TaskStatus   *Type         `json:"status"`
	Title        string        `json:"title"`
	Text         string        `json:"text,omitempty"`
	CreationDate int32         `json:"creationDate"`
	EndDate      int32         `json:"endDate"`
	UpdateDate   int32         `json:"updateTime"`
	Attachments  []*Attachment `json:"attachments"`
}

type Tasks struct {
	Tasks []*Task `json:"tasks"`
}
