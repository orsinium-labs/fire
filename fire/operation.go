package fire

import "net/http"

type Operation struct {
	Verb        Verb
	Path        string
	Interaction Interaction
	Type        Type
	Handler     http.Handler
}

func (op *Operation) Register(mux *http.ServeMux) {
	mux.Handle(string(op.Verb)+" "+op.Path, op.Handler)
}
