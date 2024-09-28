package tasks

import (
	"encoding/json"
	"log"
	"os"
)

const FILE_NAME = "tasks-cli.json"

type Save func(tasks TasksList)
type Load func() (TasksList, error)

type TaskDataBase struct {
	Save      Save
	LoadTasks Load
}

var TaskStorage TaskDataBase = TaskDataBase{
	Save: writeToFile,
	LoadTasks: loadTaskList,
}

func writeToFile(taskList TasksList) {

	file, err := os.Create(FILE_NAME)

	if err != nil {
		log.Fatal("Error creating the file")
		return
	}

	defer file.Close()

	data, _ := json.Marshal(taskList)

	_, err = file.Write(data)

	if err != nil {
		log.Fatal("Error writing to file")
		return
	}
}

func loadTaskList() (TasksList, error) {
	file, err := os.ReadFile(FILE_NAME)
	if err != nil {
		return TasksList{Tasks: []Task{}}, nil
	}

	var tasks TasksList
	json.Unmarshal(file, &tasks)

	if err != nil {
		log.Fatal("Unable to read data")
	}

	return tasks, nil
}
