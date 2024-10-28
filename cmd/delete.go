package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"strconv"
	"task-tracker/internal/task"
)

func init() {
	RootCmd.AddCommand(NewDeleteCmd())
}

func NewDeleteCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "delete",
		Short: "Delete a task",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return RunDeleteTaskCmd(args)
		},
	}
}

func RunDeleteTaskCmd(args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("not enough arguments")
	}

	ids := make([]int64, len(args))
	for i := 0; i < len(args); i++ {
		id, err := strconv.ParseInt(args[i], 10, 64)
		if err != nil {
			return err
		}
		ids[i] = id
	}

	return task.DeleteTasks(ids)

}
