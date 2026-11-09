package fire

import (
	"encoding/json/v2"
	"errors"
	"io"
	"net/http"
)

type Resource interface {
	json.Marshaler
	ResourceType() string
}

type (
	ReadHandler[R Resource]        func(ReadRequest[R]) Response
	VReadHandler[R Resource]       func(VReadRequest[R]) Response
	UpdateHandler[R Resource]      func(UpdateRequest[R]) Response
	PatchHandler[R Resource]       func(PatchRequest[R]) Response
	DeleteHandler[R Resource]      func(DeleteRequest[R]) Response
	CreateHandler[R Resource]      func(CreateRequest[R]) Response
	SearchHandler[R Resource]      func(SearchRequest[R]) Response
	InstanceExtHandler[R Resource] func(InstanceExtRequest[R]) Response
	TypeExtHandler[R Resource]     func(TypeExtRequest[R]) Response
)

func (h ReadHandler[R]) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	base := parseRequest[R](r)
	req := ReadRequest[R]{
		BaseRequest: base,
		ID:          r.PathValue("id"),
	}
	resp := h(req)
	resp.write(req.BaseRequest, w)
}

func (h VReadHandler[R]) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	base := parseRequest[R](r)
	req := VReadRequest[R]{
		BaseRequest: base,
		ID:          r.PathValue("id"),
		VID:         r.PathValue("vid"),
	}
	resp := h(req)
	resp.write(req.BaseRequest, w)
}

func (h UpdateHandler[R]) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	raw, err := parseRawBody(r)
	if err != nil {
		sendErr(w, err)
		return
	}
	base := parseRequest[R](r)
	body, err := parseBody(base, raw)
	if err != nil {
		sendErr(w, err)
		return
	}
	req := UpdateRequest[R]{
		BaseRequest: base,
		ID:          r.PathValue("id"),
		Body:        body,
		RawBody:     raw,
	}
	resp := h(req)
	resp.write(req.BaseRequest, w)
}

func (h PatchHandler[R]) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	raw, err := parseRawBody(r)
	if err != nil {
		sendErr(w, err)
		return
	}
	base := parseRequest[R](r)
	body, err := parseBody(base, raw)
	if err != nil {
		sendErr(w, err)
		return
	}
	req := PatchRequest[R]{
		BaseRequest: base,
		ID:          r.PathValue("id"),
		Body:        body,
		RawBody:     raw,
	}
	resp := h(req)
	resp.write(req.BaseRequest, w)
}

func (h DeleteHandler[R]) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	base := parseRequest[R](r)
	req := DeleteRequest[R]{
		BaseRequest: base,
		ID:          r.PathValue("id"),
	}
	resp := h(req)
	resp.write(req.BaseRequest, w)
}

func (h CreateHandler[R]) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	raw, err := parseRawBody(r)
	if err != nil {
		sendErr(w, err)
		return
	}
	base := parseRequest[R](r)
	body, err := parseBody(base, raw)
	if err != nil {
		sendErr(w, err)
		return
	}
	req := CreateRequest[R]{
		BaseRequest: base,
		Body:        body,
		RawBody:     raw,
	}
	resp := h(req)
	resp.write(req.BaseRequest, w)
}

func (h SearchHandler[R]) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	base := parseRequest[R](r)
	req := SearchRequest[R]{
		BaseRequest: base,
		Body:        r.Form,
	}
	resp := h(req)
	resp.write(req.BaseRequest, w)
}

func (h InstanceExtHandler[R]) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	base := parseRequest[R](r)
	req := InstanceExtRequest[R]{
		BaseRequest: base,
		ID:          r.PathValue("id"),
	}
	resp := h(req)
	resp.write(req.BaseRequest, w)
}

func (h TypeExtHandler[R]) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	base := parseRequest[R](r)
	req := TypeExtRequest[R]{
		BaseRequest: base,
	}
	resp := h(req)
	resp.write(req.BaseRequest, w)
}

func parseRawBody(r *http.Request) ([]byte, error) {
	const maxFormSize = int64(10 << 20) // 10 MB is a lot of text.
	reader := io.LimitReader(r.Body, maxFormSize+1)
	return io.ReadAll(reader)
}

func parseBody[R Resource](r *BaseRequest[R], raw []byte) (R, error) {
	var body R
	var err error
	switch r.ContentType {
	case ContentTypeForm:
		return body, errors.New("unsupported Content-Type")
	case ContentTypeJSON:
		err = json.Unmarshal(raw, &body)
	case ContentTypeTurtle:
		return body, errors.New("unsupported Content-Type")
	case ContentTypeXML:
		return body, errors.New("unsupported Content-Type")
	default:
		return body, errors.New("unsupported Content-Type")
	}
	return body, err
}

func sendErr(w http.ResponseWriter, err error) {
	http.Error(w, err.Error(), http.StatusBadRequest)
}
