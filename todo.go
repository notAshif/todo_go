package main

import (
	"errors"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/aquasecurity/table"
)

type Todo struct {
	Title       string
	Complete    bool
	CreateAt    time.Time
	CompletedAt *time.Time
}

type Todos []Todo

func (t *Todos) validateIndex(index int) error {
	if index < 0 || index >= len(*t) {
		err := errors.New("INVALID INDEX ❌")
		log.Printf("INVALID INDEX ❌ %s", err)
		return err
	}
	return nil
}

// Add Todo
func (t *Todos) add(title string) {
	todo := Todo{
		Title:       title,
		Complete:    false,
		CreateAt:    time.Now(),
		CompletedAt: nil,
	}

	*t = append(*t, todo)
}

// update todo by title
func (t *Todos) update(index int, title string) {
	if err := t.validateIndex(index); err != nil {
		log.Printf("FAILED TO UPDATE/EDIT TODoooo SORRY!! 😔 %s", err)
	}
	todo:= *t
	todo[index].Title = title
}

// delete by index
func (t *Todos) delete(index int) {
	todos := *t
	if err := todos.validateIndex(index); err != nil {
		log.Printf("FAILED TO REMOVE TODooo SORRY!!!! 😔 %s", err)
	}
	log.Println("TODO IS DELETED SUCCESSFULLY!!! 😊")
	*t = append(todos[:index], todos[:index+1]...)
}

// toggle
func (t *Todos) toggle(index int) {
	if err := t.validateIndex(index); err != nil {
		log.Printf("TODO IS NOT FOUND!!!🤷‍♂️ %s", err)
	}

	todos := *t
	todo := &todos[index]

	if !todo.Complete {
		completedTime := time.Now()
		todo.CompletedAt = &completedTime
	} else {
		todo.CompletedAt = nil
	}
	todo.Complete = !todo.Complete
}

// DISPLAY THE TODO
func (t *Todos) print() {
	table := table.New(os.Stdout)
	table.SetAlignment()
	table.SetRowLines(false)
	table.SetHeaders("#", "Title", "Completed", "Created At", "Completed At")

	for i, todo := range *t {
		completed := "❌"
		completedAt := ""

		if todo.Complete  {
			completed = "✅"
			if todo.CompletedAt != nil {
				completedAt = todo.CompletedAt.Format(time.RFC3339)
			}
		}
		table.AddRow(
			strconv.Itoa(i),
			todo.Title,
			completed,
			todo.CreateAt.Format(time.RFC3339),
			completedAt,
		)
	}
	table.Render()
}
