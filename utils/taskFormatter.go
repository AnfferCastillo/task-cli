package utils

import (
	"fmt"
	"log"

	"github.com/AnfferCastillo/task-cli/tasks"
)

func FormatTasks(tasks []tasks.TaskItem) string {
	var result string = ""
	if len(tasks) == 0 {
		return "No tasks found"
	}

	tasksByStatus, err := splitTasksByStatus(tasks)

	if err != nil {
		log.Fatal("Error while splitting tasks by status", err)
		return "Unable to return tasks"
	}

	for _, status := range []string{"To Do", "In Progress", "Done"} {
		tasks := tasksByStatus[status]
		if len(tasks) == 0 {
			continue
		}
		result += fmt.Sprintf("Tasks in (%s):\n", status)
		for _, task := range tasks {
			result += fmt.Sprintf("\t%d. %s\n", task.ID, task.Description)
		}
	}

	return result
}

func splitTasksByStatus(theTasks []tasks.TaskItem) (map[string][]tasks.TaskItem, error) {
	if theTasks == nil {
		return nil, fmt.Errorf("tasks array is nil")
	}

	taskMap := make(map[string][]tasks.TaskItem)
	for _, task := range theTasks {
		taskMap[task.Status] = append(taskMap[task.Status], task)
	}

	return taskMap, nil
}
