package metanetworks

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"os"
	"os/user"
	"path/filepath"
	"time"
)

const (
	baseURL            string = "https://api.nsof.io"
	oauthURL           string = "/v1/oauth/token"
	maxIdleConnections int    = 10
	requestTimeout     int    = 60
	configPath         string = ".metanetworks/credentials.json"
)

// Config ...
type Config struct {
	APIKey    string `json:"api_key"`
	APISecret string `json:"api_secret"`
	Org       string `json:"org"`
}

// Client ...
type Client struct {
	APIKey           string
	APISecret        string
	Org              string
	TokenRefreshed   int64
	OAUTHToken       *Token
	HTTPClient       *http.Client
	terraformVersion string
}

// Token ...
type Token struct {
	Token         string `json:"access_token"`
	Expiry        int64  `json:"expires_in"`
	RefreshToken  string `json:"refresh_token"`
	RefreshExpiry int    `json:"refresh_expires_in"`
	TokenType     string `json:"token_type"`
}

// Credentials ...
type Credentials struct {
	GrantType    string `json:"grant_type"`
	Scope        string `json:"scope"`
	ClientID     string `json:"client_id,omitempty"`
	ClientSecret string `json:"client_secret,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty"`
}

// NewClientFromConfig Returns a Client from credentials found in a config file
func NewClientFromConfig() (*Client, error) {
	usr, _ := user.Current()
	dir := usr.HomeDir
	path := filepath.Join(dir, configPath)
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	configBytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	var config Config
	err = json.Unmarshal(configBytes, &config)
	if err != nil {
		return nil, errors.New("Could not parse credentials file, needs to contain one json object with keys: api_key, api_secret and org. " + err.Error())
	}

	return NewClient(config.APIKey, config.APISecret, config.Org)
}

// NewClient Returns a Client from credentials passed as parameters
func NewClient(key, secret, org string) (*Client, error) {
	httpClient := &http.Client{
		Transport: &http.Transport{
			MaxIdleConnsPerHost: maxIdleConnections,
		},
		Timeout: time.Duration(requestTimeout) * time.Second,
	}

	credentialData := Credentials{
		GrantType:    "client_credentials",
		Scope:        "org:" + org,
		ClientID:     key,
		ClientSecret: secret,
	}

	token, err := MakeAuthReqest(&credentialData, httpClient)
	if err != nil {
		return nil, err
	}

	now := time.Now()

	return &Client{
		APIKey:         key,
		APISecret:      secret,
		Org:            org,
		TokenRefreshed: now.Unix(),
		OAUTHToken:     token,
		HTTPClient:     httpClient,
	}, nil
}

// MakeAuthReqest ...
func MakeAuthReqest(credentials *Credentials, client *http.Client) (*Token, error) {
	jsonData, err := json.Marshal(credentials)
	if err != nil {
		return nil, err
	}

	resp, err := client.Post(baseURL+oauthURL, "application/json", bytes.NewReader(jsonData))
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

// RefreshToken ..
func (c *Client) RefreshToken() error {
	credentialData := Credentials{
		GrantType:    "refresh_token",
		Scope:        "org:" + c.Org,
		RefreshToken: c.OAUTHToken.RefreshToken,
	}

	token, err := MakeAuthReqest(&credentialData, c.HTTPClient)
	if err != nil {
		return err
	}

	c.OAUTHToken = token

	return nil
}
