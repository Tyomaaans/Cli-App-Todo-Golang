package user

import (
    "context"
    "fmt"

    "github.com/spf13/cobra"

    "cli-todo-app-golang/internal/usecase/user"
    "cli-todo-app-golang/internal/interface/dto"
)

type UpdateUsernameHandler struct {
    updateusernameUserUC *user.UpdateUsernameUseCase
}

func NewUpdateUsernameHandler(updateusernameUC *user.UpdateUsernameUseCase) *UpdateUsernameHandler {
    return &UpdateUsernameHandler{
        updateusernameUserUC: updateusernameUC,
    }
}

func (h *UpdateUsernameHandler) UpdateUsername(cmd *cobra.Command, args []string) error {
    username, _ := cmd.Flags().GetString("new-username")

    req := dto.UpdateUsernameRequest{
        Username: username,
    }

    if err := h.updateusernameUserUC.Execute(context.Background(), req); err != nil {
        return err
    }

    fmt.Printf("Update Username Success!")

    return nil
}