package fire

import "net/http"

type (
	ReadHandler   func(ReadRequest) Response
	VReadHandler  func(VReadRequest) Response
	UpdateHandler func(UpdateRequest) Response
	PatchHandler  func(PatchRequest) Response
	DeleteHandler func(DeleteRequest) Response
	CreateHandler func(CreateRequest) Response
	SearchHandler func(SearchRequest) Response
)

func (h ReadHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	base := &BaseRequest{HTTPRequest: r}
	err := base.Parse()
	if err != nil {
		writeErr(w, err)
		return
	}
	req := ReadRequest{
		BaseRequest: base,
		ID:          r.PathValue("id"),
	}
	sendResp(w, h(req))
}

func (h VReadHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	base := &BaseRequest{HTTPRequest: r}
	err := base.Parse()
	if err != nil {
		writeErr(w, err)
		return
	}
	req := VReadRequest{
		BaseRequest: base,
		ID:          r.PathValue("id"),
		VID:         r.PathValue("vid"),
	}
	sendResp(w, h(req))
}

func (h UpdateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	base := &BaseRequest{HTTPRequest: r}
	err := base.Parse()
	if err != nil {
		writeErr(w, err)
		return
	}
	req := UpdateRequest{
		BaseRequest: base,
		ID:          r.PathValue("id"),
	}
	sendResp(w, h(req))
}

func (h PatchHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	base := &BaseRequest{HTTPRequest: r}
	err := base.Parse()
	if err != nil {
		writeErr(w, err)
		return
	}
	req := PatchRequest{
		BaseRequest: base,
		ID:          r.PathValue("id"),
	}
	sendResp(w, h(req))
}

func (h DeleteHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	base := &BaseRequest{HTTPRequest: r}
	err := base.Parse()
	if err != nil {
		writeErr(w, err)
		return
	}
	req := DeleteRequest{
		BaseRequest: base,
		ID:          r.PathValue("id"),
	}
	sendResp(w, h(req))
}

func (h CreateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	base := &BaseRequest{HTTPRequest: r}
	err := base.Parse()
	if err != nil {
		writeErr(w, err)
		return
	}
	req := CreateRequest{
		BaseRequest: base,
	}
	sendResp(w, h(req))
}

func (h SearchHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	base := &BaseRequest{HTTPRequest: r}
	err := base.Parse()
	if err != nil {
		writeErr(w, err)
		return
	}
	req := SearchRequest{
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
