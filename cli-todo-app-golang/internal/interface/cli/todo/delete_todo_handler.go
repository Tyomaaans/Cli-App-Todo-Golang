package user

import (
    "context"
    "fmt"
	
    "github.com/spf13/cobra"

    "cli-todo-app-golang/internal/usecase/todo"
    "cli-todo-app-golang/internal/interface/dto"
)

type DeleteTodoHandler struct {
    deletetodoUserUC *todo.DeleteTodoUseCase
}

func NewDeleteTodoHandler(deletetodoUC *todo.DeleteTodoUseCase) *DeleteTodoHandler {
    return &DeleteTodoHandler{
        deletetodoUserUC: deletetodoUC,
    }
}

func (h *DeleteTodoHandler) DeleteTodo(cmd *cobra.Command, args []string) error {
    deleteindex, _ := cmd.Flags().GetInt("delete-index")

    req := dto.DeleteTodoRequest{
        DeleteIndex: deleteindex,
    }

    if err := h.deletetodoUserUC.Execute(context.Background(), req); err != nil {
        return err
    }

    fmt.Printf("Delete Todo Success!")

    return nil
}