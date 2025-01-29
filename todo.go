package main

import (
	"errors"
	"fmt"
	"time"
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
func (todos *Todos) delete(index int) error{
	t := *todos

	if err := t.validate(index); err!=nil {
        return err
	}

	*todos = append(t[:index],t[index+1:]...)

	return  nil
}
