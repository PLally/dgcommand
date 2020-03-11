package dgcommand

import (
	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
)

type Context struct {
	M    *discordgo.MessageCreate
	Args []string
	S    *discordgo.Session
}

func (ctx *Context) Reply(msg string) {
	ctx.S.ChannelMessageSend(ctx.M.ChannelID, msg)
}

func (ctx *Context) Error(err interface{}) {
	log.Error("error on message %v\n%v", ctx.M.Content, err)
	ctx.Reply("Encountered an unhandled error")
}
// TODO include previous handler in command context
