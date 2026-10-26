package fire

import "net/http"

type Handler func(Request) Response

type Operation struct {
	Verb        Verb
	Path        string
	Interaction Interaction
	Type        Type
	Handler     Handler
}

func (op *Operation) Register(mux *http.ServeMux) {
	mux.HandleFunc(string(op.Verb)+" "+op.Path, op.Handler)
}
