package types
 
type AdminCreateReferralHashRequest struct {
	// A username or public name can be provided. If both are provided, public key is used.
	UserPublicKeyBase58Check string `safeForLogging:"true"`
	Username                 string `safeForLogging:"true"`

	// ReferralInfo to add for the new referral hash.
	ReferrerAmountUSDCents uint64 `safeForLogging:"true"`
	RefereeAmountUSDCents  uint64 `safeForLogging:"true"`
	MaxReferrals           uint64 `safeForLogging:"true"`
	RequiresJumio          bool   `safeForLogging:"true"`

	AdminPublicKey string `safeForLogging:"true"`
} 
type AdminCreateReferralHashResponse struct {
	ReferralInfoResponse ReferralInfoResponse `safeForLogging:"true"`
} 
type AdminDownloadRefereeCSVResponse struct {
	CSVRows [][]string
} 
type AdminDownloadReferralCSVResponse struct {
	CSVRows [][]string
} 
type AdminGetAllReferralInfoForUserRequest struct {
	// A username or public name can be provided. If both are provided, public key is used.
	UserPublicKeyBase58Check string `safeForLogging:"true"`
	Username                 string `safeForLogging:"true"`

	AdminPublicKey string `safeForLogging:"true"`
} 
type AdminGetAllReferralInfoForUserResponse struct {
	ReferralInfoResponses []ReferralInfoResponse `safeForLogging:"true"`
} 
type AdminUpdateReferralHashRequest struct {
	// Referral hash to update.
	ReferralHashBase58 string `safeForLogging:"true"`

	// ReferralInfo to updatethe referral hash with.
	ReferrerAmountUSDCents uint64 `safeForLogging:"true"`
	RefereeAmountUSDCents  uint64 `safeForLogging:"true"`
	MaxReferrals           uint64 `safeForLogging:"true"`
	RequiresJumio          bool   `safeForLogging:"true"`
	IsActive               bool   `safeForLogging:"true"`

	AdminPublicKey string `safeForLogging:"true"`
} 
type AdminUpdateReferralHashResponse struct {
	ReferralInfoResponse ReferralInfoResponse `safeForLogging:"true"`
} 
type AdminUploadReferralCSVRequest struct {
	CSVRows [][]string
} 
type AdminUploadReferralCSVResponse struct {
	LinksCreated uint64
	LinksUpdated uint64
} 
type ReferralInfoResponse struct {
	IsActive      bool
	Info          ReferralInfo
	ReferredUsers []ProfileEntryResponse
} 
type SimpleReferralInfoResponse struct {
	IsActive bool
	Info     SimpleReferralInfo
}