package sentrywebhook

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSendHTTPPOSTRequestReal(t *testing.T) {
	t.Log("Starting real API call")
	client := NewHTTPClient()
	res, err := client.SendHTTPPOSTRequest("https://httpbin.org/post", []byte(`{"hello": "world"}`))
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Response\n%v", string(res))
}

func TestSendHTTPPOSTRequestStub(t *testing.T) {
	t.Log("Starting stub API call")
	stubHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	client, clientClose := NewTestHTTPClient(stubHandler)
	defer clientClose()

	res, err := client.SendHTTPPOSTRequest("https://httpbin.org/post", []byte(`{"hello": "world"}`))
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("Reponse\n%v", string(res))
	assert.Equal(t, res, []byte("OK"))
}

func TestSendMessageTelegram(t *testing.T) {
	t.Log("Starting test SendMessageTelegram")
	stubHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	client, clientClose := NewTestHTTPClient(stubHandler)
	defer clientClose()

	err := client.SendMessageTelegram("Hello world")
	if err != nil {
		t.Fatal(err)
	}
}
