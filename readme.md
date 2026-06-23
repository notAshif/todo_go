# Go Todo CLI

A simple command-line Todo application built in Go for learning and practicing:

* Structs and Methods
* Slices
* JSON Storage
* File Handling
* Command Line Flags
* Generics
* Error Handling
* Time Management
* CRUD Operations

## Features

* Add new todos
* List all todos
* Edit existing todos
* Delete todos
* Mark todos as completed
* Store data persistently in a JSON file
* Display creation and completion timestamps

## Project Structure

```text
.
├── main.go
├── command.go
├── todo.go
├── storage.go
├── todos.json
└── README.md
```

## Installation

Clone the repository:

```bash
git clone <repository-url>
cd todo-cli
```

Install dependencies:

```bash
go mod tidy
```

Run the application:

```bash
go run .
```

## Usage

### Add a Todo

```bash
go run . -Add "Build a Go Todo CLI"
```

### List Todos

```bash
go run . -ls
```

### Edit a Todo

```bash
go run . -ed "1:Learn Advanced Go"
```

Format:

```text
<index>:<new_title>
```

### Toggle Completion Status

Mark a task as completed or incomplete:

```bash
go run . -tg 1
```

### Delete a Todo

```bash
go run . -rm 1
```

## Example Workflow

```bash
go run . -Add "Learn Go"
go run . -Add "Build Todo App"

go run . -ls

go run . -tg 0

go run . -ed "1:Build Advanced Todo App"

go run . -rm 0
```

## Data Storage

Todos are stored locally in:

```text
todos.json
```

The application automatically:

1. Loads todos from storage on startup.
2. Executes the requested command.
3. Saves updated data back to storage.

## Learning Objectives

This project was built as a practice exercise to understand:

* Go structs
* Receiver methods
* Pointer receivers
* Slice manipulation
* JSON encoding and decoding
* Command-line applications
* Generic types
* Persistent storage
* Error handling patterns

## Future Improvements

* Search todos
* Due dates
* Priority levels
* Categories and tags
* Colored terminal output
* Interactive TUI interface
* Unit tests
* Sorting and filtering
* Export and import functionality

## Tech Stack

* Go
* Standard Library
* aquasecurity/table

## License

This project is intended for educational and learning purposes.
