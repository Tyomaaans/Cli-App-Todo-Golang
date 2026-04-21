package cmd

import (
    "github.com/spf13/cobra"
)

func NewUpdateEmailCommand(updateemailFunc func(cmd *cobra.Command, args []string) error) *cobra.Command {
    cmd := &cobra.Command{
        Use: "update-email",
        RunE: updateemailFunc,
    }

    cmd.Flags().String("new-email", "", "New Email for Update")

    return cmd
}   