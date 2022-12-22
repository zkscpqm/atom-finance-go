package asset

import (
	"encoding/json"
	"github.com/zkscpqm/atom-finance-go/pkg/market"
)

type Type string

func (t Type) String() string { return string(t) }

const (
	Equity Type = "equity"
	Fund   Type = "fund"
	Crypto Type = "crypto"
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

func New(identifier, value string, assetType Type, market market.Code) Asset {
	return Asset{
		Identifier: identifier,
		Value:      value,
		AssetType:  assetType.String(),
		Market:     market.String(),
	}
}

func (a Asset) render() string {
	b, _ := json.Marshal(a)
	return string(b)
}
