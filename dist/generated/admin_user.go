package types
 
type AdminGetAllUserGlobalMetadataRequest struct {
	NumToFetch int `safeForLogging:"true"`
} 
type AdminGetAllUserGlobalMetadataResponse struct {
	// A mapping between the PublicKeyBase58Check string and the user's global metadata.
	PubKeyToUserGlobalMetadata map[string]*UserMetadata
	PubKeyToUsername           map[string]string
} 
type AdminGetUserAdminDataRequest struct {
	UserPublicKeyBase58Check string
} 
type AdminGetUserAdminDataResponse struct {
	// Profile Data
	Username string

	// Verifiers
	IsVerified                 bool
	LastVerifierPublicKey      string
	LastVerifyRemoverPublicKey string

	// White/Gray/Black list
	IsWhitelisted                 bool
	LastWhitelisterPublicKey      string
	LastWhitelistRemoverPublicKey string
	IsGraylisted                  bool
	LastGraylisterPublicKey       string
	LastGraylistRemoverPublicKey  string
	IsBlacklisted                 bool
	LastBlacklisterPublicKey      string
	LastBlacklistRemoverPublicKey string

	// Phone number verification
	PhoneNumber string
	Email       string
 
	// Referral Code
	ReferralHashBase58Check            string
	JumioStarterDeSoTxnHashBase58Check string
	ReferrerDeSoTxnHashBase58Check     string
} 
type AdminGetUserGlobalMetadataRequest struct {
	UserPublicKeyBase58Check string `safeForLogging:"true"`
} 
type AdminGetUserGlobalMetadataResponse struct {
	// the user's global metadata.
	UserMetadata UserMetadata

	// The User object
	UserProfileEntryResponse *ProfileEntryResponse
} 
type AdminGetUsernameVerificationAuditLogsRequest struct {
	Username string
} 
type AdminGetUsernameVerificationAuditLogsResponse struct {
	VerificationAuditLogs []VerificationUsernameAuditLogResponse
} 
type AdminGetVerifiedUsersResponse struct {
	VerifiedUsers []string
} 
type AdminGrantVerificationBadgeRequest struct {
	UsernameToVerify string `safeForLogging:"true"`
	AdminPublicKey   string
} 
type AdminGrantVerificationBadgeResponse struct {
	Message string
} 
type AdminRemoveVerificationBadgeRequest struct {
	UsernameForWhomToRemoveVerification string `safeForLogging:"true"`
	AdminPublicKey                      string
} 
type AdminRemoveVerificationBadgeResponse struct {
	Message string
} 
type AdminUpdateUserGlobalMetadataRequest struct {
	// The public key of the user to update. This will trump 'Username' if both are provided.
	UserPublicKeyBase58Check string `safeForLogging:"true"`
	// The username associated with the public key to update.
	Username string `safeForLogging:"true"`

	// Whether this is a blacklist update or not.
	IsBlacklistUpdate bool `safeForLogging:"true"`
	// Set to true if this user's content should not show up anywhere on the site.
	// Only set if IsBlacklistUpdate == true.
	RemoveEverywhere bool `safeForLogging:"true"`
	// Should be set to true if this user should not show up on the creator leaderboard.
	// Only set if IsBlacklistUpdate == true.
	RemoveFromLeaderboard bool `safeForLogging:"true"`

	// Whether this is a whitelist update or not.
	IsWhitelistUpdate bool `safeForLogging:"true"`
	// Set to true to automatically show this users posts in the global feed (max 5 per day).
	WhitelistPosts bool `safeForLogging:"true"`

	// Remove PhoneNumberMetadata to allow re-registration
	RemovePhoneNumberMetadata bool `safeForLogging:"true"`

	AdminPublicKey string
} 
type FilterAuditLog struct {
	// Time at which the filter status was granted or removed.
	TimestampNanos uint64
	// The filter type being updated
	Filter FilterType
	// Username and PKID of the admin who filtered the user.
	UpdaterUsername string
	UpdaterPKID     PKID
	// The user who was filtered or had their filter removed.
	UpdatedUsername string
	UpdatedPKID     PKID
	// Indicator of whether this request granted the filter status or removed it.
	IsRemoval bool
} 
type FilterType uint32
 
type VerificationUsernameAuditLog struct {
	// Time at which the verification was granted or removed.
	TimestampNanos uint64
	// Username and PKID of the admin who verified the user.
	VerifierUsername string
	VerifierPKID     PKID
	// The user who was verified or had their verification removed.
	VerifiedUsername string
	VerifiedPKID     PKID
	// Indicator of whether this request granted verification or removed verification.
	IsRemoval bool
} 
type VerificationUsernameAuditLogResponse struct {
	TimestampNanos               uint64
	VerifierUsername             string
	VerifierPublicKeyBase58Check string
	VerifiedUsername             string
	VerifiedPublicKeyBase58Check string
	IsRemoval                    bool
} 
type VerifiedUsernameToPKID struct {
	VerifiedUsernameToPKID map[string]PKID
}