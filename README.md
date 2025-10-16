# GoDoList CLI

A simple, interactive CLI To-Do List manager written in Go. This tool allows users to efficiently manage tasks directly from the terminal.

## Features

+ **Add Tasks:** Create new tasks with a custom title.
+ **List Tasks:** Display all tasks with their ID, title, and completion status (Pending or Completed).
+ **Mark Tasks as Completed:** Update a task's status to completed by specifying its ID.
+ **Delete Tasks:** Remove tasks from the list by their ID.
+ **Exit:** Cleanly terminate the program.


## How It Works

The CLI runs in an interactive loop, prompting the user to choose from a menu of options (1–5). Tasks are stored in memory as a slice of Task structs, each containing an ID, title, and completion status. The program supports saving tasks to a JSON file for persistence across sessions.

## Usage

Run the Program: Execute go `run main.go` in the terminal.

### Interact with the Menu:

1. Add a new task by entering a title.
2. List all tasks with their details.
3. Mark a task as completed by entering its ID.
4. Delete a task by entering its ID.
5. Exit the program.

### Example
```
1. Add Task
2. List Tasks
3. Complete Task
4. Delete Task
5. Exit

Choose an option: 1
Enter task title: Buy groceries
Task added!

Choose an option: 2
ID: 1, Title: Buy groceries, Status: Pending

Choose an option: 3
Enter task ID to complete: 1
Task marked as completed!

Choose an option: 2
ID: 1, Title: Buy groceries, Status: Completed
```

## Technical Details

- **Language:** Go
- **Dependencies:** Standard library only (`bufio`, `encoding/json`, `fmt`, `os`, `strconv`, `strings`)
- **Data Structure:** Uses a `Task` struct (`ID int`, `Title string`, `Completed bool`) and a `GoDoList` struct with a slice of tasks (`Tasks []Task`).
- **Input Handling:** Utilizes `bufio.Scanner` for line-by-line input from `os.Stdin`.


## Getting Started

Clone the Repository:
```
git clone <repository-url>
cd <repository-directory>
```

Initialize the Go Module:
```
go mod init go-do-cli
```


Run the Program:
```
go run main.go
```

### Future Enhancements

- Add file persistence to automatically save/load tasks.
- Implement task filtering (e.g., show only pending tasks).
- Support task priorities or due dates.
- Add an undo feature for deletions.

[![Go Version](https://img.shields.io/badge/Go-1.22-blue)](https://golang.org/)
