package fire

import (
	"encoding/json/v2"
	"net/http"

	"github.com/orsinium-labs/fhir-server/statuses"
)

type Response struct {
	Status   statuses.Status
	Resource any

	RequestID       string
	CorrelationID   string
	ETag            string
	LastModified    string
	Location        string
	ContentLocation string
	Date            string
}

func (r Response) write(w http.ResponseWriter) {
	r.writeHeaders(w)
	w.WriteHeader(int(r.Status))
	_ = json.MarshalWrite(w, r.Resource)
}

func (r Response) writeHeaders(w http.ResponseWriter) {
	h := w.Header()
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
