package metanetworks

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/user"
	"path/filepath"
	"time"
)

const (
	BaseURL  string = "https://api.nsof.io"
	OAUTHURL string = "/v1/oauth/token"

	MaxIdleConnections int    = 10
	RequestTimeout     int    = 60
	ConfigFile         string = ".metanetworks/credentials.json"
)

type credentials struct {
	ApiKey    string `json:"api_key"`
	ApiSecret string `json:"api_secret"`
	Org       string `json:"org"`
}

type Client struct {
	ApiKey         string
	ApiSecret      string
	Org            string
	TokenRefreshed int64
	OauthToken     *Token
	HTTPClient     *http.Client
}

type Token struct {
	Token         string `json:"access_token"`
	Expiry        int64  `json:"expires_in"`
	RefreshToken  string `json:"refresh_token"`
	RefreshExpiry int    `json:"refresh_expires_in"`
	TokenType     string `json:"token_type"`
}

type CredentialData struct {
	GrantType    string `json:"grant_type"`
	Scope        string `json:"scope"`
	ClientId     string `json:"client_id,omitempty"`
	ClientSecret string `json:"client_secret,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty"`
}

// Returns a Client from credentials found in a config file
func NewClientFromConfig() (*Client, error) {
	usr, _ := user.Current()
	dir := usr.HomeDir
	path := filepath.Join(dir, ConfigFile)
	configFile, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	configBytes, err := ioutil.ReadAll(configFile)
	if err != nil {
		return nil, err
	}
	var config credentials
	err = json.Unmarshal(configBytes, &config)
	if err != nil {
		return nil, errors.New("Could not parse credentials file, needs to contain one json object with keys: api_key, api_secret and org. " + err.Error())
	}

	return NewClient(config.ApiKey, config.ApiSecret, config.Org)

}

// Returns a Client from credentials passed as parameters
func NewClient(key, secret, org string) (*Client, error) {

	http_client := &http.Client{
		Transport: &http.Transport{
			MaxIdleConnsPerHost: MaxIdleConnections,
		},
		Timeout: time.Duration(RequestTimeout) * time.Second,
	}

	credential_data := CredentialData{
		GrantType:    "client_credentials",
		Scope:        "org:" + org,
		ClientId:     key,
		ClientSecret: secret,
	}

	token, err := MakeAuthReqest(&credential_data, http_client)
	if err != nil {
		return nil, err
	}

	now := time.Now()

	return &Client{
		ApiKey:         key,
		ApiSecret:      secret,
		Org:            org,
		TokenRefreshed: now.Unix(),
		OauthToken:     token,
		HTTPClient:     http_client,
	}, nil
}

func MakeAuthReqest(credentials *CredentialData, client *http.Client) (*Token, error) {

	json_data, err := json.Marshal(credentials)
	if err != nil {
		return nil, err
	}

	resp, err := client.Post(BaseURL+OAUTHURL, "application/json", bytes.NewReader(json_data))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}
	var token Token
	err = json.Unmarshal(body, &token)
	if err != nil {
		return nil, err
	}
	return &token, nil

}

func (c *Client) RefreshToken() error {

	credential_data := CredentialData{
		GrantType:    "refresh_token",
		Scope:        "org:" + c.Org,
		RefreshToken: c.OauthToken.RefreshToken,
	}

	token, err := MakeAuthReqest(&credential_data, c.HTTPClient)
	if err != nil {
		return err
	}

	c.OauthToken = token

	return nil
}

func (c *Client) ReadRequest(endpoint string) ([]byte, error) {
	resp, err := c.MakeRequest(endpoint, "GET", nil, "application/json")
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *Client) CreateRequest(endpoint string, data []byte) ([]byte, error) {
	resp, err := c.MakeRequest(endpoint, "POST", data, "application/json")
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *Client) UpdateRequest(endpoint string, data []byte) ([]byte, error) {
	resp, err := c.MakeRequest(endpoint, "PATCH", data, "application/merge-patch+json")
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *Client) DeleteRequest(endpoint string) error {
	_, err := c.MakeRequest(endpoint, "DELETE", nil, "application/json")
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) MakeRequest(endpoint, method string, data []byte, content_type string) ([]byte, error) {

	if content_type == "" {
		content_type = "application/json"
	}
	// Do we need to refresh the token? Do this first because
	// the token might be expired, but the refresh token is ok.
	now := time.Now()
	if ((c.TokenRefreshed + c.OauthToken.Expiry) - now.Unix()) < 30 {
		err := c.RefreshToken()
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, BaseURL+endpoint, bytes.NewReader(data))

	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", content_type)
	req.Header.Add("Authorization", "Bearer "+c.OauthToken.Token)

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, errors.New(string(body))
	}

	log.Printf("Response: %s", body)
	return body, nil

}
