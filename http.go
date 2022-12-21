package atom

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

func (c *Client) buildURL(parts ...any) (string, error) {
	strParts := []string{"api", fmt.Sprintf("%d.%d", c.apiVersionMajor, c.apiVersionMinor)}
	for _, p := range parts {
		strParts = append(strParts, fmt.Sprintf("%v", p))
	}
	u, err := url.JoinPath(c.baseUrl, strParts...)
	if err != nil {
		return "", fmt.Errorf("failed to build URL with parts %s: %v", strParts, err)
	}
	return fmt.Sprintf("%s?api_key=%s", u, c.apiKey), nil
}

func (c *Client) makeRequest(method string, url string, data map[string]interface{}, extraHeaders map[string]string) (*http.Request, error) {
	body, err := json.Marshal(data)
	_b, _ := json.MarshalIndent(data, "", "  ")
	fmt.Println(string(_b))
	if err != nil {
		return nil, fmt.Errorf("failed to marshal [%s %s] request body: %v", method, url, err)
	}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("failed to create [%s %s] request: %v", method, url, err)
	}
	if extraHeaders != nil {
		for k, v := range extraHeaders {
			req.Header.Set(k, v)
		}
	}
	if c.defaultHeaders != nil {
		for k, v := range c.defaultHeaders {
			req.Header.Set(k, v)
		}
	}

	return req, nil
}

func (c *Client) post(url string, data map[string]interface{}, extraHeaders map[string]string) (*http.Response, error) {

	req, err := c.makeRequest("POST", url, data, extraHeaders)
	if err != nil {
		return nil, fmt.Errorf("failed to build POST request to %s: %v", url, err)
	}
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to perform POST %s: %v", url, err)
	}
	return resp, nil
}
