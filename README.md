### Todo list service
(simple CRUD functionality, postgresql as a database)


Interfaces:
- GRPC
- HTTP (GRPC-gateway)
- Telegram bot

Telegram bot commands:

- `/help` - list all commands
- `/get <task_id>` - get task with id
- `/list` - list all current tasks
- `/update <task_id> <due_date:YYYY-MM-DD> <task>` - update task with id
- `/delete <task_id>` - delete task with id
- `/add <due_date:YYYY-MM-DD> <task>` - add a new task
