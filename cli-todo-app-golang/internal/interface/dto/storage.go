package dto

import "time"

type UserStorage struct {
	UserId    string    `json:"user_id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
}

type UserLogStorage struct {
	LogId  string    `json:"log_id"`
	UserId string    `json:"user_id"`
	Action string    `json:"action"`
	Time   time.Time `json:"time"`
}

type SessionStorage struct {
	UserId    string    `json:"user_id"`
	Name      string    `json:"name"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	LoggedIn  bool      `json:"logged_in"`
	LoginTime time.Time `json:"login_time"`
}

type TodoStorage struct {
	UserId string       `json:"user_id"`
    Todos  []TodoItem   `json:"todos"`
}

type TodoItem struct {
    ID        string    `json:"id"`
    Task      string    `json:"task"`
    DueDate   time.Time `json:"due_date"`
    Priority  string    `json:"priority"`
    Progress  string    `json:"progress"`
    CreatedAt time.Time `json:"created_at"`
}