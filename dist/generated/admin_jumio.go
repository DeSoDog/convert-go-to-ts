package types
 
type AdminJumioCallback struct {
	PublicKeyBase58Check string
	Username             string
	CountryAlpha3        string
} 
type AdminResetJumioRequest struct {
	PublicKeyBase58Check string
	Username             string
	JWT                  string
} 
type AdminSetJumioVerifiedRequest struct {
	PublicKeyBase58Check string
	Username             string
} 
type AdminUpdateJumioCountrySignUpBonusRequest struct {
	CountryCode             string
	CountryLevelSignUpBonus CountryLevelSignUpBonus
} 
type AdminUpdateJumioDeSoRequest struct {
	JWT       string
	DeSoNanos uint64
} 
type AdminUpdateJumioDeSoResponse struct {
	DeSoNanos uint64
} 
type AdminUpdateJumioKickbackUSDCentsRequest struct {
	JWT      string
	USDCents uint64
} 
type AdminUpdateJumioKickbackUSDCentsResponse struct {
	USDCents uint64
} 
type AdminUpdateJumioUSDCentsRequest struct {
	JWT      string
	USDCents uint64
} 
type AdminUpdateJumioUSDCentsResponse struct {
	USDCents uint64
} 
type CountrySignUpBonusResponse struct {
	CountryLevelSignUpBonus CountryLevelSignUpBonus
	CountryCodeDetails      Alpha3CountryCodeDetails
} 
type GetAllCountryLevelSignUpBonusResponse struct {
	SignUpBonusMetadata        map[string]CountrySignUpBonusResponse
	DefaultSignUpBonusMetadata CountryLevelSignUpBonus
}