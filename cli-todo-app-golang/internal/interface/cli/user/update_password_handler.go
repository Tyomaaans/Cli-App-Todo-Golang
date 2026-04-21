package user

import (
    "context"
    "fmt"

    "github.com/spf13/cobra"

    "cli-todo-app-golang/internal/usecase/user"
	"cli-todo-app-golang/internal/interface/dto"
)

type UpdatePasswordHandler struct {
    updatepasswordUserUC *user.UpdatePasswordUseCase
}

func NewUpdatePaswordHandler(updatepasswordUC *user.UpdatePasswordUseCase) *UpdatePasswordHandler {
    return &UpdatePasswordHandler{
        updatepasswordUserUC: updatepasswordUC,
    }
}

func (h *UpdatePasswordHandler) UpdatePassword(cmd *cobra.Command, args []string) error {
    password, _ := cmd.Flags().GetString("new-password")
	confirmpassword,_ := cmd.Flags().GetString("confirm-password")

	req := dto.UpdatePasswordRequest{
		Password:        password,
		ConfirmPassword: confirmpassword,
	}

    if err := h.updatepasswordUserUC.Execute(context.Background(), req); err != nil {
        return err
    }

    fmt.Printf("Update Password Success!")

    return nil
}