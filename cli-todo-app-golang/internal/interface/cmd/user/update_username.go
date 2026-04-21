package cmd

import (
    "github.com/spf13/cobra"
)

func NewUpdateUsernameCommand(updateusernameFunc func(cmd *cobra.Command, args []string) error) *cobra.Command {
    cmd := &cobra.Command{
        Use: "update-username",
        RunE: updateusernameFunc,
    }

    cmd.Flags().String("new-username", "", "New Username for Update")

    return cmd
}