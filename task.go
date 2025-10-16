package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// ---Define Data Structure
type Task struct {
	ID        int
	Title     string
	Completed bool
}

//Holds a slice of Task structs to store tasks

type GoDoList struct {
	Tasks []Task
}

//---Methods

// Add a Task
// Creates a new task with an ID based on the current length of the slice plus 1
// Set the task's Title to the provided string and Completed to false by default
// Appends the task to the Tasks slice using append
func (g *GoDoList) AddTask(title string) {
	task := Task{
		ID:        len(g.Tasks) + 1,
		Title:     title,
		Completed: false,
	}
	g.Tasks = append(g.Tasks, task)
}

// Complete a Task
// Iterates through the Tasks slice
// If the task ID matches the input id, updates the Completed field and returns true
// Returns fals if no task with the given ID is found
func (g *GoDoList) CompleteTask(id int) bool {
	for i, task := range g.Tasks {
		if task.ID == id {
			g.Tasks[i].Completed = true

			return true
		}
	}

	return false
}

// Delete a Task
// Iterates through the Tasks slice to find the matching id
// Removes the task using slice manipulation
// Reassigns IDs to all remaining tasks
func (g *GoDoList) DeleteTask(id int) bool {
	for i, task := range g.Tasks {
		if task.ID == id {
			//Remove the task by slicing
			g.Tasks = append(g.Tasks[:i], g.Tasks[i+1:]...)
			//Reassign IDs to mantain consecutive numbering
			for j := range g.Tasks {
				g.Tasks[j].ID = j + 1
			}

			return true
		}
	}

	return false
}

// List the Tasks
// Iterates through the Tasks slice
// For each task, sets status based on the Completed field
// Prints each task's details
func (g *GoDoList) ListTasks() {
	for _, task := range g.Tasks {
		status := "Pending"
		if task.Completed {
			status = "Completed"
		}
		fmt.Printf("ID: %d, Title: %s, Status: %s\n", task.ID, task.Title, status)

	}
}

// Save to file
func (g *GoDoList) SaveToFile(filename string) error {
	data, err := json.Marshal(g.Tasks)
	if err != nil {
		return err
	}
	return os.WriteFile(filename, data, 0644)
}
func (g *GoDoList) LoadFromFile(filename string) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, &g.Tasks)
}

// ---Handling User Input
func main() {
	//Instance to store tasks
	list := &GoDoList{}
	//Initialize bufio.Scanner to read input from os.Stdin (terminal)
	scanner := bufio.NewScanner(os.Stdin)

	//Main loop
	for {
		fmt.Println("\n1. Add Task\n2. List Tasks\n3. Complete Task\n4 Delete Task \n5. Exit")
		fmt.Print("Choose an option: ")
		//Read user input
		scanner.Scan()
		//get the user input as string
		choice := scanner.Text()

		//Choice's cases
		switch choice {
		//Calls AddTask
		case "1":
			fmt.Print("Enter task title: ")
			scanner.Scan()
			title := scanner.Text()
			list.AddTask(title)
			fmt.Println("Task added!")
		//Calls ListTasks
		case "2":
			list.ListTasks()
		//Prompts for a task ID, reads it and converts it to an integer
		case "3":
			fmt.Print("Enter task ID to complete: ")
			scanner.Scan()
			idStr := scanner.Text()
			//Checking for conversion errors (non-numeric input)
			id, err := strconv.Atoi(idStr)
			if err != nil {
				fmt.Println("Invalid ID. Please enter a number")
				continue
			}
			//Calls CompleteTask or prints failure message
			if list.CompleteTask(id) {
				fmt.Println("Task marked as completed!")
			} else {
				fmt.Println("Task ID not found.")
			}
		//Prompts for a Task ID, reads it and trims it
		case "4":
			fmt.Print("Enter task ID to delete: ")
			scanner.Scan()
			idStr := strings.TrimSpace(scanner.Text())
			//Converting ID to an integer and checking for errors
			id, err := strconv.Atoi(idStr)
			if err != nil {
				fmt.Println("Invalid ID. Please enter a number")
				continue
			}
			//Calls DeleteTask and prints success or failure message
			if list.DeleteTask(id) {
				fmt.Println("Task deleted successfully!")
			} else {
				fmt.Println("Task ID not found.")
			}
		//Exiting the program (0 = successful termination)
		case "5":
			os.Exit(0)
		default:
			fmt.Println("Invalid option")
		}
	}

}
