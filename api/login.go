package api

import (
	"net/http"
	"text/template"

	log "github.com/sirupsen/logrus"
)

func serveLogin(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/html; charset=utf-8")
	t, err := template.ParseFiles("templates/login.html")
	if err != nil {
		log.Errorf("what happened: %s", err)
		return
	}
	err = t.Execute(w, nil)
	if err != nil {
		log.Errorf("what happened: %s", err)
		return
	}
}
