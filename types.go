package dexigo

import (
	"github.com/JustBugLord/dexigo/chains"
	"github.com/JustBugLord/dexigo/channels"
)

type Event string

const (
	Subscribe   Event = "subscribe"
	Unsubscribe Event = "unsubscribe"
	Pong        Event = "pong"
	Update      Event = "update"
)

type Argument struct {
	ChainId      chains.Chain     `json:"chainId"`
	Channel      channels.Channel `json:"channel"`
	TokenAddress string           `json:"tokenAddress"`
}

type WSRequest struct {
	Op   Event      `json:"op"`
	Args []Argument `json:"args"`
}

type WSResponse struct {
	Arg    Argument    `json:"arg"`
	Args   []Argument  `json:"args"`
	Data   []TokenData `json:"data"`
	Event  Event       `json:"event"`
	ConnId string      `json:"connId"`
}

type TokenData struct {
	BundleHoldingRatio    string     `json:"bundleHoldingRatio"`
	ChainId               string     `json:"chainId"`
	Change                string     `json:"change"`
	Change1H              string     `json:"change1H"`
	Change4H              string     `json:"change4H"`
	Change5M              string     `json:"change5M"`
	ChangeUtc0            string     `json:"changeUtc0"`
	ChangeUtc8            string     `json:"changeUtc8"`
	CirculatingSupply     string     `json:"circulatingSupply"`
	DevHoldingRatio       string     `json:"devHoldingRatio"`
	EarlyBuyerHoldAmount  string     `json:"earlyBuyerHoldAmount"`
	Fdv                   string     `json:"fdv"`
	FirstPriceTime        string     `json:"firstPriceTime"`
	Holders               string     `json:"holders"`
	LaunchedTokenCount    string     `json:"launchedTokenCount"`
	Liquidity             string     `json:"liquidity"`
	MarketCap             string     `json:"marketCap"`
	MaxPrice              string     `json:"maxPrice"`
	MaxSupply             string     `json:"maxSupply"`
	MinPrice              string     `json:"minPrice"`
	Price                 string     `json:"price"`
	Progress              string     `json:"progress"`
	RiskControlLevel      string     `json:"riskControlLevel"`
	RiskLevel             string     `json:"riskLevel"`
	SnipersClear          string     `json:"snipersClear"`
	SnipersTotal          string     `json:"snipersTotal"`
	SuspiciousRatio       string     `json:"suspiciousRatio"`
	TagList               [][]string `json:"tagList"`
	TokenContractAddress  string     `json:"tokenContractAddress"`
	TokenFee              string     `json:"tokenFee"`
	TotalEarlyBuyerAmount string     `json:"totalEarlyBuyerAmount"`
	TradeNum              string     `json:"tradeNum"`
	Txs                   string     `json:"txs"`
	Txs1H                 string     `json:"txs1H"`
	Txs4H                 string     `json:"txs4H"`
	Txs5M                 string     `json:"txs5M"`
	UniqueTraders         string     `json:"uniqueTraders"`
	UniqueTraders1H       string     `json:"uniqueTraders1H"`
	UniqueTraders4H       string     `json:"uniqueTraders4H"`
	UniqueTraders5M       string     `json:"uniqueTraders5M"`
	Volume                string     `json:"volume"`
	Volume1H              string     `json:"volume1H"`
	Volume4H              string     `json:"volume4H"`
	Volume5M              string     `json:"volume5M"`
}

type NetworkToken struct {
	Amount               string       `json:"amount"`
	AmountNum            string       `json:"amountNum"`
	AmountNumBigDecimal  int          `json:"amountNumBigDecimal"`
	BuyTaxes             string       `json:"buyTaxes"`
	ChainBWLogoUrl       string       `json:"chainBWLogoUrl"`
	ChainId              chains.Chain `json:"chainId"`
	ChainLogoUrl         string       `json:"chainLogoUrl"`
	ChainName            string       `json:"chainName"`
	Change               string       `json:"change"`
	CollectTime          string       `json:"collectTime"`
	CurrencyAmount       string       `json:"currencyAmount"`
	Decimals             int          `json:"decimals"`
	ExplorerUrl          string       `json:"explorerUrl"`
	FromCurrency         int          `json:"fromCurrency"`
	FromMarket           int          `json:"fromMarket"`
	IsAuth               int          `json:"isAuth"`
	IsCollectToken       int          `json:"isCollectToken"`
	IsCustomToken        int          `json:"isCustomToken"`
	IsDefault            int          `json:"isDefault"`
	IsHoneypot           int          `json:"isHoneypot"`
	IsLeveraged          int          `json:"isLeveraged"`
	IsNativeToken        int          `json:"isNativeToken"`
	IsSafeMoonToken      int          `json:"isSafeMoonToken"`
	IsSubscribe          int          `json:"isSubscribe"`
	Liquidity            string       `json:"liquidity"`
	MarketCap            string       `json:"marketCap"`
	Price                string       `json:"price"`
	RiskLevel            string       `json:"riskLevel"`
	SellTaxes            string       `json:"sellTaxes"`
	TagList              [][]string   `json:"tagList"`
	TokenContractAddress string       `json:"tokenContractAddress"`
	TokenLogoUrl         string       `json:"tokenLogoUrl"`
	TokenName            string       `json:"tokenName"`
	TokenSymbol          string       `json:"tokenSymbol"`
	TokenType            string       `json:"tokenType"`
	TopPlacement         int          `json:"topPlacement"`
	TvlUsd               string       `json:"tvlUsd"`
	Volume               string       `json:"volume"`
}

type TokenInfo struct {
	BundleHoldingRatio        string                    `json:"bundleHoldingRatio"`
	ChainBWLogoUrl            string                    `json:"chainBWLogoUrl"`
	ChainLogoUrl              string                    `json:"chainLogoUrl"`
	ChainName                 string                    `json:"chainName"`
	Change                    string                    `json:"change"`
	Change1H                  string                    `json:"change1H"`
	Change4H                  string                    `json:"change4H"`
	Change5M                  string                    `json:"change5M"`
	ChangeUtc0                string                    `json:"changeUtc0"`
	ChangeUtc8                string                    `json:"changeUtc8"`
	CirculatingSupply         string                    `json:"circulatingSupply"`
	DappList                  []string                  `json:"dappList"`
	DevHoldingRatio           string                    `json:"devHoldingRatio"`
	EarlyBuyerStatisticsInfo  TokenEarlyBuyerStatistics `json:"earlyBuyerStatisticsInfo"`
	Holders                   string                    `json:"holders"`
	IsCollected               string                    `json:"isCollected"`
	IsNotSupportTxNativeToken string                    `json:"isNotSupportTxNativeToken"`
	IsSubscribe               string                    `json:"isSubscribe"`
	IsSupportBlinksShareUrl   string                    `json:"isSupportBlinksShareUrl"`
	IsSupportHolder           string                    `json:"isSupportHolder"`
	IsSupportHolderExpandData string                    `json:"isSupportHolderExpandData"`
	IsSupportMarketCapKline   string                    `json:"isSupportMarketCapKline"`
	IsTxPrice                 string                    `json:"isTxPrice"`
	Liquidity                 string                    `json:"liquidity"`
	MarketCap                 string                    `json:"marketCap"`
	MaxPrice                  string                    `json:"maxPrice"`
	MinPrice                  string                    `json:"minPrice"`
	ModuleType                string                    `json:"moduleType"`
	NativeTokenSymbol         string                    `json:"nativeTokenSymbol"`
	Price                     string                    `json:"price"`
	RiskControlLevel          string                    `json:"riskControlLevel"`
	RiskLevel                 string                    `json:"riskLevel"`
	SnipersClear              string                    `json:"snipersClear"`
	SnipersTotal              string                    `json:"snipersTotal"`
	SupportLimitOrder         string                    `json:"supportLimitOrder"`
	SupportMemeMode           string                    `json:"supportMemeMode"`
	SupportSingleChainSwap    string                    `json:"supportSingleChainSwap"`
	SupportSwap               string                    `json:"supportSwap"`
	SupportTrader             string                    `json:"supportTrader"`
	SuspiciousHoldingRatio    string                    `json:"suspiciousHoldingRatio"`
	T                         []struct {
		E struct {
		} `json:"e"`
		K string `json:"k"`
		M int    `json:"m"`
	} `json:"t"`
	TagList                     [][]interface{}    `json:"tagList"`
	TokenContractAddress        string             `json:"tokenContractAddress"`
	TokenFee                    string             `json:"tokenFee"`
	TokenLargeLogoUrl           string             `json:"tokenLargeLogoUrl"`
	TokenLogoUrl                string             `json:"tokenLogoUrl"`
	TokenName                   string             `json:"tokenName"`
	TokenSymbol                 string             `json:"tokenSymbol"`
	TokenThirdPartInfo          TokenThirdPartInfo `json:"tokenThirdPartInfo"`
	Top10HoldAmountPercentage   string             `json:"top10HoldAmountPercentage"`
	TradeNum                    string             `json:"tradeNum"`
	TransactionNum              string             `json:"transactionNum"`
	Volume                      string             `json:"volume"`
	WrapperTokenContractAddress string             `json:"wrapperTokenContractAddress"`
}

type TokenEarlyBuyerStatistics struct {
	ChainId               int    `json:"chainId"`
	EarlyBuyerHoldAmount  string `json:"earlyBuyerHoldAmount"`
	TokenContractAddress  string `json:"tokenContractAddress"`
	TotalEarlyBuyerAmount string `json:"totalEarlyBuyerAmount"`
}

type TokenThirdPartInfo struct {
	OkxDarkDefaultLogo         string `json:"okxDarkDefaultLogo"`
	OkxDarkHoverLogo           string `json:"okxDarkHoverLogo"`
	OkxLightDefaultLogo        string `json:"okxLightDefaultLogo"`
	OkxLightHoverLogo          string `json:"okxLightHoverLogo"`
	OkxWebSiteName             string `json:"okxWebSiteName"`
	OkxWebSiteUrl              string `json:"okxWebSiteUrl"`
	ThirdPartyWebSiteColorLogo string `json:"thirdPartyWebSiteColorLogo"`
	ThirdPartyWebSiteGreyLogo  string `json:"thirdPartyWebSiteGreyLogo"`
	ThirdPartyWebSiteName      string `json:"thirdPartyWebSiteName"`
	ThirdPartyWebSiteUrl       string `json:"thirdPartyWebSiteUrl"`
}

type DexResponse struct {
	Code         int    `json:"code"`
	DetailMsg    string `json:"detailMsg"`
	ErrorCode    string `json:"error_code"`
	ErrorMessage string `json:"error_message"`
	Msg          string `json:"msg"`
}

type AllNetworkTokensResponse struct {
	DexResponse
	Data []NetworkToken `json:"data"`
}

type TokenInfoResponse struct {
	DexResponse
	Data *TokenInfo `json:"data"`
}

type SearchTokenData struct {
	InputContent string         `json:"inputContent"`
	SystemList   []NetworkToken `json:"systemList"`
}

type SearchTokenResponse struct {
	DexResponse
	Data *SearchTokenData `json:"data"`
}
