package sentrywebhook

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	_ "github.com/joho/godotenv/autoload"
	log "github.com/sirupsen/logrus"
)

// HTTPClient wrap a http.Client to make API call
type HTTPClient struct {
	Client http.Client
}

// NewHTTPClient create new HTTPClient instance
func NewHTTPClient() HTTPClient {
	client := HTTPClient{}
	return client
}

// SendHTTPPOSTRequest send JSON POST request with body to url
func (client HTTPClient) SendHTTPPOSTRequest(url string, message []byte) ([]byte, error) {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(message))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	log.Debugf("Send JSON POST request to %v with data %v\n", url, string(message))
	res, err := client.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

// SendMessageTelegram send a text message to a Chat ID from envvar
func (client HTTPClient) SendMessageTelegram(message string) error {
	URL := fmt.Sprintf("https://api.telegram.org/bot%v/sendMessage", os.Getenv("API_TOKEN"))
	dataReq := fmt.Sprintf(`{"text": "%v", "chat_id": "%v"}`, "Test Message", os.Getenv("CHAT_ID"))

	res, err := client.SendHTTPPOSTRequest(URL, []byte(dataReq))
	if err != nil {
		return err
	}
	log.Debug("Response from Telegram", string(res))
	return nil
}
