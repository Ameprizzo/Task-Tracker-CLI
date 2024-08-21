# Task Tracker CLI

A simple and efficient command-line task tracker built with Go. This tool allows you to manage your tasks seamlessly from the terminal, with all tasks persisted in a JSON file for easy access and management.

## Table of Contents

- [Features](#features)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Usage](#usage)
  - [Adding a New Task](#adding-a-new-task)
  - [Listing All Tasks](#listing-all-tasks)
  - [Listing Tasks by Status](#listing-tasks-by-status)
  - [Updating a Task](#updating-a-task)
  - [Deleting a Task](#deleting-a-task)
  - [Marking a Task as In-Progress](#marking-a-task-as-in-progress)
  - [Marking a Task as Done](#marking-a-task-as-done)
- [Data Storage](#data-storage)
- [Contributing](#contributing)
- [License](#license)

---

## Features

- **Add Tasks**: Quickly add new tasks with a simple command.
- **Update Tasks**: Modify existing task descriptions as needed.
- **Delete Tasks**: Remove tasks that are no longer relevant.
- **Mark Tasks**: Change task status to "in-progress" or "done".
- **List Tasks**: View all tasks or filter them by status ("todo", "in-progress", "done").
- **Persistent Storage**: All tasks are saved in a JSON file, ensuring data persists between sessions.

## Prerequisites

- **Go**: You need to have Go installed on your system.
  - **Installation**: Download and install Go from the official website: [golang.org/dl](https://golang.org/dl/).
  - **Verify Installation**:

    ```bash
    go version
    ```

    You should see output similar to:

    ```bash
    go version go1.16.3 darwin/amd64
    ```

## Installation

Follow these steps to set up and run the Task Tracker CLI on your local machine.

### 1. Clone the Repository

Clone the repository to your local machine using Git:

```bash
git clone https://github.com/Ameprizzo/task-tracker-cli.git
cd task-tracker-cli
```

### 2. Build the Application

Compile the application using the Go build tool:

 ```bash
 go build -o task-cli
```

This command will generate an executable file named task-cli in your current directory.

### 3. Run the Application

You can now start using the Task Tracker CLI by running the task-cli executable followed by the desired command.

## Usage

Below are examples of how to use the Task Tracker CLI to manage your tasks.

### Adding a New Task

**Command:**

```bash
./task-cli add "Buy groceries"
```

**Output:**

```bash
Task added successfully (ID: 1)
```

### Listing All Tasks

**Command:**

```bash
./task-cli list
```

**Output:**

```bash
ID: 1
Description: Buy groceries
Status: todo
CreatedAt: 2024-08-21 14:30:00
UpdatedAt: 2024-08-21 14:30:00
```

### Listing Tasks by Status

```bash
./task-cli list todo
```

**Command:**

```bash
./task-cli list todo
```

**Output:**

```bash
ID: 1
Description: Buy groceries
Status: todo
CreatedAt: 2024-08-21 14:30:00
UpdatedAt: 2024-08-21 14:30:00
```

#### Available Status Filters

- **`todo`**
- **`in-progress`**
- **`done`**

### Updating a Task

### Command

```bash
./task-cli update 1 "Buy groceries and cook dinner"
```

**Output:**

``` bash
Task updated successfully (ID: 1)
```

### Deleting a Task

**Command:**

```bash
./task-cli delete 1
```

**Output:**

``` bash
Task deleted successfully (ID: 1)
```

### Marking a Task as In-Progress

**Command:**

```bash
./task-cli mark-in-progress 1
```

**Output:**

```bash
Task marked as in-progress (ID: 1)
```

### Marking a Task as Done

**Command:**

```bash
./task-cli mark-done 1
```

**Output:**

``` bash
Task marked as done (ID: 1)
```

## Data Storage

All tasks are stored in a `tasks.json` file located in the root directory of the project. This file is automatically created when you add your first task.

### Example `tasks.json` Content

```json
[
  {
    "id": 1,
    "description": "Buy groceries and cook dinner",
    "status": "done",
    "createdAt": "2024-08-21T14:30:00Z",
    "updatedAt": "2024-08-21T16:45:00Z"
  },
  {
    "id": 2,
    "description": "Read a book",
    "status": "in-progress",
    "createdAt": "2024-08-21T15:00:00Z",
    "updatedAt": "2024-08-21T15:30:00Z"
  }
]
```

**Note:**
 Ensure you have write permissions in the directory where the application is running so that it can create and modify the `tasks.json` file as needed.

## Contributing

Contributions are welcome! If you'd like to improve this project, follow these steps:

 1. **Fork the Repository:** Click the 'Fork' button on the top right of this page to create a copy of this repository on your GitHub account.
 2. **Clone Your Fork:**

    ```bash
    git clone https://github.com/yourusername/task-tracker-cli.git
    ```

 3. **Create a New Branch:**

    ```bash
    git checkout -b feature/YourFeatureName
    ```

 4. **Make Your Changes:** Implement your feature or bug fix.

    ```bash
    git commit -m "Add YourFeatureName"
    ```

 5. **Commit Your Changes:**

    ```bash
    git commit -m "Add YourFeatureName"
    ```

 6. **Push to Your Fork:**

    ```bash
    git push origin feature/YourFeatureName
    ```
    
7. **Submit a Pull Request:** Go to the original repository and click the 'New Pull Request' button to submit your changes for review.

## License

This project is licensed under the MIT License.

## Contact

If you have any questions, suggestions, or issues, feel free to open an issue on the GitHub repository or contact me directly at <amedeus926@gmail.com>.
