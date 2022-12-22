package atom

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/zkscpqm/atom-finance-go/pkg/financial/asset"
	"net/http"
)

func (c *Client) AnalystEstimates(ctx context.Context, request asset.AnalystEstimateRequest) (rv asset.AnalystEstimateResponse, err error) {
	targetAsset := asset.EquityAsset(request.Ticker, request.Market)
	body := map[string]interface{}{
		"asset": targetAsset,
	}
	u, err := c.buildURL("equity", "estimates")
	if err != nil {
		return rv, fmt.Errorf("failed build analyst estimates URL: %v", err)
	}
	resp, err := c.post(ctx, u, body, request.ExtraHeaders)
	if err != nil {
		return rv, fmt.Errorf("failed to perform equity estimate: %v", err)
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&rv)
	if err != nil {
		return rv, fmt.Errorf("failed to decode analyst estimates response: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		return rv, fmt.Errorf("failed to perform equity estimate: [%d]", resp.StatusCode)
	}
	rv.Asset = targetAsset
	return
}
