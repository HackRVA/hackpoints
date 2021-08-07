package models

type ScoreStore interface {
	Get() (int, error)
}

// swagger:response scoreResponse
type scoreResponse struct {
	Body Score
}

type Score struct {
	Score int `json:"score"`
}
