package dgcommand

import (
	"github.com/bwmarrin/discordgo"
	"github.com/plally/dgcommand/parsing"
	"github.com/plally/dgcommand/parsing/util"
	"strings"
)

type Handler interface {
	Handle(ctx CommandContext)
}

type HandlerFunc func(ctx CommandContext)

func (h HandlerFunc) Handle(ctx CommandContext) {
	h(ctx)
}

// prefix handler
type PrefixHandler struct {
	Prefix   func(ctx CommandContext) string
	Callback func(ctx CommandContext)
}

func DiscordHandle(h Handler) func(*discordgo.Session, *discordgo.MessageCreate) {

	return func(s *discordgo.Session, m *discordgo.MessageCreate) {
		h.Handle(CommandContext{
			S: s,
			M: m,
			Args: []string{
				m.Content,
			},
		})
	}
}

func (h *PrefixHandler) Handle(ctx CommandContext) {
	prefix := h.Prefix(ctx)
	if !strings.HasPrefix(ctx.Args[0], prefix) {
		return
	}
	ctx.Args[0] = ctx.Args[0][len(prefix):]
	h.Callback(ctx)
}

func WithPrefix(h Handler, getPrefix func(ctx CommandContext) string) *PrefixHandler {
	return &PrefixHandler{
		Prefix:   getPrefix,
		Callback: h.Handle,
	}
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

func (h *CommandRoutingHandler) Commands() map[string]Handler {
	return h.commands
}

//command_name [args...]
func (h *CommandRoutingHandler) Handle(ctx CommandContext) {
	first := util.FirstWord(ctx.Args[0])
	text := ctx.Args[0][len(first):]
	util.ConsumeSpaces(&text)

	ctx.Args = []string{
		first,
		text,
	}

	var cmd string
	if len(ctx.Args) > 1 {
		cmd = ctx.Args[0]
	}

	handler, ok := h.commands[cmd]
	if !ok {
		return
	}
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
		Callback:          callback,
		CommandDefinition: parsing.NewCommandDefinition(definition),
	}
}
