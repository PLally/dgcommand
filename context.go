package dgcommand

import (
	"github.com/bwmarrin/discordgo"
	"github.com/plally/dgcommand/embed"
	log "github.com/sirupsen/logrus"
	"io"
)

type Context interface {
	Error(interface{})

	Reply(string)
	SendFile(name string, r io.Reader)
	SendEmbed(*embed.Embed)

	Args() []string
	SetArgs([]string)

	Message() *discordgo.Message
}


// a struct that other structs can embed to implement the context args functions
type ArgContainer struct {
	args []string
}

func (c *ArgContainer) Args() []string {
	return c.args
}

func (c *ArgContainer) SetArgs(args []string) {
	c.args = args
}


// context object for discordgo
type DiscordContext struct {
	*ArgContainer
	M    *discordgo.MessageCreate
	S    *discordgo.Session
}

func (ctx *DiscordContext) Reply(msg string) {
	ctx.S.ChannelMessageSend(ctx.M.ChannelID, msg)
}

func (ctx *DiscordContext) Error(err interface{}) {
	log.Errorf("error on message %v\n%v", ctx.M.Content, err)
	ctx.Reply("Encountered an unhandled error")
}

func (ctx *DiscordContext) SendFile(name string, r io.Reader) {
	ctx.S.ChannelFileSend(ctx.M.ChannelID, name, r)
}

func (ctx *DiscordContext) SendEmbed(e *embed.Embed) {
	ctx.S.ChannelMessageSendEmbed(ctx.M.ChannelID, e.MessageEmbed)
}

func (ctx *DiscordContext) Message() *discordgo.Message {
	return ctx.M.Message
}

func CreateDiscordContext(s *discordgo.Session, m *discordgo.MessageCreate) *DiscordContext {
	return &DiscordContext{
		ArgContainer: &ArgContainer{args: []string{
			m.Content,
		}},
		M:           m,
		S:           s,
	}
}
