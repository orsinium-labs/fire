package fire

import "net/http"

type Response struct {
	Status   int
	Resource any

	RequestID       string
	CorrelationID   string
	ETag            string
	LastModified    string
	Location        string
	ContentLocation string
}

func Ok(r any) Response {
	return Response{
		Status:   http.StatusOK,
		Resource: r,
	}
}
