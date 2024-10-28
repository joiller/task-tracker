package cmd

import "github.com/spf13/cobra"

var RootCmd = NewRootCmd()

func NewRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "task-tracker",
		Short: "A simple CLI task tracker",
	}
	return cmd
}
