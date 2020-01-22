package dgcommand

import "github.com/bwmarrin/discordgo"

type CommandContext struct {
	M *discordgo.MessageCreate
	Args    []string
	S *discordgo.Session
}

func (ctx *CommandContext) Reply(msg string) {
	ctx.S.ChannelMessageSend(ctx.M.ChannelID, msg)
}