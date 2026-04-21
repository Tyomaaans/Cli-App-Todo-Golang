package dto

type RegisterRequest struct {
	Name            string `json:"name" validate:"required,alphaspaceunicode,min=3"`
	Email           string `json:"email" validate:"required,email"`
	Username        string `json:"username" validate:"required,username,min=4"`
	Password        string `json:"password" validate:"required,password,min=8"`
	ConfirmPassword string `validate:"required"`
}

type ForgotPasswordRequest struct {
	Email           string `json:"email" validate:"required,email"`
	Password        string `json:"password" validate:"required,password,min=8"`
	ConfirmPassword string `validate:"required"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"omitempty,email"`
	Username string `json:"username"`
	Password string `json:"password" validate:"required"`
}

type UpdatePasswordRequest struct {
	Password string `json:"password" validate:"required,password,min=8"`
	ConfirmPassword string `validate:"required"`
}

type UpdateNameRequest struct {
	Name string `json:"name" validate:"required,alphaspaceunicode,min=3"`
}

type UpdateEmailRequest struct {
	Email string `json:"email" validate:"required,email`
}

type UpdateUsernameRequest struct {
	Username string `json:"username" validate:"required,username,min=4"`
}

type AddTodoRequest struct {
    Task     string `json:"task"     validate:"required,alphaspaceunicode,min=3"`
    Priority string `json:"priority" validate:"required,oneof=high mid low"`
    DueDate  string `json:"due_date" validate:"required,datetime=2006-01-02,futuredate"`
}

type DeleteTodoRequest struct {
	DeleteIndex int `validate:"gte=1"`
}

type UpdateTodoRequest struct {
	UpdateIndex int    `validate:"gte=1"`
	Task        string `json:"task"     validate:"omitempty,alphaspaceunicode,min=3"`
	Priority    string `json:"priority" validate:"omitempty,oneof=high mid low"`
	Progress    string `json:"progress" validate:"omitempty,oneof=pending on-progress done"`
    DueDate     string `json:"due_date" validate:"omitempty,datetime=2006-01-02,futuredate"`
}

type UpdateUserRequest struct {
	Name            string `json:"name" validate:"omitempty,alphaspaceunicode,min=3"`
	Email           string `json:"email" validate:"omitempty,email"`
	Username        string `json:"username" validate:"omitempty,username,min=4"`
	Password        string `json:"password" validate:"omitempty,password,min=8"`
	ConfirmPassword string `validate:"omitempty"`
}