# Todo CLI App in Go

This is a simple command-line todo application written in Go. It allows you to manage a todo list by adding, updating, deleting, and listing tasks, all stored in a JSON file.

## Features

- Add new tasks
- Mark tasks as in-progress or done
- Update task descriptions
- Delete tasks
- List tasks by status (todo, in-progress, done)

## Usage

### Build and Run

1. Install Go: [Go installation](https://golang.org/dl/)
2. Clone the repository:
   ```bash
   git clone <repository_url>
   cd <repository_folder>
   ```
3. Build the app:
   ```bash
    go build -o todolist
   ```
4. Run the app:

   ```bash
   ./todolist <command> [options]

   ```

### Available Commands

1. add: Add a new todo

   ```bash
   ./todolist add --task "Task description"
   ```

2. list: List all todos or filter by status

   ```bash
   ./todolist list [status]
   ```

3. update: Update a task by ID

   ```bash
   ./todolist update <id> "Updated task description"
   ```

4. mark-done: Mark a task as done by ID

   ```bash
   ./todolist mark-done <id>
   ```

5. mark-in-progress: Mark a task as in-progress by ID

   ```bash
   ./todolist mark-in-progress <id>
   ```

6. delete: Delete a task by ID

   ```bash
   ./todolist delete <id>
   ```

### Example

1. Add a new todo:

   ```bash
   ./todolist add --task "Write a blog post"
   ```

2. List all todos:

   ```bash
   ./todolist list
   ```

3. Mark a todo as done:

   ```bash
   ./todolist mark-done 123456789
   ```

4. Delete a todo:

   ```bash
   ./todolist delete 123456789
   ```

### JSON Storage

All todos are stored in a todos.json file in the same directory as the app. The file is automatically created if it doesn't exist.
