package types
 
type GetBuyDeSoFeeBasisPointsResponse struct {
	BuyDeSoFeeBasisPoints uint64
} 
type GetUSDCentsToDeSoExchangeRateResponse struct {
	USDCentsPerDeSo uint64
} 
type SetBuyDeSoFeeBasisPointsRequest struct {
	BuyDeSoFeeBasisPoints uint64
	AdminPublicKey        string
} 
type SetBuyDeSoFeeBasisPointsResponse struct {
	BuyDeSoFeeBasisPoints uint64
} 
type SetUSDCentsToDeSoExchangeRateRequest struct {
	USDCentsPerDeSo uint64
	AdminPublicKey  string
} 
type SetUSDCentsToDeSoExchangeRateResponse struct {
	USDCentsPerDeSo uint64
}