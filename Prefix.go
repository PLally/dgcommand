package dgcommand

import "strings"



func OnPrefix(prefix string, next Handler) HandlerFunc {
	return func(ctx Context) {
		args := ctx.Args()
		if !strings.HasPrefix(args[0], prefix) {
			return
		}
		args[0] = args[0][len(prefix):]
		next.Handle(ctx)
	}

}