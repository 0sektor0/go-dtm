package models

type Comment struct {
	Id          int    `json: "id"`
	Developer   *User  `json: "developer"`
	DeveloperId int    `json: "developerId"`
	Task        *Task  `json: "task"`
	TaskId      int    `json: "taskId"`
	Text        string `json: "text"`
}
