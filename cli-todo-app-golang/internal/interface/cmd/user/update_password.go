package cmd

import (
    "github.com/spf13/cobra"
)

func NewUpdatePasswordCommand(updatepasswordFunc func(cmd *cobra.Command, args []string) error) *cobra.Command {
    cmd := &cobra.Command{
        Use: "update-password",
        RunE: updatepasswordFunc,
    }

    cmd.Flags().String("new-password", "", "New Password for Update")
	cmd.Flags().String("confirm-password", "", "Confirm Password")

    return cmd
}