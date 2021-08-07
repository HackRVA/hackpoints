package score

import (
	"encoding/json"
	"errors"
	"hackpoints/datastore"
	"hackpoints/models"
	"net/http"
)

type ScoreServer struct {
	Store datastore.ScoreStore
}

var (
	ErrGettingScore = errors.New("could not get the score from datastore")
)

func (s *ScoreServer) Get(w http.ResponseWriter, r *http.Request) {
	score, err := s.Store.GetScore()
	if err != nil {
		http.Error(w, ErrGettingScore.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	scoreResponse := models.Score{Score: score}
	j, _ := json.Marshal(scoreResponse)
	w.Write(j)
}
