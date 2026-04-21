package entity

import "time"

type User struct {
	UserId    string    
	Name      string    
	Email     string    
	Username  string    
	Password  string    
	CreatedAt time.Time 
}
