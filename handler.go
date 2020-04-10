package dgcommand

type Handler interface {
	Handle(ctx Context)
}

type HandlerFunc func(ctx Context)

type MiddlewareFunc func(h HandlerFunc) HandlerFunc

func (h HandlerFunc) Handle(ctx Context) {
	h(ctx)
}
