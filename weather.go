package netatmo

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"golang.org/x/oauth2"
)

const (
	// DefaultBaseURL is netatmo api url
	baseURL = "https://api.netatmo.net/"
	// DefaultAuthURL is netatmo auth url
	authURL = baseURL + "oauth2/token"
	// DefaultDeviceURL is netatmo device url
	deviceURL = baseURL + "/api/getstationsdata"
)

// Config is used to specify credential to Netatmo API
type Config struct {
	// ClientID from netatmo app registration at http://dev.netatmo.com/dev/listapps
	ClientID string
	// ClientSecret Client app secret
	ClientSecret string
	// Username Your netatmo account username
	Username string
	// Password Your netatmo account password
	Password string
}

// Client use to make request to Netatmo API
type Client struct {
	oauth      *oauth2.Config
	httpClient *http.Client
}

// NewClient create a handle authentication to NetAtmo API
func NewClient(config Config) (*Client, error) {
	oauth := &oauth2.Config{
		ClientID:     config.ClientID,
		ClientSecret: config.ClientSecret,
		Scopes:       []string{"read_station"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  baseURL,
			TokenURL: authURL,
		},
	}

	token, err := oauth.PasswordCredentialsToken(oauth2.NoContext, config.Username, config.Password)
	if err != nil {
		return nil, err
	}

	return &Client{
		oauth:      oauth,
		httpClient: oauth.Client(oauth2.NoContext, token),
	}, nil
}

// Read returns the list of stations owned by the user and their modules
func (c *Client) Read() (*DeviceCollection, error) {
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
