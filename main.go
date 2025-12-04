package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/joho/godotenv"

	"github.com/dmandevv/task-tracker/internal/commands"
	"github.com/dmandevv/task-tracker/internal/config"
	"github.com/dmandevv/task-tracker/internal/json"
	. "github.com/dmandevv/task-tracker/internal/task"
)

func main() {
	err := godotenv.Load("task-tracker.env")
	if err != nil {
		fmt.Println("No .env file found, proceeding without it.")
	}

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
			newDesc := strings.Join(cmds[2:], " ")
			err = commands.UpdateTask(cfg, idInt, newDesc)
			if err != nil {
				fmt.Println(err)
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
	fmt.Println("\nDisplay this help message:")
	fmt.Println("help")

	fmt.Println("\nAdd a new task:")
	fmt.Println("add <task-description>")

	fmt.Println("\nUpdate new task:")
	fmt.Println("update <task-id> <new-task-description>")

	fmt.Println("\nDelete a task:")
	fmt.Println("delete <task-id>")

	fmt.Println("\nChange a task's status:")
	fmt.Println("mark-in-progress <task-id>")
	fmt.Println("mark-done <task-id>")

	fmt.Println("\nList all tasks:")
	fmt.Println("list")

	fmt.Println("\nList tasks by status:")
	fmt.Println("list todo")
	fmt.Println("list in-progress")
	fmt.Println("list done")

	fmt.Println("\nSave data and exit program:")
	fmt.Println("exit")

}
