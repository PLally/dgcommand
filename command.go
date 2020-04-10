package dgcommand

import (
	"github.com/plally/dgcommand/parsing"
)

type Command struct {
	HandlerMeta
	Callback HandlerFunc
	parsing.CommandDefinition
}

func (cmd *Command) Handle(ctx Context) {
	args, err := cmd.ParseInput(ctx.Args()[0])
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.SetArgs(args)
	cmd.Callback(ctx)
}

func (cmd *Command) Use(middleware ...MiddlewareFunc) *Command {
	for _, mid := range middleware {
		cmd.Callback = mid(cmd.Callback)
	}
	return cmd
}
func NewCommand(definition string, callback HandlerFunc) *Command {
	return &Command{
		Callback:          callback,
		CommandDefinition: parsing.NewCommandDefinition(definition),
	}
}
