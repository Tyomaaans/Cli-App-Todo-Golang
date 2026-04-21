package auth

import (
    "context"
    "fmt"

    "github.com/spf13/cobra"

    "cli-todo-app-golang/internal/usecase/auth"
    "cli-todo-app-golang/internal/interface/dto"
)

type RegisterHandler struct {
    registerUserUC *auth.RegisterUseCase
}

func NewRegisterHandler(registerUC *auth.RegisterUseCase) *RegisterHandler {
    return &RegisterHandler{
        registerUserUC: registerUC,
    }
}

func (h *RegisterHandler) Register(cmd *cobra.Command, args []string) error {
	name, _ := cmd.Flags().GetString("name")
	email, _ := cmd.Flags().GetString("email")
    username, _ := cmd.Flags().GetString("username")
    password, _ := cmd.Flags().GetString("password")
	confirmpassword, _ := cmd.Flags().GetString("confirm-password")

    req := dto.RegisterRequest{
		Name:            name,
		Email:           email,
		Username:        username,
		Password:        password,
		ConfirmPassword: confirmpassword,
	}

	if err := h.registerUserUC.Execute(context.Background(), req); err != nil {
		return err
	}

    fmt.Printf("Register success! User, %s\n", username)

    return nil
}