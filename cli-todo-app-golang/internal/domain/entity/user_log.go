package entity

import "time"

type UserLog struct {
	LogId  string 
	UserId string
	Action string
	Time   time.Time
}
