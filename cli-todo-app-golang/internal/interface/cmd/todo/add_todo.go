package cmd

import (
    "github.com/spf13/cobra"
)

func NewAddTodoCommand(addtodoFunc func(cmd *cobra.Command, args []string) error) *cobra.Command {
    cmd := &cobra.Command{
        Use: "add-todo",
        RunE: addtodoFunc,
    }

    cmd.Flags().String("task", "", "Task for New Todo")
	cmd.Flags().String("priority", "", "Priority for New Todo")
	cmd.Flags().String("due-date", "", "DueDate for New Todo")

    return cmd
}