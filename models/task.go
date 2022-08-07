package Models

import (
	"time"
	"github.com/jinzhu/gorm"
)

// Task type model
type Task struct{
	gorm.Model
	Name string
	Description string
	Deadline time.Time
	Status string
	Finisihed *bool `gorm:"default:false" gorm:"not null"`
}


// {
// 	"id":99,
// 	"name": "Membuat CRUD",
// 	"description": "Membuat CRUD untuk Tabel users",
// 	"deadline": 2022-08-10,
// 	"status": "process,finished,pending"
// 	"is_finished" : true
// }