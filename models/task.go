package models

type Task struct {
	Id           int    `json:"id"`
	Creator      *User  `json:"creator"`
	CreatorId    int    `json:"creatorId"`
	Asignee      *User  `json:"asignee"`
	AsigneeId    int    `json:"asigneeId"`
	TaskType     *Type  `json:"type"`
	TaskTypeId   int    `json:"taskTypeId"`
	TaskStatus   *Type  `json:"status"`
	TaskStatusId int    `json:"taskStatusId"`
	Title        string `json:"title"`
	Text         string `json:"text"`
	CreationDate string `json:"creationDate"`
	EndDate      string `json:"endDate"`
	UpdateDate   string `json:"updateTime"`
}
