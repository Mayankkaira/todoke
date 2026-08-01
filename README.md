# todoke

A command-line task manager written in Go. Add, list, delete, and toggle tasks — saves to a local file so it persists between runs.

Running it
go run .

Or build a binary:
go build -o todoke .
./todoke

You'll get a numbered menu: add a task, list tasks, delete a task, toggle a task's complete/incomplete status, or exit.

Structure
.
├── main.go       
├── display.go   
├── task/
│   └── task.go   
└── storage/
    └── storage.go 