package types
 
type BlockchainDotcomResponse struct {
	USD struct {
		FifteenMinutePrice float64 `json:"15m"`
	} `json:"USD"`
} 
type CoinbaseResponse struct {
	Data struct {
		Amount string `json:"amount"`
	} `json:"data"`
} 
type CoingeckoResponse struct {
	Bitcoin struct {
		USD float64 `json:"usd"`
	} `json:"bitcoin"`
} 
type GeminiResponse struct {
	Last string `json:"last"`
} 
type KrakenResponse struct {
	Result struct {
		Ticker struct {
			LastPriceList []string `json:"c"`
		} `json:"XXBTZUSD"`
	} `json:"result"`
}