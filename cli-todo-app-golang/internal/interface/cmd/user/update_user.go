package cmd

import (
    "github.com/spf13/cobra"
)

func NewUpdateUserCommand(updateuserFunc func(cmd *cobra.Command, args []string) error) *cobra.Command {
    cmd := &cobra.Command{
        Use: "update-user",
        RunE: updateuserFunc,
    }

    cmd.Flags().String("new-name", "", "New Name for Update")
	cmd.Flags().String("new-username", "", "New Username for Update")
	cmd.Flags().String("new-email", "", "New Email for Update")
	cmd.Flags().String("new-password", "", "New Passwrord for Update")
	cmd.Flags().String("confirm-password", "", "Confirm Password")

    return cmd
}