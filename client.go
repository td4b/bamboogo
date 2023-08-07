package bamboogo

import (
	"encoding/base64"
	"net/http"
)

func NewClient(host string, company string, apikey string) (*Client, error) {
	client := &http.Client{}

	// Create basic authentication header
	authHeader := apikey + ":x"
	encodedAuth := base64.StdEncoding.EncodeToString([]byte(authHeader))

	headers := map[string]string{
		"Authorization": "Basic " + encodedAuth,
	}

	client.Transport = &transportWithHeaders{
		headers:   headers,
		transport: http.DefaultTransport,
	}
	// Create and send GET request
	c := Client{
		HostURL:    host,
		Company:    company,
		HTTPClient: client,
	}
	return &c, nil
}

type transportWithHeaders struct {
	headers   map[string]string
	transport http.RoundTripper
}

func (t *transportWithHeaders) RoundTrip(req *http.Request) (*http.Response, error) {
	for key, value := range t.headers {
		req.Header.Set(key, value)
	}
	return t.transport.RoundTrip(req)
}
