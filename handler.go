package dgcommand

import (
	"github.com/bwmarrin/discordgo"
	"github.com/plally/dgcommand/parsing"
)

type Handler interface {
	Handle(ctx Context)
}

type HandlerFunc func(ctx Context)

type MiddlewareFunc func(Handler) Handler

func (h HandlerFunc) Handle(ctx Context) {
	h(ctx)
}

func DiscordHandle(h Handler) func(*discordgo.Session, *discordgo.MessageCreate) {

	return func(s *discordgo.Session, m *discordgo.MessageCreate) {
		h.Handle(Context{
			S: s,
			M: m,
			Args: []string{
				m.Content,
			},
		})
	}
}

// command
type Command struct {
	Callback func(ctx Context)
	parsing.CommandDefinition
}

func (cmd *Command) Handle(ctx Context) {
	var err error
	ctx.Args, err = cmd.ParseInput(ctx.Args[0])
	if err != nil {
		return
	}
	cmd.Callback(ctx)
}
func NewCommand(definition string, callback func(ctx Context)) *Command {
	return &Command{
		Callback:          callback,
		CommandDefinition: parsing.NewCommandDefinition(definition),
	}
}
