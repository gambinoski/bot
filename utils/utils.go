package utils

/*
* Most of this file has been made by @pollen5 on GitHub.
 */
import (
	"fmt"
	"math/rand"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/gambinoski/bot/config"
	"github.com/pollen5/minori"
	"github.com/sapphire-cord/sapphire"
)

// Logger is a minori log
var Logger = minori.GetLogger("Bot")

var Rand = rand.New(rand.NewSource(time.Now().Unix()))

func RandNumber(r int) int {
	return Rand.Intn(r)
}

func Roll(ls []string) string {
	return ls[RandNumber(len(ls))]
}

// ErrorHandler send nice info into a private LogChannel
func ErrorHandler(bot *sapphire.Bot, err interface{}) {
	if cmd, ok := err.(*sapphire.CommandError); ok {
		ctx := cmd.Context

		g := "DM"
		gid := ctx.Channel.ID

		if ctx.Guild != nil {
			g = ctx.Guild.Name
			gid = ctx.Guild.ID
		}

		bot.Session.ChannelMessageSendEmbed(config.Conf.Bot.LogChannel, &discordgo.MessageEmbed{
			Title:       "Command Error - " + ctx.Command.Name,
			Description: fmt.Sprintf("Content: %s\n```\n%s```", ctx.Message.Content, cmd.Error()),
			Color:       0xDFAC7C,
			Footer: &discordgo.MessageEmbedFooter{
				Text: fmt.Sprintf("Author: %s (%s), Guild: %s (%s)", ctx.Author.Username, ctx.Author.ID, g, gid),
			},
		})

		Logger.Errorf("Command Error (%s): %s (Author: %s (%s), Guild: %s (%s))", ctx.Command.Name, cmd.Error(), ctx.Author.Username, ctx.Author.ID, g, gid)
		ctx.ReplyLocale("COMMAND_ERROR")
		return
	}

	bot.Session.ChannelMessageSendEmbed(config.Conf.Bot.LogChannel, &discordgo.MessageEmbed{
		Title:       "Panic Recovered",
		Description: fmt.Sprint(err),
		Color:       0xDFAC7C,
	})

	Logger.Error(err)
}
