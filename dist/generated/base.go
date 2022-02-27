package types

import (
	"github.com/deso-protocol/core/lib"
)
  
type BlockchainDeSoTickerResponse struct {
	Symbol         string  `json:"symbol"`
	Price24H       float64 `json:"price_24h"`
	Volume24H      float64 `json:"volume_24h"`
	LastTradePrice float64 `json:"last_trade_price"`
} 
type CoinbaseDeSoTickerResponse struct {
	Data struct {
		Base     string `json:"base"`
		Currency string `json:"currency"`
		Amount   string `json:"amount"` // In USD
	} `json:"data"`
} 
type GetAppStateRequest struct {
	PublicKeyBase58Check string
} 
type GetAppStateResponse struct {
	MinSatoshisBurnedForProfileCreation uint64
	BlockHeight                         uint32
	IsTestnet                           bool

	HasStarterDeSoSeed    bool
	HasTwilioAPIKey       bool
	CreateProfileFeeNanos uint64
	CompProfileCreation   bool
	DiamondLevelMap       map[int64]uint64
	HasWyreIntegration    bool
	HasJumioIntegration   bool
	BuyWithETH            bool

	USDCentsPerDeSoExchangeRate uint64
	JumioDeSoNanos              uint64 // Deprecated
	JumioUSDCents               uint64
	JumioKickbackUSDCents       uint64
	// CountrySignUpBonus is the sign-up bonus configuration for the country inferred from a request's IP address.
	CountrySignUpBonus CountryLevelSignUpBonus

	DefaultFeeRateNanosPerKB uint64
	TransactionFeeMap        map[string][]TransactionFee

	// Address to which we want to send ETH when used to buy DESO
	BuyETHAddress string

	Nodes map[uint64]lib.DeSoNode

	USDCentsPerBitCloutExchangeRate uint64 // Deprecated
	JumioBitCloutNanos              uint64 // Deprecated
} 
type GetExchangeRateResponse struct {
	// BTC
	SatoshisPerDeSoExchangeRate    uint64
	USDCentsPerBitcoinExchangeRate uint64

	// ETH
	NanosPerETHExchangeRate    uint64
	USDCentsPerETHExchangeRate uint64

	// DESO
	NanosSold                          uint64
	USDCentsPerDeSoExchangeRate        uint64
	USDCentsPerDeSoReserveExchangeRate uint64
	BuyDeSoFeeBasisPoints              uint64

	SatoshisPerBitCloutExchangeRate        uint64 // Deprecated
	USDCentsPerBitCloutExchangeRate        uint64 // Deprecated
	USDCentsPerBitCloutReserveExchangeRate uint64 // Deprecated
}