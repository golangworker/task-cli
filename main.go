package main

import "app/list"

func main() {
	list, _ := list.Init()
	defer list.Save()

}
