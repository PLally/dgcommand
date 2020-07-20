package dgcommand

import (
	"context"
	"github.com/bwmarrin/discordgo"
	"github.com/plally/dgcommand/embed"
	log "github.com/sirupsen/logrus"
	"io"
)

type CommandContext struct {
	context.Context
	Session *discordgo.Session
	Message *discordgo.MessageCreate
	args []string
}

func (ctx *CommandContext) Args() []string {
	return ctx.args
}

func (ctx *CommandContext) WithValue(k, v interface{}) {
	ctx.Context = context.WithValue(ctx.Context, k, v)
}
func (ctx *CommandContext) SetArgs(args []string) {
	ctx.args = args
}


// context object for discordgo
func (ctx *CommandContext) Reply(msg string) {
	ctx.Session.ChannelMessageSend(ctx.Message.ChannelID, msg)
}

func (ctx *CommandContext) Error(err interface{}) {
	log.Errorf("error on message %v\n%v", ctx.Message.Content, err)
	ctx.Reply("Encountered an unhandled error")
}

func (ctx *CommandContext) SendFile(name string, r io.Reader) {
	ctx.Session.ChannelFileSend(ctx.Message.ChannelID, name, r)
}

func (ctx *CommandContext) SendEmbed(e *embed.Embed) {
	ctx.Session.ChannelMessageSendEmbed(ctx.Message.ChannelID, e.MessageEmbed)
}

func CreatContext(s *discordgo.Session, m *discordgo.MessageCreate) CommandContext{
	ctx := CommandContext{
		Context: context.Background(),
		Session: s,
		Message: m,
	}
	ctx.SetArgs([]string{m.Content})
	return ctx
}
