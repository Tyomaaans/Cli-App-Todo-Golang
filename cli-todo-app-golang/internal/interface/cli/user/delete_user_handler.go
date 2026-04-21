package user

import (
    "context"
    "fmt"

    "github.com/spf13/cobra"

    "cli-todo-app-golang/internal/usecase/user"
)

type DeleteUserHandler struct {
    deleteuserUserUC *user.DeleteUserUseCase
}

func NewDeleteUserlHandler(deleteuserUC *user.DeleteUserUseCase) *DeleteUserHandler {
    return &DeleteUserHandler{
        deleteuserUserUC: deleteuserUC,
    }
}

func (h *DeleteUserHandler) DeleteUser(cmd *cobra.Command, args []string) error {
    if err := h.deleteuserUserUC.Execute(context.Background()); err != nil {
        return err
    }

    fmt.Printf("Delete User Success!")

    return nil
}