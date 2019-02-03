package sentrywebhook

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIndexHandler(t *testing.T) {
	var server Server
	handler := http.HandlerFunc(server.IndexHandler)
	rr := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	handler.ServeHTTP(rr, req)

	assert.Equal(t, rr.Code, http.StatusOK)
	assert.Equal(t, rr.Body.String(), "Hello world")
}

func TestHookHandler(t *testing.T) {
	sampleResponseData, err := ioutil.ReadFile("sampleResponse.json")
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("POST", "/", bytes.NewBuffer(sampleResponseData))
	if err != nil {
		t.Fatal(err)
	}

	var server Server
	rr := httptest.NewRecorder()
	hanlder := http.HandlerFunc(server.HookHandler)
	hanlder.ServeHTTP(rr, req)

	fmt.Println(string(rr.Body.String()))
}
