package types
 
type GetReferralInfoForReferralHashRequest struct {
	ReferralHash string
} 
type GetReferralInfoForReferralHashResponse struct {
	ReferralInfoResponse *SimpleReferralInfoResponse
	CountrySignUpBonus   CountryLevelSignUpBonus
} 
type GetReferralInfoForUserRequest struct {
	PublicKeyBase58Check string `safeForLogging:"true"`

	JWT string
} 
type GetReferralInfoForUserResponse struct {
	ReferralInfoResponses []ReferralInfoResponse `safeForLogging:"true"`
}