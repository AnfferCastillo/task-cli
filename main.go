package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/AnfferCastillo/task-cli/tasks"
	"github.com/AnfferCastillo/task-cli/utils"
)

type Command interface {
	Execute(args []string) string
}

type add struct {
}

func (a add) Execute(args []string) string {
	id := tasks.AddTask(args[0])
	if id < 0 {
		return "There was a problem, could not add new task."
	}
	return fmt.Sprintf("Task added successfully (ID: %v)", id)
}

type update struct{}

func (u update) Execute(args []string) {
	id, err := strconv.Atoi(args[0])
	if err != nil {
		log.Fatal("invalid task ID")
	}

	tasks.UpdateTask(id, args[1])
}

type list struct{}

func (l list) Execute(args []string) string {
	status := "all"
	if len(args) != 0 {
		status = args[0]

	}

	tasks, err := tasks.ListTasks(status)
	if err != nil {
		log.Fatal(err)
		return "Could not load tasks"
	}

	return utils.FormatTasks(tasks)

}

func FormatTasks(tasks []tasks.TaskItem, status string) {
	panic("unimplemented")
}

type delete struct{}

func (d delete) Execute(args []string) {
	id, err := strconv.Atoi(args[0])
	if err != nil {
		return
	}
	tasks.Delete(id)
}

type mark struct{}

func (m mark) Execute(args []string) {
	id, err := strconv.Atoi(args[1])
	if err != nil {
		return
	}

	status := getStatus(args[0])
	tasks.Mark(id, status)
}

func getStatus(args string) string {
	if strings.Contains(args, "todo") {
		return "todo"
	}
	if strings.Contains(args, "done") {
		return "done"
	}
	if strings.Contains(args, "in-progress") {
		return "in-progress"
	}

	return ""
}

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Println("No options provided. Please run task-cli help to see available options.")
		return
	}

	command := args[0]
	var result string = ""

	switch command {
	case "add":
		result = add{}.Execute(args[1:])
	case "update":
		update{}.Execute(args[1:])
	case "list":
		result = list{}.Execute(args[1:])
	case "delete":
		delete{}.Execute(args[1:])

	case "mark-in-progress":
		mark{}.Execute(args)

	case "mark-done":
		mark{}.Execute(args)
	default:
		fmt.Printf("Unknown command. Run help to get available commands")
	}

	fmt.Print(result)
}
