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
	complete    bool
	createAt    time.Time
	completedAt *time.Time
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
		complete:    false,
		createAt:    time.Now(),
		completedAt: nil,
	}

	if t != nil {
		log.Printf("I'M SAY TO YOU SORRY, FAILED TO ADD TODO😭")
	}
	*t = append(*t, todo)
}

// update todo by title
func (t *Todos) update(index int, title string) {
	if err := t.validateIndex(index); err != nil {
		log.Printf("FAILED TO UPDATE/EDIT TODoooo SORRY!! 😔 %s", err)
	}
	res := (*t)[index].Title == title
	log.Println("UPDATED TODO -> ", res)
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

	if !todo.complete {
		completedTime := time.Now()
		todo.completedAt = &completedTime
	} else {
		todo.completedAt = nil
	}
	todo.complete = !todo.complete
}

// DISPLAY THE TODO
func (t *Todos) print() {
	table := table.New(os.Stdout)
	table.SetAlignment()
	table.SetRowLines(false)
	table.SetHeaders("#", "Title", "Completed", "Create At", "Completed At")

	for i, todo := range *t {
		completed := "❌"
		completedAt := ""

		if todo.complete {
			completed = "✅"

			if todo.completedAt != nil {
				completedAt = time.Now().Format(time.RFC1123)
			}
			table.AddRow(
				strconv.Itoa(i),
				todo.Title,
				completed,
				todo.createAt.Format(time.RFC1123),
				completedAt,
			)
		}
	}
	table.Render()
}
