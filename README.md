# Task Tracker 📝

A simple command-line task management application written in Go.

## Installation 📦

To download and install Task Tracker, run:

```shell
go install github.com/dmandevv/task-tracker@latest
```

Alternatively, clone the repository and build a local binary:

```shell
git clone https://github.com/dmandevv/task-tracker.git
cd task-tracker
go build -o task-tracker
```

## Setup ⚙️

Task Tracker persists tasks to a JSON file. By default the app uses `mytasks` in the current working directory. You can override the path using the `TASK_TRACKER_SAVE_FILE_PATH` environment variable.

Create a `.env` file in the project root (or set the variable in your shell) with the following content:

```env
TASK_TRACKER_SAVE_FILE_PATH="ENTER_FILEPATH_HERE"
```

- If `TASK_TRACKER_SAVE_FILE_PATH` is not set, the default file is `mytasks`.
- The program will create the file automatically when you save tasks.

## Running the Program 🚀

Start the CLI from the repository root (or run the installed binary):

```sh
./task-tracker
```

You can also run with the `-v` flag to show help text on startup:

```sh
./task-tracker -v
```

Once running you'll see a prompt (`>`). Type commands described below and press Enter.

## Commands 🧰

All commands are interactive at the prompt. Replace values in angle brackets.

- **Add a task**: `add <task-description>`

	Example:

	```
	> add Buy groceries and milk
	Buy groceries and milk added
	```

- **List tasks**: `list` (all) or `list <status>` (filter by status)

	Status values: `todo`, `in-progress`, `done`.

	Examples:

	```
	> list
	1. (ID: 1) [todo] Buy groceries and milk (Created: 04 Dec 25 12:34 UTC)

	> list done
	No tasks found
	```

- **Update a task description**: `update <task-id> <new-description>`

	Example:

	```
	> update 1 Buy groceries and cook dinner
	Buy groceries and milk updated to Buy groceries and cook dinner
	```

- **Delete a task**: `delete <task-id>`

	Example:

	```
	> delete 1
	Buy groceries and cook dinner deleted
	```

- **Mark task in progress**: `mark-in-progress <task-id>`

	Example:

	```
	> mark-in-progress 2
	Write blog post marked as IN_PROGRESS
	```

- **Mark task done**: `mark-done <task-id>`

	Example:

	```
	> mark-done 2
	Write blog post marked as DONE
	```

- **Help**: `help` — display a summary of available commands.

- **Exit**: `exit` — save tasks to the JSON file and exit the program.

	Example:

	```
	> exit
	Tasks have been saved. Exiting Task Tracker. Goodbye!
	```

## Implementation notes 🧭

- Tasks are stored in a `Config` object and saved/loaded as JSON. The code reads `SAVE_FILE_PATH` (environment) and falls back to `task_tracker.json`.
- The main logic and command handlers live in `main.go` and `internal/commands`:
	- `internal/commands/add.go` — `AddTask(cfg, description)`
	- `internal/commands/list.go` — `GetTasksByFilter(cfg, status)`
	- `internal/commands/update.go` — `UpdateTask(cfg, id, description)` and `MarkTask(cfg, id, status)`
	- `internal/commands/delete.go` — `DeleteTask(cfg, id)`