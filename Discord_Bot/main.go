package main

import (
	"fmt"

	"github.com/rohitvarshney/GoPrac/discord_bot/bot"
	"github.com/rohitvarshney/GoPrac/discord_bot/config"
)

func main() {
	err := config.ReadConfig()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	bot.Start()

	<-make(chan struct{})
	return
}
