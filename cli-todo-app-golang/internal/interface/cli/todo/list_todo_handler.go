package user

import (
    "context"
    "fmt"
	"strings"
	"unicode/utf8"
	
    "github.com/spf13/cobra"

    "cli-todo-app-golang/internal/usecase/todo"
)

type ListTodoHandler struct {
    listtodoUserUC *todo.ListTodoUseCase
}

func NewListTodoHandler(listtodoUC *todo.ListTodoUseCase) *ListTodoHandler {
    return &ListTodoHandler{
        listtodoUserUC: listtodoUC,
    }
}

func (h *ListTodoHandler) ListTodo(cmd *cobra.Command, args []string) error {
    todos, err := h.listtodoUserUC.Execute(context.Background()) 
	if err != nil {
        return err
    }

    headers := []string{"Index", "ID", "Task", "Due Date", "Priority", "Progress", "Created At"}
	rows := [][]string{}
	for i, todo := range todos {
		Index := i + 1
		rows = append(rows, []string{
			fmt.Sprintf("%d", Index),
			Truncate(todo.ID, 15),
			todo.Task,
			todo.DueDate.Format("2006-01-02"),
			todo.Priority,
			todo.Progress,
			todo.CreatedAt.Format("2006-01-02"),
		})
	}

	printTable(headers, rows)

    return nil
}

func printTable(headers []string, rows [][]string) {
	colWidths := make([]int, len(headers))
	for i, h := range headers {
		colWidths[i] = len(h)
	}
	for _, row := range rows {
		for i, cell := range row {
			if i < len(colWidths) && len(cell) > colWidths[i] {
				colWidths[i] = len(cell)
			}
		}
	}

	separator := func() string {
		s := "+"
		for _, w := range colWidths {
			s += strings.Repeat("-", w+2) + "+"
		}
		return s
	}

	formatRow := func(cols []string) string {
		s := "|"
		for i, col := range cols {
			pad := colWidths[i] - len(col)
			s += " " + col + strings.Repeat(" ", pad) + " |"
		}
		return s
	}

	sep := separator()
	fmt.Println(sep)
	fmt.Println(formatRow(headers))
	fmt.Println(sep)
	for _, row := range rows {
		fmt.Println(formatRow(row))
	}
	fmt.Println(sep)
}

func Truncate(s string, maxLength int) string {
	if utf8.RuneCountInString(s) <= maxLength {
		return s
	}
	return string([]rune(s)[:maxLength]) + "..."
}