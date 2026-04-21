package cmd

import (
    "github.com/spf13/cobra"
)

func NewListTodoCommand(listtodoFunc func(cmd *cobra.Command, args []string) error) *cobra.Command {
    cmd := &cobra.Command{
        Use: "list",
        RunE: listtodoFunc,
    }

    return cmd
}