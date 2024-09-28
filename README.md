# Task CLI Tool

This is a simple CLI tool to manage tasks. You can list, delete, and mark tasks with different statuses.

It will create a task-cli.json file in your current directory to keep track of your tasks.

## Prerequisites

- Go (version 1.16 or higher)

## Installation

1. Clone the repository:
    ```sh
    git clone https://github.com/yourusername/task-cli-tool.git
    cd task-cli-tool
    ```

2. Build the CLI tool:
    ```sh
    go build -o task-cli
    ```
3. Install the CLI tool:
    ```sh
    go install 
    ```

## Usage

Run the CLI tool with the following command:
```sh
    task-cli <command> [arguments]
```

## Commands

#### Add a Tasks
Add a new task to your list.

```sh
 task-cli add <task_description>
```
- `task_description`: The task description to add.

#### List Tasks
List all tasks or tasks with a specific status.

```sh
 task-cli list <status>
```
- `status` (optional):  The status of the tasks to list (todo, done, in-progress). If not provided, all task will be listed.

#### Upddate a Tasks
Update a task by ID.

```sh
 task-cli udpate <ID> <task_description>
```
- `ID`: ID of the task to update.
- `task_description`: The new task description.


#### Delete Task
Delete a task by ID.

```sh
 task-cli delete <task_id>
```
- `task_id`: The ID of the task to delete.

#### Mark Task
Mark a task with a specific status.

```sh
task-cli mark <status> <task_id>
```

- `status`: The status to mark the task with (e.g., todo, done).
- `task_id`: The ID of the task to mark.

