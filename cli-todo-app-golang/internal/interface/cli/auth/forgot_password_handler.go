package auth

import (
    "context"
    "fmt"

    "github.com/spf13/cobra"

    "cli-todo-app-golang/internal/usecase/auth"
    "cli-todo-app-golang/internal/interface/dto"
)

type ForgotPasswordHandler struct {
    forgotpasswordUserUC *auth.ForgotPasswordUseCase
}

func NewForgotPasswordHandler(forgotpasswordUC *auth.ForgotPasswordUseCase) *ForgotPasswordHandler {
    return &ForgotPasswordHandler{
        forgotpasswordUserUC: forgotpasswordUC,
    }
}

func (h *ForgotPasswordHandler) ForgotPassword(cmd *cobra.Command, args []string) error {
    email, _ := cmd.Flags().GetString("email")
    password, _ := cmd.Flags().GetString("new-password")
	confirmpassword, _ := cmd.Flags().GetString("confirm-password")

    req := dto.ForgotPasswordRequest{
		Email:           email,
		Password:        password,
		ConfirmPassword: confirmpassword,
	}

    if err := h.forgotpasswordUserUC.Execute(context.Background(), req); err != nil {
        return err
    }

    fmt.Printf("Change Password Success!",)

    return nil
}