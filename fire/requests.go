package fire

import (
	"errors"
	"net/http"
	"strings"
)

type BaseRequest struct {
	HTTPRequest *http.Request

	ContentType     ContentType
	FHIRVersion     FHIRVersion
	Accept          ContentType
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

func (r *BaseRequest) Parse() error {
	raw := r.HTTPRequest

	// Parse headers.
	ct, fv, _ := strings.Cut(raw.Header.Get("Content-Type"), "; ")
	r.ContentType = ContentType(ct)
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
		err := raw.ParseForm()
		if err != nil {
			return err
		}
		q = raw.Form
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

	return r.validate()
}

func (r *BaseRequest) validate() error {
	switch r.ContentType {
	case ContentTypeForm:
	case ContentTypeJSON:
	case ContentTypeTurtle:
		return errors.New("Turtle format is not supported")
	case ContentTypeXML:
	case "":
	default:
		return errors.New("unsupported Content-Type")
	}

	switch r.Accept {
	case ContentTypeJSON:
	case ContentTypeTurtle:
		return errors.New("Turtle format is not supported")
	case ContentTypeXML:
	case "":
	default:
		return errors.New("unsupported Accept")
	}

	if r.FHIRVersion != "" && r.FHIRVersion != FHIRVersionR5 {
		return errors.New("unsupported FHIR version")
	}

	return nil
}

type ReadRequest struct {
	*BaseRequest
	ID string
}

type VReadRequest struct {
	*BaseRequest
	ID  string
	VID string
}

type UpdateRequest struct {
	*BaseRequest
	ID string
}

type PatchRequest struct {
	*BaseRequest
	ID string
}

type DeleteRequest struct {
	*BaseRequest
	ID string
}

type CreateRequest struct {
	*BaseRequest
}

type SearchRequest struct {
	*BaseRequest
}
