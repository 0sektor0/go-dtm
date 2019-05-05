package models

type Comment struct {
	Id          int  	`json:"id"`
	Developer   *User  	`json:"developer"`
	DeveloperId int  	`json:"developerId"`
	Text        string 	`json:"text"`
}
