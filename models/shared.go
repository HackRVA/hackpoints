package models

// swagger:response endpointSuccessResponse
type endpointSuccessResponse struct {
	Body EndpointSuccess
}

// EndpointSuccess -- success response
type EndpointSuccess struct {
	Ack bool `json:"ack"`
}
