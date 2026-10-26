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
	req := ReadRequest[R]{
		BaseRequest: parseRequest[R](r),
		ID:          r.PathValue("id"),
	}
	h(req).write(w)
}

func (h VReadHandler[R]) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	req := VReadRequest[R]{
		BaseRequest: parseRequest[R](r),
		ID:          r.PathValue("id"),
		VID:         r.PathValue("vid"),
	}
	h(req).write(w)
}

func (h UpdateHandler[R]) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	req := UpdateRequest[R]{
		BaseRequest: parseRequest[R](r),
		ID:          r.PathValue("id"),
	}
	h(req).write(w)
}

func (h PatchHandler[R]) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	req := PatchRequest[R]{
		BaseRequest: parseRequest[R](r),
		ID:          r.PathValue("id"),
	}
	h(req).write(w)
}

func (h DeleteHandler[R]) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	req := DeleteRequest[R]{
		BaseRequest: parseRequest[R](r),
		ID:          r.PathValue("id"),
	}
	h(req).write(w)
}

func (h CreateHandler[R]) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	req := CreateRequest[R]{
		BaseRequest: parseRequest[R](r),
	}
	h(req).write(w)
}

func (h SearchHandler[R]) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	req := SearchRequest[R]{
		BaseRequest: parseRequest[R](r),
	}
	h(req).write(w)
}
