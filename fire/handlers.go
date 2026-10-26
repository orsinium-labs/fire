package fire

import (
	"encoding/json/v2"
	"net/http"
)

type Resource json.Marshaler

type (
	ReadHandler[R Resource]   func(ReadRequest[R]) Response
	VReadHandler[R Resource]  func(VReadRequest[R]) Response
	UpdateHandler[R Resource] func(UpdateRequest[R]) Response
	PatchHandler[R Resource]  func(PatchRequest[R]) Response
	DeleteHandler[R Resource] func(DeleteRequest[R]) Response
	CreateHandler[R Resource] func(CreateRequest[R]) Response
	SearchHandler[R Resource] func(SearchRequest[R]) Response
)

func (h ReadHandler[R]) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	base := &BaseRequest[R]{HTTPRequest: r}
	err := base.parse()
	if err != nil {
		writeErr(w, err)
		return
	}
	req := ReadRequest[R]{
		BaseRequest: base,
		ID:          r.PathValue("id"),
	}
	sendResp(w, h(req))
}

func (h VReadHandler[R]) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	base := &BaseRequest[R]{HTTPRequest: r}
	err := base.parse()
	if err != nil {
		writeErr(w, err)
		return
	}
	req := VReadRequest[R]{
		BaseRequest: base,
		ID:          r.PathValue("id"),
		VID:         r.PathValue("vid"),
	}
	sendResp(w, h(req))
}

func (h UpdateHandler[R]) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	base := &BaseRequest[R]{HTTPRequest: r}
	err := base.parse()
	if err != nil {
		writeErr(w, err)
		return
	}
	req := UpdateRequest[R]{
		BaseRequest: base,
		ID:          r.PathValue("id"),
	}
	sendResp(w, h(req))
}

func (h PatchHandler[R]) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	base := &BaseRequest[R]{HTTPRequest: r}
	err := base.parse()
	if err != nil {
		writeErr(w, err)
		return
	}
	req := PatchRequest[R]{
		BaseRequest: base,
		ID:          r.PathValue("id"),
	}
	sendResp(w, h(req))
}

func (h DeleteHandler[R]) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	base := &BaseRequest[R]{HTTPRequest: r}
	err := base.parse()
	if err != nil {
		writeErr(w, err)
		return
	}
	req := DeleteRequest[R]{
		BaseRequest: base,
		ID:          r.PathValue("id"),
	}
	sendResp(w, h(req))
}

func (h CreateHandler[R]) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	base := &BaseRequest[R]{HTTPRequest: r}
	err := base.parse()
	if err != nil {
		writeErr(w, err)
		return
	}
	req := CreateRequest[R]{
		BaseRequest: base,
	}
	sendResp(w, h(req))
}

func (h SearchHandler[R]) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	base := &BaseRequest[R]{HTTPRequest: r}
	err := base.parse()
	if err != nil {
		writeErr(w, err)
		return
	}
	req := SearchRequest[R]{
		BaseRequest: base,
	}
	sendResp(w, h(req))
}

func sendResp(w http.ResponseWriter, resp Response) {
	panic("todo")
}

func writeErr(w http.ResponseWriter, err error) {
	panic("todo")
}
