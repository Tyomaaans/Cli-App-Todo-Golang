package user

import (
    "context"
    "fmt"
	
    "github.com/spf13/cobra"

    "cli-todo-app-golang/internal/usecase/todo"
    "cli-todo-app-golang/internal/interface/dto"
)

type AddTodoHandler struct {
    addtodoUserUC *todo.AddTodoUseCase
}

func NewAddTodoHandler(addtodoUC *todo.AddTodoUseCase) *AddTodoHandler {
    return &AddTodoHandler{
        addtodoUserUC: addtodoUC,
    }
}

func (h *AddTodoHandler) AddTodo(cmd *cobra.Command, args []string) error {
    task, _ := cmd.Flags().GetString("task")
	priority, _ := cmd.Flags().GetString("priority")
	duedate, _ := cmd.Flags().GetString("due-date")

    req := dto.AddTodoRequest{
        Task:     task,
        Priority: priority,
        DueDate:  duedate,
    }

    if err := h.addtodoUserUC.Execute(context.Background(), req); err != nil {
        return err
    }

    fmt.Printf("Add Todo Success!")

    return nil
}