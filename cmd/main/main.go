package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"gitlab.ozon.dev/ralexa2000/todo-bot/internal/config"
	botPkg "gitlab.ozon.dev/ralexa2000/todo-bot/internal/pkg/bot"
	cmdAddPkg "gitlab.ozon.dev/ralexa2000/todo-bot/internal/pkg/bot/command/add"
	cmdDeletePkg "gitlab.ozon.dev/ralexa2000/todo-bot/internal/pkg/bot/command/delete"
	cmdGetPkg "gitlab.ozon.dev/ralexa2000/todo-bot/internal/pkg/bot/command/get"
	cmdHelpPkg "gitlab.ozon.dev/ralexa2000/todo-bot/internal/pkg/bot/command/help"
	cmdListPkg "gitlab.ozon.dev/ralexa2000/todo-bot/internal/pkg/bot/command/list"
	cmdUpdatePkg "gitlab.ozon.dev/ralexa2000/todo-bot/internal/pkg/bot/command/update"
	taskPkg "gitlab.ozon.dev/ralexa2000/todo-bot/internal/pkg/core/task"
	"gitlab.ozon.dev/ralexa2000/todo-bot/internal/pkg/repository"
	"log"
)

var psqlConn = fmt.Sprintf(
	"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
	config.DbHost,
	config.DbPort,
	config.DbUser,
	config.DbPwd,
	config.DbName,
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	pool := connectToDb(ctx)
	defer pool.Close()

	repo := repository.New(pool)

	var task taskPkg.Interface
	{
		task = taskPkg.New(ctx, repo)
	}

	bot := registerBot(task)
	go runBot(bot)
	go runREST(ctx)
	runGRPCServer(task)
}

func connectToDb(ctx context.Context) *pgxpool.Pool {
	pool, err := pgxpool.Connect(ctx, psqlConn)
	if err != nil {
		log.Fatal("can't connect to database", err)
	}

	if err := pool.Ping(ctx); err != nil {
		log.Fatal("ping database error", err)
	}
	return pool
}

func runBot(bot botPkg.Interface) {
	if err := bot.Run(); err != nil {
		log.Fatal(err)
	}
}

func registerBot(task taskPkg.Interface) botPkg.Interface {
	var bot botPkg.Interface
	{
		bot = botPkg.MustNew()

		commandAdd := cmdAddPkg.New(task)
		bot.RegisterHandler(commandAdd)

		commandGet := cmdGetPkg.New(task)
		bot.RegisterHandler(commandGet)

		commandList := cmdListPkg.New(task)
		bot.RegisterHandler(commandList)

		commandUpdate := cmdUpdatePkg.New(task)
		bot.RegisterHandler(commandUpdate)

		commandDelete := cmdDeletePkg.New(task)
		bot.RegisterHandler(commandDelete)

		commandHelp := cmdHelpPkg.New(map[string][2]string{
			commandAdd.Name():    {commandAdd.Arguments(), commandAdd.Description()},
			commandGet.Name():    {commandGet.Arguments(), commandGet.Description()},
			commandList.Name():   {commandList.Arguments(), commandList.Description()},
			commandUpdate.Name(): {commandUpdate.Arguments(), commandUpdate.Description()},
			commandDelete.Name(): {commandDelete.Arguments(), commandDelete.Description()},
		})
		bot.RegisterHandler(commandHelp)
	}
	return bot
}
