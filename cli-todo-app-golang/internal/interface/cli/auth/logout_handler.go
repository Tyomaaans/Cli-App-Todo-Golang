package auth

import (
    "context"
    "fmt"

    "github.com/spf13/cobra"

    "cli-todo-app-golang/internal/usecase/auth"
)

type LogoutHandler struct {
    logoutUserUC *auth.LogoutUseCase
}

func NewLogoutHandler(logoutUC *auth.LogoutUseCase) *LogoutHandler {
    return &LogoutHandler{
        logoutUserUC: logoutUC,
    }
}

func (h *LogoutHandler) Logout(cmd *cobra.Command, args []string) error {
    if err := h.logoutUserUC.Execute(context.Background()); err != nil {
        return err
    }

    fmt.Printf("Logout Success! See You Again!")

    return nil
}