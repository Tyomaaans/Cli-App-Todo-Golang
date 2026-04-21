package cmd

import (
    "github.com/spf13/cobra"
)

func NewForgotPasswordCommand(forgotpasswordFunc func(cmd *cobra.Command, args []string) error) *cobra.Command {
    cmd := &cobra.Command{
        Use: "forgot-password",
        RunE: forgotpasswordFunc,
    }

    cmd.Flags().String("email", "", "Email for Forgot Password")
    cmd.Flags().String("new-password", "", "New Password for Update")
	cmd.Flags().String("confirm-password", "", "Confirm Password")

    return cmd
}