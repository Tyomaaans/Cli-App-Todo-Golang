package cmd

import (
    "github.com/spf13/cobra"
)

func NewLogoutCommand(logoutFunc func(cmd *cobra.Command, args []string) error) *cobra.Command {
    cmd := &cobra.Command{
        Use: "logout",
        RunE: logoutFunc,
    }

    return cmd
}