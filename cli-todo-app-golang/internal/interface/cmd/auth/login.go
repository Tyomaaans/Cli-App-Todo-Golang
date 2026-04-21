package cmd

import (
    "github.com/spf13/cobra"
)

func NewLoginCommand(loginFunc func(cmd *cobra.Command, args []string) error) *cobra.Command {
    cmd := &cobra.Command{
        Use: "login",
        RunE: loginFunc,
    }

    cmd.Flags().String("username", "", "Username for Login")
    cmd.Flags().String("email", "", "Email for Login")
    cmd.Flags().String("password", "", "Password for Login")

    return cmd
}