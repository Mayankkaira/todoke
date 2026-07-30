package storage

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/mayankkaira/taskmanager/task"
)

func SaveTasks(tasks []task.Task) error {
	var content strings.Builder
	file, err := os.Create("tasks.txt")
	if err != nil {
		return err
	}
	defer file.Close()
	for _, value := range tasks {
		content.WriteString(value.Name)
		content.WriteString("|")
		content.WriteString(strconv.FormatBool(value.Completed))
		content.WriteString("\n")
	}
	_, err = file.WriteString(content.String())
	if err != nil {
		return err
	}
	return nil
}

func LoadTasks() ([]task.Task, error) {
	file, err := os.Open("tasks.txt")
	if err != nil {
		if os.IsNotExist(err) {
			return []task.Task{}, nil
		}
		return nil, err
	}
	defer file.Close()
	tasks := []task.Task{}
	reader := bufio.NewReader(file)
	for {
		data, err := reader.ReadString('\n')
		if data != "" {
			data = strings.TrimSpace(data)
			parts := strings.Split(data, "|")
			if len(parts) != 2 {
				return nil, fmt.Errorf("invalid task format")
			}
			part, err := strconv.ParseBool(parts[1])
			if err != nil {
				return nil, err
			}
			task := task.Task{Name: parts[0], Completed: part}
			tasks = append(tasks, task)
		}

		if err == io.EOF {
			return tasks, nil
		}
		if err != nil {
			return nil, err
		}

	}
}
