package cmd

import (
    "github.com/spf13/cobra"
)

func NewDeleteTodoCommand(deletetodoFunc func(cmd *cobra.Command, args []string) error) *cobra.Command {
    cmd := &cobra.Command{
        Use: "delete-todo",
        RunE: deletetodoFunc,
    }

    cmd.Flags().Int("delete-index", 0, "Index Number of Todo Want to Delete")

    return cmd
}