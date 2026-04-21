package entity

import "time"

type Todo struct {
	ID        string      
	UserId    string
	Task      string
	DueDate   time.Time
	Priority  string
	Progress  string
	CreatedAt time.Time
}
