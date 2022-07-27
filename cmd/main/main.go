package main

import (
	botPkg "gitlab.ozon.dev/ralexa2000/todo-bot/internal/pkg/bot"
	cmdAddPkg "gitlab.ozon.dev/ralexa2000/todo-bot/internal/pkg/bot/command/add"
	cmdHelpPkg "gitlab.ozon.dev/ralexa2000/todo-bot/internal/pkg/bot/command/help"
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

		commandHelp := cmdHelpPkg.New(map[string][2]string{
			commandAdd.Name(): {commandAdd.Arguments(), commandAdd.Description()},
		})
		bot.RegisterHandler(commandHelp)
	}

	if err := bot.Run(); err != nil {
		log.Fatal(err)
	}
}
