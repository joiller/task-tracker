package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"strconv"
	"task-tracker/internal/task"
)

func init() {
	RootCmd.AddCommand(NewUpdateCmd())
	RootCmd.AddCommand(NewMarkTodoCmd())
	RootCmd.AddCommand(NewMarkDoneCmd())
	RootCmd.AddCommand(NewMarkInProgressCmd())
}

func RunUpdateTaskCmd(args []string) error {
	id, err := strconv.ParseInt(args[0], 10, 64)
	if err != nil {
		return err
	}
	return task.UpdateTaskDescription(id, args[1])
}

func NewUpdateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update",
		Short: "Update a task description",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			return RunUpdateTaskCmd(args)
		},
	}
	return cmd

}

func RunUpdateTaskStatusCmd(args []string, status task.TaskStatus) error {
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

	return task.UpdateTaskStatus(ids, status)
}

func NewMarkTodoCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "mark-todo",
		Short: "Update a task status as todo",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return RunUpdateTaskStatusCmd(args, task.TaskStatusFromString("todo"))
		},
	}
	return cmd
}

func NewMarkDoneCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "mark-done",
		Short: "Update a task status as done",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return RunUpdateTaskStatusCmd(args, task.TaskStatusFromString("done"))
		},
	}
	return cmd
}

func NewMarkInProgressCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "mark-in-progress",
		Short: "Update a task status as in progress",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return RunUpdateTaskStatusCmd(args, task.TaskStatusFromString("in_progress"))
		},
	}
	return cmd
}
