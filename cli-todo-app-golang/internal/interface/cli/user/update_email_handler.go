package user

import (
    "context"
    "fmt"

    "github.com/spf13/cobra"

    "cli-todo-app-golang/internal/usecase/user"
    "cli-todo-app-golang/internal/interface/dto"
)

type UpdateEmailHandler struct {
    updateemailUserUC *user.UpdateEmailUseCase
}

func NewUpdateEmailHandler(updateemailUC *user.UpdateEmailUseCase) *UpdateEmailHandler {
    return &UpdateEmailHandler{
        updateemailUserUC: updateemailUC,
    }
}

func (h *UpdateEmailHandler) UpdateEmail(cmd *cobra.Command, args []string) error {
    email, _ := cmd.Flags().GetString("new-email")

    req := dto.UpdateEmailRequest{
        Email: email,
    }

    if err := h.updateemailUserUC.Execute(context.Background(), req); err != nil {
        return err
    }

    fmt.Printf("Update Email Success!")

    return nil
}