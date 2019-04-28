package models

type Attachment struct {
	Id           int	`json:"id"`
	Task         *Task  `json:"task"`
	TaskId       int  	`json:"taskId"`
	Type         *Type  `json:"type"`
	TypeId       int  	`json:"typeId"`
	Path         string `json:"path"`
	CreationDate int32  `json:"creationDate"`
}
