package cli

import "flag"

var (
	Add       = flag.String("add", "", "add a description of the new task to the list")
	Delete    = flag.String("del", "", "add ID to delete a task from the list")
	Mark      = flag.String("mark", "", "mark a task as done")
	AllTasks  = flag.Bool("all", false, "print all tasks in list")
	DoneTasks = flag.Bool("done", false, "print done tasks")
	TodoTasks = flag.Bool("todo", false, "print in-progress tasks")
)
