package dgcommand

import "github.com/plally/dgcommand/parsing/util"

type CommandGroup struct {
	HandlerMeta
	Commands map[string]Handler
}

func Group() *CommandGroup {
	return &CommandGroup{
		Commands: make(map[string]Handler),
	}
}

func (h *CommandGroup) AddHandler(name string, handler Handler) {
	h.Commands[name] = handler
}

func (r *CommandGroup) On(name string, f HandlerFunc) {
	r.AddHandler(name, f)
}

func (r *CommandGroup) Command(definition string, f HandlerFunc) *Command {
	cmd := NewCommand(definition, f)
	r.AddHandler(cmd.Name, cmd)
	return cmd
}

func (h *CommandGroup) Handle(ctx Context) {
	args := ctx.Args()

	first := util.FirstWord(args[0])
	text := args[0][len(first):]
	util.ConsumeSpaces(&text)

	args = []string{
		first,
		text,
	}

	handler, ok := h.Commands[args[0]]
	if !ok { return }
	ctx.SetArgs(args[1:])
	handler.Handle(ctx)
}
