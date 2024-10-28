package main

import (
	"fmt"
	"task-tracker/cmd"
)

func main() {
	err := cmd.RootCmd.Execute()
	if err != nil {
		fmt.Println(err)
	}
}
