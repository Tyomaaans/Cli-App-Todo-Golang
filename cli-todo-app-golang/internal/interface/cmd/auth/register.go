package cmd

import (
    "github.com/spf13/cobra"
)

func NewRegisterCommand(registerFunc func(cmd *cobra.Command, args []string) error) *cobra.Command {
    cmd := &cobra.Command{
        Use: "register",
        RunE: registerFunc,
    }

	cmd.Flags().String("name", "", "Name for Register")
	cmd.Flags().String("email", "", "Email for Register")
    cmd.Flags().String("username", "", "Username for Register")
    cmd.Flags().String("password", "", "Passowrd for Register")
	cmd.Flags().String("confirm-password", "", "Confirm Password")

    return cmd
}