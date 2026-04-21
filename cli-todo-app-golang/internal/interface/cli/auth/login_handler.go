package auth

import (
    "context"
    "fmt"

    "github.com/spf13/cobra"

    "cli-todo-app-golang/internal/usecase/auth"
    "cli-todo-app-golang/internal/interface/dto"
)

type LoginHandler struct {
    loginUserUC *auth.LoginUseCase
}

func NewLoginHandler(loginUC *auth.LoginUseCase) *LoginHandler {
    return &LoginHandler{
        loginUserUC: loginUC,
    }
}

func (h *LoginHandler) Login(cmd *cobra.Command, args []string) error {
    username, _ := cmd.Flags().GetString("username")
    email, _ := cmd.Flags().GetString("email")
    password, _ := cmd.Flags().GetString("password")

    req := dto.LoginRequest{
		Email:    email,
		Username: username,
		Password: password,
	}

    if err := h.loginUserUC.Execute(context.Background(), req); err != nil {
        return err
    }
    
    fmt.Printf("Login success! Welcome Back!")

    return nil
}