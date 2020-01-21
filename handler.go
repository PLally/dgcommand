package dgcommand

import (
	"github.com/bwmarrin/discordgo"
	"github.com/plally/dgcommand/parsing"
)

type CommandContext struct {
	M *discordgo.MessageCreate
	Args    []string
	S *discordgo.Session
}

type Handler interface {
	Handle(ctx CommandContext)
}

// command router. routes commands based on the first word
type CommandRoutingHandler struct {
	commands map[string]Handler
}

func NewCommandHandler() *CommandRoutingHandler {
	return &CommandRoutingHandler{
		commands: make(map[string]Handler),
	}
}

func (h *CommandRoutingHandler) DiscordHandle(s  *discordgo.Session, m *discordgo.MessageCreate) {

	h.Handle( CommandContext{
		S: s,
		M: m,
		Args: []string{
			m.Content,
		},
	} )
}

//command_name [args...]
func (h *CommandRoutingHandler) Handle(ctx CommandContext) {
	first := firstWord(ctx.Args[0])
	text := ctx.Args[0][len(first):]
	ctx.Args = []string{
		first,
		text,
	}

	var cmd string
	if len(ctx.Args) > 1 {
		cmd = ctx.Args[0]
	}



	handler, ok := h.commands[cmd]
	if !ok { return }
	ctx.Args = ctx.Args[1:]
	handler.Handle(ctx)
}



func (h *CommandRoutingHandler) AddHandler(name string, handler Handler) {
	h.commands[name] = handler
}

// command
type Command struct {
	Callback func(ctx CommandContext)
	parsing.CommandDefinition
}

func (cmd *Command) Handle(ctx CommandContext) {
	var err error
	ctx.Args, err = cmd.ParseInput(ctx.Args[0])
	if err != nil {
		return
	}
	cmd.Callback(ctx)
}
func NewCommand(definition string, callback func(ctx CommandContext)) *Command {
	return &Command{
		Callback: callback,
		CommandDefinition: parsing.NewCommandDefinition(definition),
	}
}