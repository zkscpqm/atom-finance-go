package asset

import (
	"encoding/json"
	"fmt"
	"github.com/zkscpqm/atom-finance-go/pkg/market"
	"strconv"
)

type AnalystEstimateRequest struct {
	Ticker       string
	Market       market.Code
	ExtraHeaders map[string]string
}

type AnalystRecommendations struct {
	Buy        int     `json:"NumberOfBuy"`
	Hold       int     `json:"NumberOfHold"`
	Sell       int     `json:"NumberOfSell"`
	StrongBuy  int     `json:"NumberOfStrongBuy"`
	StrongSell int     `json:"NumberOfStrongSell"`
	Overall    float64 `json:"overall"`
}

func (r AnalystRecommendations) TotalBuy() int {
	return r.Buy + r.StrongBuy
}

func (r AnalystRecommendations) TotalSell() int {
	return r.Sell + r.StrongSell
}

type PriceTarget struct {
	CurrentTarget string `json:"currentTarget"`
	OldTarget     string `json:"oldTarget,omitempty"`
}

func (p PriceTarget) CurrentTargetFloat() float64 {
	f, err := strconv.ParseFloat(p.CurrentTarget, 64)
	if err != nil {
		return 0.
	}
	return f
}

func (p PriceTarget) OldTargetFloat() float64 {
	f, err := strconv.ParseFloat(p.OldTarget, 64)
	if err != nil {
		return 0.
	}
	return f
}

type BrokerageAction struct {
	Action   string `json:"action"`
	Date     string `json:"date"`
	Firm     string `json:"firm"`
	Position struct {
		Current string `json:"current,omitempty"`
		Old     string `json:"old,omitempty"`
	} `json:"position"`
	Price PriceTarget `json:"price"`
}

type FinancialData struct {
	BookValuePerShare         float64 `json:"bookValuePerShare,omitempty"`
	DividendPerShare          float64 `json:"dividendPerShare,omitempty"`
	Ebit                      float64 `json:"ebit,omitempty"`
	Ebitda                    float64 `json:"ebitda,omitempty"`
	Eps                       float64 `json:"eps,omitempty"`
	FreeCashFlowPerShare      float64 `json:"freeCashFlowPerShare,omitempty"`
	FinancialYear             int     `json:"fy"`
	GaExpense                 float64 `json:"gaExpense,omitempty"`
	LoanLossProvision         float64 `json:"loanLossProvision,omitempty"`
	NetIncome                 float64 `json:"netIncome,omitempty"`
	NetInterestIncome         float64 `json:"netInterestIncome,omitempty"`
	NetInvestmentIncome       float64 `json:"netInvestmentIncome,omitempty"`
	NetPerShare               float64 `json:"netPerShare,omitempty"`
	NonInterestExpense        float64 `json:"nonInterestExpense,omitempty"`
	NonPerformingAssets       float64 `json:"nonPerformingAssets,omitempty"`
	OperatingCashFlowPerShare float64 `json:"operatingCashFlowPerShare,omitempty"`
	Revenue                   float64 `json:"revenue,omitempty"`
	RiskWeightedAssets        float64 `json:"riskWeightedAssets,omitempty"`
	TangibleBookValuePerShare float64 `json:"tangibleBookValuePerShare,omitempty"`
	Tier1Capital              float64 `json:"tier1Capital,omitempty"`
	TotalDebt                 float64 `json:"totalDebt,omitempty"`
	TotalInterestExpense      float64 `json:"totalInterestExpense,omitempty"`
	TotalOperatingExpense     float64 `json:"totalOperatingExpense,omitempty"`
}

type Growth struct {
	EbitGrowth    float64 `json:"ebitGrowth"`
	EbitdaGrowth  float64 `json:"ebitdaGrowth"`
	EpsGrowth     float64 `json:"epsGrowth"`
	FinancialYear int     `json:"fy"`
	RevenueGrowth float64 `json:"revenueGrowth"`
}

type MarginAndRatioInfo struct {
	EbitMargin             float64 `json:"ebitMargin,omitempty"`
	EbitdaMargin           float64 `json:"ebitdaMargin,omitempty"`
	FinancialYear          int     `json:"fy"`
	GrossMargin            float64 `json:"grossMargin,omitempty"`
	IntExpMargin           float64 `json:"intExpMargin,omitempty"`
	NIEOverRevenue         float64 `json:"nIEOverRevenue,omitempty"`
	NIIOverRevenue         float64 `json:"nIIOverRevenue,omitempty"`
	NetIncomeMargin        float64 `json:"netIncomeMargin,omitempty"`
	OperatingExpenseMargin float64 `json:"operatingExpenseMargin,omitempty"`
	ReturnOnAssets         float64 `json:"returnOnAssets,omitempty"`
	ReturnOnEquity         float64 `json:"returnOnEquity,omitempty"`
	Tier1CapitalRatio      float64 `json:"tier1CapitalRatio,omitempty"`
}

type Valuation struct {
	EvEBIT                    float64 `json:"evEBIT,omitempty"`
	EvEBITDA                  float64 `json:"evEBITDA,omitempty"`
	EvRev                     float64 `json:"evRev,omitempty"`
	FinancialYear             int     `json:"fy"`
	PricePerBookValue         float64 `json:"pricePerBookValue,omitempty"`
	PricePerEarnings          float64 `json:"pricePerEarnings,omitempty"`
	PricePerNAV               float64 `json:"pricePerNAV,omitempty"`
	PricePerTangibleBookValue float64 `json:"pricePerTangibleBookValue,omitempty"`
}

type Guidance struct {
	AdjEpsMax           float64 `json:"adjEpsMax,omitempty"`
	AdjEpsMin           float64 `json:"adjEpsMin,omitempty"`
	GaapEpsMax          float64 `json:"gaapEpsMax,omitempty"`
	GaapEpsMin          float64 `json:"gaapEpsMin,omitempty"`
	AdjEpsPriorMax      float64 `json:"adjEpsPriorMax,omitempty"`
	AdjEpsPriorMin      float64 `json:"adjEpsPriorMin,omitempty"`
	GaapEpsPriorMax     float64 `json:"gaapEpsPriorMax,omitempty"`
	GaapEpsPriorMin     float64 `json:"gaapEpsPriorMin,omitempty"`
	AdjRevenueMax       float64 `json:"adjRevenueMax,omitempty"`
	AdjRevenueMin       float64 `json:"adjRevenueMin,omitempty"`
	GaapRevenueMax      int     `json:"gaapRevenueMax,omitempty"`
	GaapRevenueMin      int     `json:"gaapRevenueMin,omitempty"`
	AdjRevenuePriorMax  float64 `json:"adjRevenuePriorMax,omitempty"`
	AdjRevenuePriorMin  float64 `json:"adjRevenuePriorMin,omitempty"`
	GaapRevenuePriorMax float64 `json:"gaapRevenuePriorMax,omitempty"`
	GaapRevenuePriorMin float64 `json:"gaapRevenuePriorMin,omitempty"`
	Period              string  `json:"period,omitempty"`
	PeriodYear          int     `json:"periodYear,omitempty"`
}

type AnalystEstimateResponse struct {
	Asset     Asset `json:"asset"`
	Estimates struct {
		AnalystRecommendations AnalystRecommendations `json:"analystRecommendations"`
		BrokerageActions       []BrokerageAction      `json:"brokerageActions"`
		Consensus              struct {
			Financials struct {
				Annual    []FinancialData `json:"annual"`
				Quarterly []FinancialData `json:"quarterly"`
			} `json:"financials"`

			Growth struct {
				Annual []Growth `json:"annual"`
			} `json:"growth"`

			MarginsAndRatios struct {
				Annual    []MarginAndRatioInfo `json:"annual"`
				Quarterly []MarginAndRatioInfo `json:"quarterly"`
			} `json:"marginsAndRatios"`

			Valuation struct {
				Annual []Valuation `json:"annual"`
			} `json:"valuation"`
		} `json:"consensus"`

		Guidance struct {
			Annual    []Guidance `json:"annual"`
			Quarterly []Guidance `json:"quarterly"`
		} `json:"guidance"`
	} `json:"estimates"`
}

func (aer AnalystEstimateResponse) Json() (string, error) {
	prettyResp, err := json.MarshalIndent(&aer, "", "    ")
	if err != nil {
		return "", fmt.Errorf("failed to marshal Analyst Estimate Response: %v", err)
	}
	return string(prettyResp), nil
}
