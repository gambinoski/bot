// Package commands is the main entry point where all commands are registered.
// Auto-Generated with spgen DO NOT EDIT.
// To use this file import the package in your entry file and initialize it with commands.Init(bot)
package commands

import (
	"github.com/gambinoski/bot/discord/commands/general"
	"github.com/sapphire-cord/sapphire"
)

func Init(bot *sapphire.Bot) {
	bot.AddCommand(sapphire.NewCommand("ping", "General", general.Ping))
}
