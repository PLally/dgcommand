package dgcommand

type HandlerMeta struct {
	Description string
	DisplayName string
}

func (h *HandlerMeta) Metadata() *HandlerMeta {
	return h
}
func (h *HandlerMeta) Desc(text string) {
	h.Description = text
}

func (h *HandlerMeta) Display(text string) {
	h.DisplayName = text
}

