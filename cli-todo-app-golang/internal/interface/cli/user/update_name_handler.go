package user

import (
    "context"
    "fmt"

    "github.com/spf13/cobra"

    "cli-todo-app-golang/internal/usecase/user"
    "cli-todo-app-golang/internal/interface/dto"
)

type UpdateNameHandler struct {
    updatenameUserUC *user.UpdateNameUseCase
}

func NewUpdateNameHandler(updatenameUC *user.UpdateNameUseCase) *UpdateNameHandler {
    return &UpdateNameHandler{
        updatenameUserUC: updatenameUC,
    }
}

func (h *UpdateNameHandler) UpdateName(cmd *cobra.Command, args []string) error {
    name, _ := cmd.Flags().GetString("new-name")

    req := dto.UpdateNameRequest{
        Name: name,
    }

    if err := h.updatenameUserUC.Execute(context.Background(), req); err != nil {
        return err
    }

    fmt.Printf("Update Name Success!")

    return nil
}