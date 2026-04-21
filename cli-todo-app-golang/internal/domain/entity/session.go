package entity

import "time"

type Session struct {
	UserId    string
	Name      string
	Username  string
	Email     string
	LoggedIn  bool
	LoginTime time.Time
}
