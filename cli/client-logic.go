package cli

import (
	"app/list"
	"app/task"
	"flag"
	"fmt"
)

// runLoop is the main loop of the task-cli application.
func Runloop() {

	flag.Parse()

	switch {
	case *Add != "":
		add(*Add)
	case *Delete != "":
		delete(*Delete)
	case *Done:
		mark()
	case *AllTasks:
		allTasks()
	case *DoneTasks:
		doneTasks()
	case *TodoTasks:
		todoTasks()
	default:
		welcome()
	}
}

func welcome() {
	fmt.Println("Welcome to task-cli!")
	fmt.Println("Type '--help' for a list of commands.")
}

func add(v string) {
	task := task.Init(v)
	list, err := list.Init()
	if err != nil {
		panic(err)
	}
	defer list.Save()
	list.Add(task)
}

func delete(v string) {
	list, err := list.Init()
	if err != nil {
		panic(err)
	}
	defer list.Save()
	list.Delete(v)
}

func mark() {

}

func allTasks() {
	list, err := list.Init()
	if err != nil {
		panic(err)
	}
	list.AllTasks()
}

func doneTasks() {
	list, err := list.Init()
	if err != nil {
		panic(err)
	}
	list.DoneTasks()
}

func todoTasks() {
	list, err := list.Init()
	if err != nil {
		panic(err)
	}
	list.TodoTasks()
}
