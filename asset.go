package atom

import (
	"encoding/json"
	"github.com/zkscpqm/atom-finance-go/market"
)

type AssetType string

func (at AssetType) String() string { return string(at) }

const (
	AssetEquity AssetType = "equity"
	AssetFund   AssetType = "fund"
	AssetCrypto AssetType = "crypto"
)

type renderable interface {
	render() string
}

type Asset struct {
	Identifier string `json:"identifier"`
	Value      string `json:"value"`
	AssetType  string `json:"assetType"`
	Market     string `json:"market"`
}

func NewAsset(identifier, value string, assetType AssetType, market market.Code) Asset {
	return Asset{
		Identifier: identifier,
		Value:      value,
		AssetType:  assetType.String(),
		Market:     market.String(),
	}
}

func NewDefaultAsset(ticker string, market market.Code) Asset {
	return NewAsset("ticker", ticker, AssetEquity, market)
}

func (a Asset) render() string {
	b, _ := json.Marshal(a)
	return string(b)
}
