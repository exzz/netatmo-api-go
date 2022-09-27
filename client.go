package netatmo

import (
	"context"
	"errors"
	"net/http"

	"golang.org/x/oauth2"
)

const (
	baseURL   = "https://api.netatmo.net/"
	authURL   = baseURL + "oauth2/authorize"
	tokenURL  = baseURL + "oauth2/token"
	deviceURL = baseURL + "/api/getstationsdata"
)

var (
	// ErrNotAuthenticated is returned from the client when it is not authenticated yet.
	ErrNotAuthenticated = errors.New("no token available")
)

// Config is used to specify credential to Netatmo API
type Config struct {
	// ClientID from netatmo app registration at http://dev.netatmo.com/dev/listapps
	ClientID string
	// ClientSecret Client app secret
	ClientSecret string
}

// Client use to make request to Netatmo API
type Client struct {
	oauth      *oauth2.Config
	httpClient *http.Client
}

// NewClient creates an unauthenticated NetAtmo API client.
func NewClient(config Config) *Client {
	oauth := &oauth2.Config{
		ClientID:     config.ClientID,
		ClientSecret: config.ClientSecret,
		Scopes:       []string{"read_station"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  authURL,
			TokenURL: tokenURL,
		},
	}

	return &Client{
		oauth: oauth,
	}
}

// AuthCodeURL creates an authentication URL that can be passed to the user.
func (c *Client) AuthCodeURL(redirectURL, state string) string {
	c.oauth.RedirectURL = redirectURL
	return c.oauth.AuthCodeURL(state)
}

// Exchange converts an authentication code into a token and authenticates the client.
func (c *Client) Exchange(ctx context.Context, code, state string) error {
	token, err := c.oauth.Exchange(ctx, code, oauth2.SetAuthURLParam("state", state))
	if err != nil {
		return err
	}

	c.httpClient = c.oauth.Client(ctx, token)
	return nil
}

// CurrentToken retrieves the token for persisting state.
func (c *Client) CurrentToken() (*oauth2.Token, error) {
	if c.httpClient == nil {
		return nil, ErrNotAuthenticated
	}

	transport := c.httpClient.Transport.(*oauth2.Transport)
	source := transport.Source
	return source.Token()
}

// InitWithToken initializes the client with an existing token.
func (c *Client) InitWithToken(ctx context.Context, token *oauth2.Token) {
	c.httpClient = c.oauth.Client(ctx, token)
}
