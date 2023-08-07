package bamboogo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

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
