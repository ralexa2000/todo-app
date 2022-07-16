package main

import (
	"gitlab.ozon.dev/ralexa2000/todo-bot/internal/commander"
	"log"
)

func main() {
	c, err := commander.Init()
	if err != nil {
		log.Panic(err)
	}
	if err := c.Run(); err != nil {
		log.Panic(err)
	}
}
