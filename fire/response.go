package fire

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
