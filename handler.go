package dgcommand

type Handler interface {
	Handle(ctx CommandContext)
}

type HandlerFunc func(ctx CommandContext)

type MiddlewareFunc func(h HandlerFunc) HandlerFunc

func (h HandlerFunc) Handle(ctx CommandContext) {
	h(ctx)
}
