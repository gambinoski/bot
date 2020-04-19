package main

import (
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/gambinoski/bot/config"
	"github.com/gambinoski/bot/database"
	"github.com/gambinoski/bot/discord/commands"
	"github.com/gambinoski/bot/discord/events"
	"github.com/sapphire-cord/sapphire"
)

func main() {
	config.Load()
	var token = config.Conf.Bot.Token

	if token == "" {
		log.Fatalln("No $BOT_TOKEN given.")
	}
	dg, err := discordgo.New(token)
	if err != nil {
		log.Fatal(err)
	}

	bot := sapphire.New(dg)

	bot.OwnerID = config.Conf.Bot.OwnerID

	database.Connect()
	database.Setup()

	// Initialize
	bot.LoadBuiltins()
	commands.Init(bot)
	events.Init(dg)

	bot.MustConnect()
	bot.Wait()
}
