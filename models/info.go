package models

// swagger:response infoResponse
type infoResponse struct {
	Body InfoResponse
}

// InfoResponse -- response of info request
type InfoResponse struct {
	// Info Message
	//
	// Example: hello, world!
	Message string `json:"message"`
	Version int    `json:"version"`
	Commit  string `json:"commit"`
}
