package main

import (
	botPkg "gitlab.ozon.dev/ralexa2000/todo-bot/internal/pkg/bot"
	cmdAddPkg "gitlab.ozon.dev/ralexa2000/todo-bot/internal/pkg/bot/command/add"
	cmdDeletePkg "gitlab.ozon.dev/ralexa2000/todo-bot/internal/pkg/bot/command/delete"
	cmdHelpPkg "gitlab.ozon.dev/ralexa2000/todo-bot/internal/pkg/bot/command/help"
	cmdListPkg "gitlab.ozon.dev/ralexa2000/todo-bot/internal/pkg/bot/command/list"
	cmdUpdatePkg "gitlab.ozon.dev/ralexa2000/todo-bot/internal/pkg/bot/command/update"
	taskPkg "gitlab.ozon.dev/ralexa2000/todo-bot/internal/pkg/core/task"
	"log"
)

func main() {
	var task taskPkg.Interface
	{
		task = taskPkg.New()
	}

	var bot botPkg.Interface
	{
		bot = botPkg.MustNew()

		commandAdd := cmdAddPkg.New(task)
		bot.RegisterHandler(commandAdd)

		commandList := cmdListPkg.New(task)
		bot.RegisterHandler(commandList)

		commandUpdate := cmdUpdatePkg.New(task)
		bot.RegisterHandler(commandUpdate)

		commandDelete := cmdDeletePkg.New(task)
		bot.RegisterHandler(commandDelete)

		commandHelp := cmdHelpPkg.New(map[string][2]string{
			commandAdd.Name():    {commandAdd.Arguments(), commandAdd.Description()},
			commandList.Name():   {commandList.Arguments(), commandList.Description()},
			commandUpdate.Name(): {commandUpdate.Arguments(), commandUpdate.Description()},
			commandDelete.Name(): {commandDelete.Arguments(), commandDelete.Description()},
		})
		bot.RegisterHandler(commandHelp)
	}

	if err := bot.Run(); err != nil {
		log.Fatal(err)
	}
}
