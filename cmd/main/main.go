package main

import (
	"gitlab.ozon.dev/ralexa2000/todo-bot/internal/commander"
	"gitlab.ozon.dev/ralexa2000/todo-bot/internal/handlers"
	"log"
)

func main() {
	c, err := commander.Init()
	if err != nil {
		log.Panic(err)
	}
	handlers.AddHandlers(c)
	if err := c.Run(); err != nil {
		log.Panic(err)
	}
}
