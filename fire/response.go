package fire

import (
	"encoding/json/jsontext"
	"encoding/json/v2"
	"encoding/xml"
	"net/http"

	"github.com/orsinium-labs/fhir-server/statuses"
)

type Response struct {
	Status   statuses.Status
	Resource any

	ContentType     ContentType
	RequestID       string
	CorrelationID   string
	ETag            string
	LastModified    string
	Location        string
	ContentLocation string
	Date            string
}

func (resp *Response) write[R Resource](req *BaseRequest[R], w http.ResponseWriter) {
	if resp.ContentType == "" {
		if req.Accept != "" {
			resp.ContentType = req.Accept
		} else {
			resp.ContentType = ContentTypeJSON
		}
	}
	resp.writeHeaders(w)
	w.WriteHeader(int(resp.Status))

	if resp.ContentType == ContentTypeJSON {
		if req.Pretty {
			_ = json.MarshalWrite(
				w, resp.Resource,
				jsontext.Multiline(true),
				jsontext.WithIndent("  "),
			)
		} else {
			_ = json.MarshalWrite(w, resp.Resource)
		}
	} else {
		e := xml.NewEncoder(w)
		if req.Pretty {
			e.Indent("", "  ")
		}
		_ = e.Encode(resp.Resource)
	}
}

func (r *Response) writeHeaders(w http.ResponseWriter) {
	h := w.Header()
	if r.ContentType != "" {
		h.Add("Content-Type", string(r.ContentType))
	}
	if r.RequestID != "" {
		h.Add("X-Request-Id", r.RequestID)
	}
	if r.CorrelationID != "" {
		h.Add("X-Correlation-Id", r.CorrelationID)
	}
	if r.ETag != "" {
		h.Add("ETag", r.ETag)
	}
	if r.LastModified != "" {
		h.Add("Last-Modified", r.LastModified)
	}
	if r.Location != "" {
		h.Add("Location", r.Location)
	}
	if r.ContentLocation != "" {
		h.Add("Content-Location", r.ContentLocation)
	}
	if r.Date != "" {
		h.Add("Date", r.Date)
	}
}
