package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"
    "github.com/aquasecurity/table"
)

type Todo struct {
	Title       string
	Completed   bool
	CreatedAt   time.Time
	CompletedAt *time.Time
}

type Todos []Todo

// adding a task

func (todos *Todos) add(title string) {
	todo := Todo{
		Title:       title,
		Completed:   false,
		CreatedAt:   time.Now(),
		CompletedAt: nil,
	}

	*todos = append(*todos, todo)

}

// validating an index
func (todos *Todos) validate(index int) error {
	if index < 0 || index >= len(*todos) {
		err := errors.New("Invalid index")
		fmt.Println(err)
		return err
	}

	return nil
}

// deleting a task by index
func (todos *Todos) delete(index int) error {
	t := *todos

	if err := t.validate(index); err != nil {
		return err
	}

	*todos = append(t[:index], t[index+1:]...)

	return nil
}

func (todos *Todos) toggle(index int) error {
	t := *todos

	if err := t.validate(index); err != nil {
		return err
	}
	iscompleted := t[index].Completed

	if !iscompleted {
		completedtime := time.Now()
		t[index].CompletedAt = &completedtime

	}

	t[index].Completed = !iscompleted

	return nil
}

func (todos *Todos) edit(index int, title string) error {
	t := *todos

	if err := t.validate(index); err != nil {
		return err
	}

	t[index].Title = title

	return nil
}

func (todos *Todos) print() {
	table := table.New(os.Stdout)
	table.SetRowLines(false)
	table.SetHeaders("#", "Title", "Completed", "Created At", "Completed At")

	for index, t := range *todos {
		completed := "❌"
		completedAt := ""
		if t.Completed {
			completed = "👍"
			if t.CompletedAt != nil {
				completedAt = t.CompletedAt.Format(time.RFC1123)
			}
		}

		table.AddRow(strconv.Itoa(index), t.Title, completed, t.CreatedAt.Format(time.RFC1123), completedAt)
	}

	table.Render()

}
