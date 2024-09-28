package tasks

import (
	"testing"
)

func TestAddTask(t *testing.T) {
	//given
	mockTaskDB()

	taskDescription := "just a simple task to do"
	//when
	var id int = AddTask(taskDescription)
	//then
	if id <= 0 {
		t.Fatalf("Adding task %q returned %v", taskDescription, id)
	}
}

func TestUpdateTask(t *testing.T) {
	//given
	mockTaskDB()
	AddTask("old Task")
	//when
	UpdateTask(1, "new task description")
	tasks, _ := ListTasks("ToDo")
	if tasks[0].Description != "new task description" {
		t.Fatalf("Error updating task, expected description: 'new task description' \n but got: %v", tasks[0].Description)
	}

}

func TestListAllTask(t *testing.T) {
	//given
	mockTaskDB()
	AddTask("Task 1")
	AddTask("Task 2")

	//then
	tasks, err := ListTasks("ToDo")

	if err != nil {
		t.Fatal(err)
	}
	if len(tasks) != 2 {
		t.Fatalf("Expected 2 tasks but found %v", len(tasks))
	}
}

func TestDeleteTask(t *testing.T) {
	//given
	mockTaskDB()
	AddTask("First task")
	tasks, _ := ListTasks("ToDo")

	if len(tasks) != 1 {
		t.Fatalf("Error adding a task")
	}

	Delete(1)
	tasks, _ = ListTasks("ToDo")
	if len(tasks) != 0 {
		t.Fatalf("Expected 0 tasks \n Found %v tasks.", len(tasks))
	}

}

func TestMarkTask(t *testing.T) {
	mockTaskDB()
	AddTask("A TODO task")

	Mark(1, "done")
	tasks, _ := ListTasks("done")
	if len(tasks) != 1 {
		t.Fatalf("Expected 0 tasks \n Found %v tasks.", len(tasks))
	}
}

func mockTaskDB() {
	tasksList := TasksList{
		Tasks: make([]Task, 0),
	}

	TaskStorage = TaskDataBase{
		Save: func(tasks TasksList) {
			tasksList = tasks
		},
		LoadTasks: func() (TasksList, error) {
			return tasksList, nil
		},
	}
}
