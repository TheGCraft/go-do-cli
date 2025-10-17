package main

import (
	"reflect"
	"testing"
)

// Tests the CompleteTask method
func TestCompleteTask(t *testing.T) {
	tables := []struct {
		name      string //Name for the test case
		tasks     []Task //Initial tasks in GoDoList
		id        int    //Input ID for CompleteTask
		want      bool   //Expected return value
		wantTasks []Task //Expected tasks after calling CompletedTask
	}{
		{
			name:      "Valid ID - single task",
			tasks:     []Task{{ID: 1, Title: "Buy food", Completed: false}},
			id:        1,
			want:      true,
			wantTasks: []Task{{ID: 1, Title: "Buy food", Completed: true}},
		},
		{
			name:      "Invalid ID - single task",
			tasks:     []Task{{ID: 1, Title: "Buy food", Completed: false}},
			id:        2,
			want:      false,
			wantTasks: []Task{{ID: 1, Title: "Buy food", Completed: false}},
		},
		{
			name:      "Empty task list",
			tasks:     []Task{},
			id:        1,
			want:      false,
			wantTasks: []Task{},
		},
		{
			name: "Multiple tasks - mark one",
			tasks: []Task{
				{ID: 1, Title: "Buy food", Completed: false},
				{ID: 2, Title: "Write email", Completed: false},
			},
			id:   1,
			want: true,
			wantTasks: []Task{
				{ID: 1, Title: "Buy food", Completed: true},
				{ID: 2, Title: "Write email", Completed: false},
			},
		},
	}

	for _, tt := range tables {
		t.Run(tt.name, func(t *testing.T) {
			//Set up the GoDoList with the initial tasks
			g := &GoDoList{Tasks: tt.tasks}

			//Call CompleteTask
			got := g.CompleteTask(tt.id)

			//Check return value
			if got != tt.want {
				t.Errorf("CompleteTask(%d) = %v, want %v", tt.id, got, tt.want)
			}

			//Check the state of the Tasks slice
			if !reflect.DeepEqual(g.Tasks, tt.wantTasks) {
				t.Errorf("Tasks after CompleteTask(%d) = %v; want %v", tt.id, g.Tasks, tt.wantTasks)
			}
		})
	}
}

// Tests the DeleteTask method
func TestDeleteTask(t *testing.T) {
	tables := []struct {
		name      string //To identify the test scenario
		tasks     []Task //Initial Tasks in GoDoList
		id        int    //Input ID for DeleteTask
		want      bool   //Expected return value
		wantTasks []Task //Expected tasks after calling DeleteTask
	}{
		{
			name:      "Valid ID - single task",
			tasks:     []Task{{ID: 1, Title: "Cook", Completed: false}},
			id:        1,
			want:      true,
			wantTasks: []Task{},
		}, {
			name:      "Invalid ID - single task",
			tasks:     []Task{{ID: 1, Title: "Cook", Completed: false}},
			id:        2,
			want:      false,
			wantTasks: []Task{{ID: 1, Title: "Cook", Completed: false}},
		}, {
			name:      "Empty list - no tasks",
			tasks:     []Task{},
			id:        1,
			want:      false,
			wantTasks: []Task{},
		}, {
			name: "Multiple tasks, deleting first element and reassigning IDs",
			tasks: []Task{
				{ID: 1, Title: "Cook", Completed: false},
				{ID: 2, Title: "Clean", Completed: true},
				{ID: 3, Title: "Do the laundry", Completed: false},
			},
			id:   1,
			want: true,
			wantTasks: []Task{
				{ID: 1, Title: "Clean", Completed: true},
				{ID: 2, Title: "Do the laundry", Completed: false},
			},
		}, {
			name: "Multiple tasks, deleting middle element and reassigning IDs",
			tasks: []Task{
				{ID: 1, Title: "Cook", Completed: false},
				{ID: 2, Title: "Clean", Completed: true},
				{ID: 3, Title: "Do the laundry", Completed: false},
			},
			id:   2,
			want: true,
			wantTasks: []Task{
				{ID: 1, Title: "Cook", Completed: false},
				{ID: 2, Title: "Do the laundry", Completed: false},
			},
		}, {
			name: "Multiple tasks, deleting last element and reassigning IDs",
			tasks: []Task{
				{ID: 1, Title: "Cook", Completed: false},
				{ID: 2, Title: "Clean", Completed: true},
				{ID: 3, Title: "Do the laundry", Completed: false},
			},
			id:   3,
			want: true,
			wantTasks: []Task{
				{ID: 1, Title: "Cook", Completed: false},
				{ID: 2, Title: "Clean", Completed: true},
			},
		}, {
			name:      "Edge Case - Last Task from the list deleted",
			tasks:     []Task{{ID: 1, Title: "Cook", Completed: false}},
			id:        1,
			want:      true,
			wantTasks: []Task{},
		},
	}

	for _, dt := range tables {
		t.Run(dt.name, func(t *testing.T) {
			//Set up the GoDoList with the initial tasks
			g := &GoDoList{Tasks: dt.tasks}

			//Call DeleteTask
			dot := g.DeleteTask(dt.id)
			//Check return value
			if dot != dt.want {
				t.Errorf("DeleteTask(%d) = %v, want %v", dt.id, dot, dt.want)

			}
			//Check the state of the Tasks slice
			if !reflect.DeepEqual(g.Tasks, dt.wantTasks) {
				t.Errorf("DeleteTask(%d) = %v. want %v", dt.id, g.Tasks, dt.want)

			}

		})

	}

}

// Tests the AddTask method
func TestAddTask(t *testing.T) {
	tables := []struct {
		name      string //To identify test scenario
		initial   []Task //Initial tasks in GoDoList
		title     string //Input Name for AddTask
		wantTasks []Task //Expected tasks after calling AddTask

	}{
		{
			name:      "Valid Name - Empty list",
			initial:   []Task{},
			title:     "Buy food",
			wantTasks: []Task{{ID: 1, Title: "Buy food", Completed: false}},
		}, {
			name:      "Invalid Title - no task added",
			initial:   []Task{},
			title:     "",
			wantTasks: []Task{{ID: 1, Title: "", Completed: false}},
		}, {
			name: "Multiple tasks - valid ID",
			initial: []Task{
				{ID: 1, Title: "Buy food", Completed: false},
				{ID: 2, Title: "Cook", Completed: true},
			},
			title: "Clean",
			wantTasks: []Task{
				{ID: 1, Title: "Buy food", Completed: false},
				{ID: 2, Title: "Cook", Completed: true},
				{ID: 3, Title: "Clean", Completed: false},
			},
		}, {
			name: "Long title",
			initial: []Task{
				{ID: 1, Title: "Write code", Completed: true},
			},
			title: "Complete a very long task description for testing purposes",
			wantTasks: []Task{
				{ID: 1, Title: "Write code", Completed: true},
				{ID: 2, Title: "Complete a very long task description for testing purposes", Completed: false},
			},
		},
	}

	for _, at := range tables {
		t.Run(at.name, func(t *testing.T) {
			//Set up the GoDoList with the initial tasks
			g := &GoDoList{Tasks: at.initial}
			//Call AddTask
			g.AddTask(at.title)
			//Check the state of the Tasks slice
			if !reflect.DeepEqual(g.Tasks, at.wantTasks) {
				t.Errorf("AddTask(%q) = %v, want %v", at.title, g.Tasks, at.wantTasks)

			}

		})

	}

}
