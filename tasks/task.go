package tasks

import (
	"errors"
	"log"
	"time"
)

type TaskItem struct {
	Description string
	Status      string
	ID          int
}

func AddTask(taskDescription string) int {
	tasks, err := TaskStorage.LoadTasks()

	if err != nil {
		log.Fatal(err)
		return -1
	}

	task := Task{
		Description: taskDescription,
		Status:      ToDo,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	tasks.Tasks = append(tasks.Tasks, task)
	TaskStorage.Save(tasks)
	return len(tasks.Tasks)
}

func UpdateTask(id int, description string) {
	tasks, err := TaskStorage.LoadTasks()
	if err != nil {
		log.Fatal("Error while loading tasks")
	}

	if len(tasks.Tasks) < id {
		log.Fatal("Task does not exist")
	}

	task := tasks.Tasks[id-1]

	task.Description = description
	task.UpdatedAt = time.Now()
	tasks.Tasks[id-1] = task

	TaskStorage.Save(tasks)
}

func Delete(id int) {
	tasks, err := TaskStorage.LoadTasks()

	if err != nil {
		log.Fatal("Error while deleting tasks tasks")
		return
	}

	if len(tasks.Tasks) < id {
		log.Fatal("Invalid task id")
		return
	}

	index := id - 1

	newTasksSlice := tasks.Tasks[0:index]

	if id < len(tasks.Tasks) {
		newTasksSlice = append(newTasksSlice, tasks.Tasks[index+1:]...)
	}

	tasks.Tasks = newTasksSlice
	TaskStorage.Save(tasks)
}

func Mark(id int, state string) {
	tasksList, err := TaskStorage.LoadTasks()
	if err != nil {
		log.Fatal("Task does not exist")
		return
	}

	task := tasksList.Tasks[id-1]
	task.Status = StatusFromString(state)
	task.UpdatedAt = time.Now()
	tasksList.Tasks[id-1] = task
	TaskStorage.Save(tasksList)
}

func ListTasks(toStatus string) ([]TaskItem, error) {
	tasksList, err := TaskStorage.LoadTasks()
	if err != nil {
		return nil, errors.New("Task does not exist")
	}

	tasks := make([]TaskItem, 0)
	status := StatusFromString(toStatus)

	for index, task := range tasksList.Tasks {
		if status == All || task.Status == status {
			tasks = append(tasks, TaskItem{
				Description: task.Description,
				ID:          index + 1,
				Status:      task.Status.String(),
			})
		}
	}

	return tasks, nil
}
