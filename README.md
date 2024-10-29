# task-tracker
project for [roadmap.sh](https://roadmap.sh/projects/task-tracker) task tracker

### How to Run
Clone the repository and run the following command:

```bash
git clone https://github.com/joiller/task-tracker.git
    
cd task-tracker
```

Run the following command to build and run the project:

```bash
go build -o task-tracker
./task-tracker --help # To see the list of available commands

# To add a task
./task-tracker add "Buy groceries"

# To update a task
./task-tracker update 1 "Buy groceries and cook dinner"

# To delete a task
./task-tracker delete 1 [2 3 4]

# To mark a task as in progress/done/todo
./task-tracker mark-in-progress 1 [2 3 4]
./task-tracker mark-done 1 [2 3 4]
./task-tracker mark-todo 1 [2 3 4]

# To list all tasks
./task-tracker list
./task-tracker list done
./task-tracker list todo
./task-tracker list in-progress
```
