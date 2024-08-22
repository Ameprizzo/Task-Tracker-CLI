# v0.1.0

You can now start using the Task Tracker CLI by running the task-cli executable followed by the desired command.

## Usage

Below are examples of how to use the Task Tracker CLI to manage your tasks.

### Adding a New Task

**Command:**

```bash
tasktracker add "Buy groceries"
```

**Output:**

```bash
Task added successfully (ID: 1)
```

### Listing All Tasks

**Command:**

```bash
tasktracker list
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
tasktracker list todo
```

**Command:**

```bash
tasktracker list todo
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
tasktracker update 1 "Buy groceries and cook dinner"
```

**Output:**

``` bash
Task updated successfully (ID: 1)
```

### Deleting a Task

**Command:**

```bash
tasktracker delete 1
```

**Output:**

``` bash
Task deleted successfully (ID: 1)
```

### Marking a Task as In-Progress

**Command:**

```bash
tasktracker mark-in-progress 1
```

**Output:**

```bash
Task marked as in-progress (ID: 1)
```

### Marking a Task as Done

**Command:**

```bash
tasktracker mark-done 1
```

**Output:**

``` bash
Task marked as done (ID: 1)
```
