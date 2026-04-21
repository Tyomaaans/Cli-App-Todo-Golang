package user

import (
    "context"
    "fmt"

    "github.com/spf13/cobra"

    "cli-todo-app-golang/internal/usecase/user"
    "cli-todo-app-golang/internal/interface/dto"
)

type UpdateUserHandler struct {
    updateuserUserUC *user.UpdateUserUseCase
}

func NewUpdateUserHandler(updateuserUC *user.UpdateUserUseCase) *UpdateUserHandler {
    return &UpdateUserHandler{
        updateuserUserUC: updateuserUC,
    }
}

func (h *UpdateUserHandler) UpdateUser(cmd *cobra.Command, args []string) error {
	name, _ := cmd.Flags().GetString("new-name")
	username, _ := cmd.Flags().GetString("new-username")
	email, _ := cmd.Flags().GetString("new-email")
	password, _ := cmd.Flags().GetString("new-password")
    confirmpassword, _ := cmd.Flags().GetString("confirm-password")

    req := dto.UpdateUserRequest{
		Name:  name,
		Username: username,
        Email: email,
		Password: password,
		ConfirmPassword: confirmpassword,
    }

    if err := h.updateuserUserUC.Execute(context.Background(), req); err != nil {
        return err
    }

    fmt.Printf("Update User Success!")

    return nil
}