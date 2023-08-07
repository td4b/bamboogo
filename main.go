package bamboogo

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Client
type Client struct {
	HostURL    string
	Company    string
	HTTPClient *http.Client
}

type User struct {
	ID         int    `json:"id"`
	EmployeeID int    `json:"employeeId"`
	FirstName  string `json:"firstName"`
	LastName   string `json:"lastName"`
	Email      string `json:"email"`
	Status     string `json:"status"`
	LastLogin  string `json:"lastLogin"`
}

type Users map[string]User

func (c *Client) Getusers() (Users, error) {
	resp, err := c.HTTPClient.Get(c.HostURL + "/" + c.Company + "/v1/meta/users/")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Read and print the response body
	if resp.StatusCode == http.StatusOK {
		// Read and unmarshal the response body
		var users Users
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(body, &users)
		if err != nil {
			return nil, err
		}
		return users, nil
	} else {
		return nil, fmt.Errorf("Invalid Response code from server: %d", resp.StatusCode)
	}
}

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
