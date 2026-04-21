package cmd

import (
    "github.com/spf13/cobra"
)

func NewDeleteUserCommand(deleteuserFunc func(cmd *cobra.Command, args []string) error) *cobra.Command {
    cmd := &cobra.Command{
        Use: "delete-user",
        RunE: deleteuserFunc,
    }

    return cmd
}