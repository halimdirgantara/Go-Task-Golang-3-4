package models

import "time"

type TaskModel struct {
	Id          int
	No			int
	Name        string
	Description string
	Employee    string
	Deadline    time.Time
	Status		string
	BtnStatus	string
	Completed 	int
	CreatedAt 	time.Time
	UpdatedAt 	time.Time
}