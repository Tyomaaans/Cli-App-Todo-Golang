package user

import (
    "context"
    "fmt"
	
    "github.com/spf13/cobra"

    "cli-todo-app-golang/internal/usecase/todo"
    "cli-todo-app-golang/internal/interface/dto"
)

type UpdateTodoHandler struct {
    updatetodoUserUC *todo.UpdateTodoUseCase
}

func NewUpdateTodoHandler(updatetodoUC *todo.UpdateTodoUseCase) *UpdateTodoHandler {
    return &UpdateTodoHandler{
        updatetodoUserUC: updatetodoUC,
    }
}

func (h *UpdateTodoHandler) UpdateTodo(cmd *cobra.Command, args []string) error {
	updateindex, _ := cmd.Flags().GetInt("update-index")
    task, _ := cmd.Flags().GetString("task")
	priority, _ := cmd.Flags().GetString("priority")
	progress, _ := cmd.Flags().GetString("progress")
	duedate, _ := cmd.Flags().GetString("due-date")

    req := dto.UpdateTodoRequest{
		UpdateIndex: updateindex,
        Task:        task,
        Priority:    priority,
		Progress:    progress,
        DueDate:     duedate,
    }

    if err := h.updatetodoUserUC.Execute(context.Background(), req); err != nil {
        return err
    }

    fmt.Printf("Update Todo Success!")

    return nil
}