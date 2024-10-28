package cmd

import (
	"github.com/spf13/cobra"
	"task-tracker/internal/task"
)

func init() {
	RootCmd.AddCommand(NewAddCmd())
}

func NewAddCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add",
		Short: "Add a task",
		RunE: func(cmd *cobra.Command, args []string) error {
			return RunAddTaskCmd(args)
		},
	}
	return cmd
}

func RunAddTaskCmd(args []string) error {
	description := args[0]
	return task.AddTask(description)
}
