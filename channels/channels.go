package channels

type Channel string

const (
	Candle1S     Channel = "dex-token-candle1s"
	Candle30S    Channel = "dex-token-candle30s"
	Candle1M     Channel = "dex-token-candle1m"
	Candle5M     Channel = "dex-token-candle5m"
	Candle15M    Channel = "dex-token-candle15m"
	Candle30M    Channel = "dex-token-candle30m"
	Candle1H     Channel = "dex-token-candle1h"
	Candle4H     Channel = "dex-token-candle4h"
	CandleDayUtc Channel = "dex-token-candle1Dutc"
	DexMarket    Channel = "dex-market"
	DexMarketV3  Channel = "dex-market-v3"
)
