package cmd

import (
	"github.com/spf13/cobra"
	"task-tracker/internal/task"
)

func init() {
	RootCmd.AddCommand(NewListCmd())
}

func NewListCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List tasks",
		RunE: func(cmd *cobra.Command, args []string) error {
			return RunListTaskCmd(args)
		},
	}
	return cmd
}

func RunListTaskCmd(args []string) error {
	if len(args) < 1 {
		return task.ListTasks(task.TaskStatusFromString("all"))
	}
	for _, status := range args {
		err := task.ListTasks(task.TaskStatusFromString(status))
		if err != nil {
			return err
		}
	}
	return nil
}
