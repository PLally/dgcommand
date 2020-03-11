package dgcommand

import "strings"

func NSFWOnly(h Handler) HandlerFunc {
	return func(ctx Context) {
		channel, err := ctx.S.State.Channel(ctx.M.ChannelID)
		if err != nil {
			ctx.Error(err)
			return
		}
		if !channel.NSFW {
			ctx.Reply("This command must be used in an nsfw channel")
			return
		}
		h.Handle(ctx)
	}
}
func Prefix(getPrefix func(Context) string, next HandlerFunc) HandlerFunc {
	return func(ctx Context) {
		prefix := getPrefix(ctx)
		if !strings.HasPrefix(ctx.Args[0], prefix) {
			return
		}
		ctx.Args[0] = ctx.Args[0][len(prefix):]
		next(ctx)
	}

}
