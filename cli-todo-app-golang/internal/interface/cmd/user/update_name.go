package cmd

import (
    "github.com/spf13/cobra"
)

func NewUpdateNameCommand(updatenameFunc func(cmd *cobra.Command, args []string) error) *cobra.Command {
    cmd := &cobra.Command{
        Use: "update-name",
        RunE: updatenameFunc,
    }

    cmd.Flags().String("new-name", "", "New Name for Update")

    return cmd
}