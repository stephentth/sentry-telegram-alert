package sentrywebhook

import (
	"context"
	"crypto/tls"
	"net"
	"net/http"
	"net/http/httptest"
)

// NewTestHTTPClient create new testing HTTPClient instance
func NewTestHTTPClient(handler http.Handler) (HTTPClient, func()) {
	client := HTTPClient{}
	server := httptest.NewTLSServer(handler)

	client.Client = http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			DialContext: func(_ context.Context, network, _ string) (net.Conn, error) {
				return net.Dial(network, server.Listener.Addr().String())
			},
		},
	}
	return client, server.Close
}
