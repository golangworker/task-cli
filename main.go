package main

import (
	"fmt"
	"task-cli/list"
	"task-cli/task"
)

func main() {
	task1 := task.Init("Brush teeth")
	list, err := list.Init()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(list)
	list.Add(task1)
	list.Save()
}
