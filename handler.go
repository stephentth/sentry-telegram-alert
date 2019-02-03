package sentrywebhook

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/sirupsen/logrus"
)

// SentryDataSchema struct prepresent format of sentry sent to server
type SentryDataSchema struct {
	ID          string `json: id`
	Project     string `json: project`
	ProjectName string `json: project_name`
	Culprit     string `json: culprit`
	Level       string `json: level`
	URL         string `json: url`
	Message     string `json: message`
}

// Server prepresent server object
type Server struct {
	HTTPClient HTTPClient
	Log        logrus.Logger
}

// IndexHandler ...
func (s Server) IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world")
}

// HookHandler ...
func (s Server) HookHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" || r.URL.Host != "" {
		w.WriteHeader(400)
		w.Write([]byte("Bad Request"))
		return
	}

	jsnByte, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("Bad Request"))
		return
	}

	var sentryJson SentryDataSchema
	err = json.Unmarshal(jsnByte, &sentryJson)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("Bad Request"))
		return
	}

	fmt.Fprintf(w, "Got %v: %v", sentryJson.Message, sentryJson.URL)
}

// RunLocal run for developement testing
func (s Server) RunLocal() {
	http.HandleFunc("/hook", s.HookHandler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
