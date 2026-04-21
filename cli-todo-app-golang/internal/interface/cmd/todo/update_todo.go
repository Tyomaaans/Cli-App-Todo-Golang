package cmd

import (
    "github.com/spf13/cobra"
)

func NewUpdateTodoCommand(updatetodoFunc func(cmd *cobra.Command, args []string) error) *cobra.Command {
    cmd := &cobra.Command{
        Use: "update-todo",
        RunE: updatetodoFunc,
    }

    cmd.Flags().Int("update-index", 0, "Index Number of Todo Want to Update")
	cmd.Flags().String("task", "", "Task for Update")
	cmd.Flags().String("priority", "", "Priority for Update")
	cmd.Flags().String("progress", "", "Progress for Update")
	cmd.Flags().String("due-date", "", "DueDate for Update")

    return cmd
}