package models

type Task struct {
	Id           int	       `json:"id"`
	Creator      *User         `json:"creator,omitempty"`
	Asignee      *User         `json:"asignee,omitempty"`
	AsigneeId    *int          `json:"asigneeId,omitempty"`
	TaskType     *Type         `json:"type"`
	TaskTypeId   *int          `json:"typeId,omitempty"`
	TaskStatus   *Type         `json:"status"`
	TaskStatusId *int          `json:"statusId,omitempty"`
	Title        string        `json:"title"`
	Text         string        `json:"text,omitempty"`
	CreationDate int32         `json:"creationDate"`
	StartDate    int32         `json:"startDate"`
	EndDate      int32         `json:"endDate"`
	UpdateDate   int32         `json:"updDateTime"`
	Attachments  []*Attachment `json:"attachments"`
	Comments     []*Comment    `json:"comments"`
}

type Tasks struct {
	Tasks []*Task `json:"tasks"`
}
