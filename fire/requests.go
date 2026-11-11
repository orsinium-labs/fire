package fire

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/orsinium-labs/fhir-server/statuses"
)

type BaseRequest[R Resource] struct {
	HTTPRequest *http.Request

	ContentType ContentType
	// Content-negotiation for MIME Type.
	Accept ContentType
	// Content-negotiation for FHIR Version.
	FHIRVersion     FHIRVersion
	IfMatch         string
	IfModifiedSince string
	IfNoneExist     string
	IfNoneMatch     string
	Prefer          Prefer
	ClientTimezone  string
	RequestID       string
	CorrelationID   string

	Pretty   bool
	Elements []string
	Summary  Summary
}

func parseRequest[R Resource](raw *http.Request) *BaseRequest[R] {
	r := &BaseRequest[R]{HTTPRequest: raw}

	// Parse headers.
	r.ContentType = ContentType(raw.Header.Get("Content-Type"))
	ct, fv, _ := strings.Cut(raw.Header.Get("Accept"), "; ")
	r.Accept = ContentType(ct)
	if fv != "" {
		r.FHIRVersion = FHIRVersion(strings.TrimSuffix(fv, "fhirVersion="))
	}
	r.Accept = ContentType(raw.Header.Get("Accept"))

	r.IfMatch = raw.Header.Get("If-Match")
	r.IfModifiedSince = raw.Header.Get("If-Modified-Since")
	r.IfNoneExist = raw.Header.Get("If-None-Exist")
	r.IfNoneMatch = raw.Header.Get("If-None-Match")
	r.RequestID = raw.Header.Get("X-Request-Id")
	r.CorrelationID = raw.Header.Get("X-Correlation-Id")
	r.Prefer = Prefer(raw.Header.Get("Prefer"))
	r.ClientTimezone = raw.Header.Get("Client-Timezone")

	// Parse query params.
	q := raw.URL.Query()
	if r.ContentType == ContentTypeForm {
		_ = raw.ParseForm()
		if raw.Form != nil {
			q = raw.Form
		}
	}
	switch q.Get("_format") {
	case "xml", "text/xml", "application/xml", "application/fhir+xml":
		r.Accept = ContentTypeXML
	case "json", "application/json", "application/fhir+json":
		r.Accept = ContentTypeJSON
	case "ttl", "application/fhir+turtle", "text/turtle":
		r.Accept = ContentTypeTurtle
	}
	r.Pretty = q.Get("_pretty") == "true"
	r.Elements = strings.Split(q.Get("_elements"), ",")
	r.Summary = Summary(q.Get("_summary"))

	return r
}

func (r *BaseRequest[R]) BadRequest(res Resource) Response {
	return Response{Status: statuses.BadRequest, Resource: res}
}

type ReadRequest[R Resource] struct {
	*BaseRequest[R]
	ID string
}

func (r ReadRequest[R]) Ok(res R) Response {
	return Response{Status: statuses.OK, Resource: res}
}

type VReadRequest[R Resource] struct {
	*BaseRequest[R]
	ID  string
	VID string
}

func (r VReadRequest[R]) Ok(res R) Response {
	return Response{Status: statuses.OK, Resource: res}
}

type UpdateRequest[R Resource] struct {
	*BaseRequest[R]
	ID      string
	Body    R
	RawBody []byte
}

type PatchRequest[R Resource] struct {
	*BaseRequest[R]
	ID      string
	Body    R
	RawBody []byte
}

type DeleteRequest[R Resource] struct {
	*BaseRequest[R]
	ID string
}

type CreateRequest[R Resource] struct {
	*BaseRequest[R]
	Body    R
	RawBody []byte
}

type SearchRequest[R Resource] struct {
	*BaseRequest[R]
	Body url.Values
}

func (r SearchRequest[R]) Ok(res []R) Response {
	return Response{Status: statuses.OK, Resource: res}
}

type InstanceExtRequest[R Resource] struct {
	*BaseRequest[R]
	ID string
}

func (r InstanceExtRequest[R]) Ok(res Resource) Response {
	return Response{Status: statuses.OK, Resource: res}
}

type TypeExtRequest[R Resource] struct {
	*BaseRequest[R]
}

func (r TypeExtRequest[R]) Ok(res Resource) Response {
	return Response{Status: statuses.OK, Resource: res}
}
