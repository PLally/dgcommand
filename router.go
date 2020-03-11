package dgcommand

import (
	"github.com/plally/dgcommand/parsing/util"
)

/*

 */
type CommandRoutingHandler struct {
	Routes map[string]Handler
	Middleware []MiddlewareFunc
	Handler HandlerFunc
}

func NewRouter() *CommandRoutingHandler {
	return &CommandRoutingHandler{
		Routes: make(map[string]Handler),
		Middleware: make([]MiddlewareFunc, 0),
	}
}

func (h *CommandRoutingHandler) AddHandler(name string, handler Handler) {
	h.Routes[name] = handler
}

func (r *CommandRoutingHandler) On(name string, f HandlerFunc) {
	r.AddHandler(name, f)
}

func (h *CommandRoutingHandler) Handle(ctx Context) {
	first := util.FirstWord(ctx.Args[0])
	text := ctx.Args[0][len(first):]
	util.ConsumeSpaces(&text)

	ctx.Args = []string{
		first,
		text,
	}

	handler, ok := h.Routes[ctx.Args[0]]
	if !ok { return }
	ctx.Args = ctx.Args[1:]
	handler.Handle(ctx)
}


func main() {
	r := NewRouter()

	getPrefix := func(ctx Context) string { return ">" }
	Prefix(getPrefix, func(ctx Context){
		r.Handle(ctx)
	})

	r.On("ping", func(ctx Context) {
		ctx.Reply("Pong")
	})

	r.On("Nsfw", NSFWOnly(func(ctx Context){

	}))
}