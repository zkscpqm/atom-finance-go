package asset

import "github.com/zkscpqm/atom-finance-go/pkg/market"

func EquityAsset(ticker string, market market.Code) Asset {
	return New("ticker", ticker, Equity, market)
}

func AmericanEquityAsset(ticker string) Asset {
	return EquityAsset(ticker, market.USA)
}

func CanadianEquityAsset(ticker string) Asset {
	return EquityAsset(ticker, market.CAN)
}

func OTCEquityAsset(ticker string) Asset {
	return EquityAsset(ticker, market.OTC)
}

func BrazilianEquityAsset(ticker string) Asset {
	return EquityAsset(ticker, market.BRA)
}
