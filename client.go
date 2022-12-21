package atom

import (
	"encoding/json"
	"fmt"
	"github.com/zkscpqm/atom-finance-go/market"
	"net/http"
	"os"
	"time"
)

type ClientConfig struct {
	APIKey string `json:"api_key"` // Your Developer API key to access Atom Finance data. See: https://portal.atom.finance/

	APIVersion struct {
		Major int `json:"major"`
		Minor int `json:"minor"`
	} `json:"api_version,omitempty"` // Version of API to query. Usually 2.0 (that's the default anyway)

	RequestTimeout int `json:"request_timeout,omitempty"` // How long to wait for the request to complete in seconds
}

func NewConfig(path string) (*ClientConfig, error) {
	/*

		You can load a config straight from a JSON file which contains the fields specified in `ClientConfig`
		An example file 'config.json' would look like:

		{
			"api_key": "YOUR_API_KEY_HERE",
			"request_timeout": 5
		}

		You can then make a config from this file by:

			cfg, err := NewConfig("config.json")

	*/

	b, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file at %s: %v", path, err)
	}

	cfg := ClientConfig{
		RequestTimeout: 3,
	}
	err = json.Unmarshal(b, &cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to parse config file %s: %v", path, err)
	}
	return &cfg, nil
}

type Client struct {
	client         *http.Client
	baseUrl        string
	defaultHeaders map[string]string

	apiVersionMajor int
	apiVersionMinor int

	apiKey string
}

func NewClient(cfg *ClientConfig) (*Client, error) {
	c := &Client{
		client: &http.Client{
			Transport: http.DefaultTransport,
			Timeout:   time.Duration(cfg.RequestTimeout) * time.Second,
		},
		baseUrl: "https://platform.atom.finance",
		defaultHeaders: map[string]string{
			"Accept":       "application/json",
			"Content-Type": "application/json",
		},

		apiVersionMajor: cfg.APIVersion.Major,
		apiVersionMinor: cfg.APIVersion.Minor,

		apiKey: cfg.APIKey,
	}
	err := c.verify()
	if err != nil {
		return nil, fmt.Errorf("failed to verify client connection: %v", err)
	}
	return c, nil
}

func (c *Client) verify() error {
	return nil
}

func (c *Client) Close() error {
	c.client.CloseIdleConnections()
	return nil
}

func (c *Client) DEBUGAnalystEstimates(ticker string, mkt market.Code) error {
	body := map[string]interface{}{
		"asset": NewDefaultAsset(ticker, mkt),
	}
	u, err := c.buildURL("equity", "estimates")
	if err != nil {
		return fmt.Errorf("failed build analyst estimates URL: %v", err)
	}
	resp, err := c.post(u, body, nil)
	if err != nil {
		return fmt.Errorf("failed to perform equity estimate: %v", err)
	}
	defer resp.Body.Close()

	tempMap := map[string]interface{}{}

	err = json.NewDecoder(resp.Body).Decode(&tempMap)
	if err != nil {
		return fmt.Errorf("failed to decode analyst estimates response: %v", err)
	}

	prettyResp, err := json.MarshalIndent(tempMap, "", "    ")
	if err != nil {
		return fmt.Errorf("failed to marshal analyst estimates response: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to perform equity estimate: [%d] %s", resp.StatusCode, string(prettyResp))
	}
	fmt.Println("RESP:\n", string(prettyResp))
	return nil
}
