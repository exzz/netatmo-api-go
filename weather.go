package netatmo

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

// Read returns the list of stations owned by the user and their modules
func (c *Client) Read() (*DeviceCollection, error) {
	if c.httpClient == nil {
		return nil, ErrNotAuthenticated
	}

	data := url.Values{"app_type": {"app_station"}}

	req, err := http.NewRequest("GET", deviceURL, nil)
	if err != nil {
		return nil, err
	}
	req.URL.RawQuery = data.Encode()

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("bad HTTP return code: %d", resp.StatusCode)
	}

	result := &DeviceCollection{}
	if err = json.NewDecoder(resp.Body).Decode(result); err != nil {
		return nil, err
	}

	return result, nil
}
