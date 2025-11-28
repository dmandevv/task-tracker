package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/dmandevv/task-tracker/internal/commands"
	"github.com/dmandevv/task-tracker/internal/config"
	"github.com/dmandevv/task-tracker/internal/json"
	. "github.com/dmandevv/task-tracker/internal/task"
)

func main() {
	cfg, err := json.LoadTasksFromFile()
	if err != nil {
		fmt.Println("No existing task file found, starting fresh.")
		cfg = &config.Config{
			Tasks:  make([]Task, 0),
			NextID: 1,
		}
	}

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)

	fmt.Println("Task Tracker initialized")

	for _, cmd := range os.Args {
		if cmd == "-v" {
			displayHelp()
		}
	}

CLI:
	for {
		fmt.Print("> ")
		scanner.Scan()
		cmds := strings.Fields(scanner.Text())
		if len(cmds) == 0 {
			continue CLI
		}
		switch cmds[0] {
		case "add":
			if len(cmds) < 2 {
				fmt.Println("Please provide a task description")
				continue CLI
			}
			desc := strings.Join(cmds[1:], " ")
			commands.AddTask(cfg, desc)
			fmt.Println("Task added")
		case "update":
			if len(cmds) < 3 {
				fmt.Println("Please provide a task ID and new description")
				continue CLI
			}
			id := cmds[1]
			idInt, err := strconv.Atoi(id)
			if err != nil {
				fmt.Println("Invalid task ID")
				continue CLI
			}
			newDesc := cmds[2]
			err = commands.UpdateTask(cfg, idInt, newDesc)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Task updated")
			}
		case "delete":
			if len(cmds) < 2 {
				fmt.Println("Please provide a task ID")
				continue CLI
			}
			id := cmds[1]
			idInt, err := strconv.Atoi(id)
			if err != nil {
				fmt.Println("Invalid task ID")
				continue CLI
			}
			err = commands.DeleteTask(cfg, idInt)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Task deleted")
			}
		case "mark-in-progress":
			if len(cmds) < 2 {
				fmt.Println("Please provide a task ID")
				continue CLI
			}
			id := cmds[1]
			idInt, err := strconv.Atoi(id)
			if err != nil {
				fmt.Println("Invalid task ID")
				continue CLI
			}
			err = commands.MarkTask(cfg, idInt, IN_PROGRESS)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Task marked as in-progress")
			}
		case "mark-done":
			if len(cmds) < 2 {
				fmt.Println("Please provide a task ID")
				continue CLI
			}
			id := cmds[1]
			idInt, err := strconv.Atoi(id)
			if err != nil {
				fmt.Println("Invalid task ID")
				continue CLI
			}
			err = commands.MarkTask(cfg, idInt, DONE)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Task marked as done")
			}
		case "list":
			var filteredTasks []Task

			if len(cmds) < 2 {
				filteredTasks = cfg.Tasks
			} else {
				status := cmds[1]
				switch status {
				case "todo":
					filteredTasks = commands.GetTasksByFilter(cfg, TODO)
				case "in-progress":
					filteredTasks = commands.GetTasksByFilter(cfg, IN_PROGRESS)
				case "done":
					filteredTasks = commands.GetTasksByFilter(cfg, DONE)
				default:
					fmt.Println("Unknown status filter")
					continue CLI
				}
			}
			if len(filteredTasks) == 0 {
				fmt.Println("No tasks found")
				continue CLI
			}
			for index, t := range filteredTasks {
				fmt.Printf("%d. (ID: %d) [%s] %s (Created: %s)\n",
					index+1,
					t.ID,
					t.Status.String(),
					t.Description,
					t.CreatedAt.Format(time.RFC822))
			}

		case "help":
			displayHelp()
		case "exit":
			err := json.SaveTasksToFile(cfg)
			if err != nil {
				fmt.Printf("Error saving tasks to file: %v\n", err)
				continue CLI
			}
			fmt.Println("Tasks have been saved. Exiting Task Tracker. Goodbye!")
			break CLI
		default:
			fmt.Println("Unknown command. Type <help> for a list of commands.")
		}
	}

}

func displayHelp() {
	fmt.Println("\nCOMMANDS")

	fmt.Println("Display this help message")
	fmt.Println("help")

	fmt.Println("Add a new task")
	fmt.Println("add <task-description>")

	fmt.Println("Update new task")
	fmt.Println("update <task-id> <new-task-description>")

	fmt.Println("Delete a task")
	fmt.Println("delete <task-id>")

	fmt.Println("Change a task's status")
	fmt.Println("mark-in-progress <task-id>")
	fmt.Println("mark-done <task-id>")

	fmt.Println("List all tasks")
	fmt.Println("list")

	fmt.Println("List tasks by status")
	fmt.Println("list todo")
	fmt.Println("list in-progress")
	fmt.Println("list done")

	fmt.Println("Save data and exit program")
	fmt.Println("exit")

}
