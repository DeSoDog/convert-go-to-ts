package types

import (
	"github.com/holiman/uint256"
)
 
type BlockPublicKeyRequest struct {
	PublicKeyBase58Check      string
	BlockPublicKeyBase58Check string
	Unblock                   bool
	JWT                       string
} 
type BlockPublicKeyResponse struct {
	BlockedPublicKeys map[string]struct{}
} 
type CoinEntryResponse struct {
	CreatorBasisPoints        uint64
	DeSoLockedNanos           uint64
	NumberOfHolders           uint64
	CoinsInCirculationNanos   uint64
	CoinWatermarkNanos        uint64

	// Deprecated: Temporary to add support for BitCloutLockedNanos
	BitCloutLockedNanos uint64 // Deprecated
} 
type DAOCoinEntryResponse struct {
	NumberOfHolders           uint64
	CoinsInCirculationNanos   uint256.Int
	MintingDisabled           bool
	TransferRestrictionStatus TransferRestrictionStatusString
} 
type DeletePIIRequest struct {
	PublicKeyBase58Check string
	JWT                  string
} 
type DiamondSenderSummaryResponse struct {
	SenderPublicKeyBase58Check   string
	ReceiverPublicKeyBase58Check string

	TotalDiamonds       uint64
	HighestDiamondLevel uint64

	DiamondLevelMap      map[uint64]uint64
	ProfileEntryResponse *ProfileEntryResponse
} 
type GetDiamondsForPublicKeyRequest struct {
	// The user we are getting diamonds for.
	PublicKeyBase58Check string `safeForLogging:"true"`

	// If true, fetch the diamonds this public key gave out instead of the diamond this public key received
	FetchYouDiamonded bool
} 
type GetDiamondsForPublicKeyResponse struct {
	DiamondSenderSummaryResponses []*DiamondSenderSummaryResponse
	TotalDiamonds                 uint64
} 
type GetFollowsResponse struct {
	PublicKeyToProfileEntry map[string]*ProfileEntryResponse `safeForLogging:"true"`
	NumFollowers            uint64
} 
type GetFollowsStatelessRequest struct {
	// Either PublicKeyBase58Check or Username can be set by the client to specify
	// which user we're obtaining follows for
	// If both are specified, PublicKeyBase58Check will supercede
	PublicKeyBase58Check        string `safeForLogging:"true"`
	Username                    string `safeForLogging:"true"`
	GetEntriesFollowingUsername bool   `safeForLogging:"true"`

	// Public Key of the last follower / followee from the previous page
	LastPublicKeyBase58Check string `safeForLogging:"true"`
	// Number of records to fetch
	NumToFetch uint64 `safeForLogging:"true"`
} 
type GetHodlersForPublicKeyRequest struct {
	// Either PublicKeyBase58Check or Username can be set by the client to specify
	// which user we're obtaining posts for
	// If both are specified, PublicKeyBase58Check will supercede
	PublicKeyBase58Check string `safeForLogging:"true"`
	Username             string `safeForLogging:"true"`

	// Public Key of the last post from the previous page
	LastPublicKeyBase58Check string `safeForLogging:"true"`
	// Number of records to fetch
	NumToFetch uint64 `safeForLogging:"true"`

	// If true, fetch DAO coin balance entries instead of creator coin balance entries
	IsDAOCoin bool `safeForLogging:"true"`

	// If true, fetch balance entries for your hodlings instead of balance entries for hodler's of your coin
	FetchHodlings bool

	// If true, fetch all hodlers/hodlings -- supercedes NumToFetch
	FetchAll bool
} 
type GetHodlersForPublicKeyResponse struct {
	Hodlers                  []*BalanceEntryResponse
	LastPublicKeyBase58Check string
} 
type GetNotificationsCountRequest struct {
	PublicKeyBase58Check string
} 
type GetNotificationsCountResponse struct {
	NotificationsCount          uint64
	LastUnreadNotificationIndex uint64
	// Whether new unread notifications were added and the user metadata should be updated
	UpdateMetadata bool
} 
type GetNotificationsRequest struct {
	// This is the index of the notification we want to start our paginated lookup at. We
	// will fetch up to "NumToFetch" notifications after it, ordered by index.  If no
	// index is provided we will return the most recent posts.
	PublicKeyBase58Check string
	FetchStartIndex      int64
	NumToFetch           int64
	// This defines notifications that should be filtered OUT of the response
	// If a field is missing from this struct, it should be included in the response
	// Accepted values are like, diamond, follow, transfer, nft, post
	FilteredOutNotificationCategories map[string]bool
} 
type GetNotificationsResponse struct {
	Notifications       []*TransactionMetadataResponse
	ProfilesByPublicKey map[string]*ProfileEntryResponse
	PostsByHash         map[string]*PostEntryResponse
	LastSeenIndex       int64
} 
type GetProfilesRequest struct {
	// When set, we return profiles starting at the given pubkey up to numEntriesToReturn.
	PublicKeyBase58Check string `safeForLogging:"true"`
	// When set, we return profiles starting at the given username up to numEntriesToReturn.
	Username string `safeForLogging:"true"`
	// When specified, we filter out all profiles that don't have this
	// string as a prefix on their username.
	UsernamePrefix string `safeForLogging:"true"`
	// When set, we filter out profiles that don't contain this string
	// in their Description field.
	Description string `safeForLogging:"true"`
	OrderBy     string `safeForLogging:"true"`
	NumToFetch  uint32 `safeForLogging:"true"`
	// Public key of the user viewing the profile (affects post entry reader state).
	ReaderPublicKeyBase58Check string `safeForLogging:"true"`
	// Moderation type (currently empty string or 'leaderboard'). Empty string is for default
	// moderation.  'Leaderboard' is a special subset of profiles only removed from the leaderboards.
	ModerationType string `safeForLogging:"true"`
	// If a single profile is requested, return a list of HODLers and amount they HODL.
	FetchUsersThatHODL bool `safeForLogging:"true"`

	// If set to true, then the posts in the response will contain a boolean about whether they're in the global feed
	AddGlobalFeedBool bool `safeForLogging:"true"`
} 
type GetProfilesResponse struct {
	ProfilesFound []*ProfileEntryResponse
	NextPublicKey *string
} 
type GetSingleProfileRequest struct {
	// When set, we return profiles starting at the given pubkey up to numEntriesToReturn.
	PublicKeyBase58Check string `safeForLogging:"true"`
	// When set, we return profiles starting at the given username up to numEntriesToReturn.
	Username string `safeForLogging:"true"`
	// When true, we don't log a 404 for missing profiles
	NoErrorOnMissing bool `safeForLogging:"true"`
} 
type GetSingleProfileResponse struct {
	Profile       *ProfileEntryResponse
	IsBlacklisted bool
	IsGraylisted  bool
} 
type GetUserDerivedKeysRequest struct {
	// Public key which derived keys we want to query.
	PublicKeyBase58Check string `safeForLogging:"true"`
} 
type GetUserDerivedKeysResponse struct {
	// DerivedKeys contains user's derived keys indexed by public keys in base58Check
	DerivedKeys map[string]*UserDerivedKey `safeForLogging:"true"`
} 
type GetUserGlobalMetadataRequest struct {
	// The public key of the user who is trying to update their metadata.
	UserPublicKeyBase58Check string `safeForLogging:"true"`

	// JWT token authenticates the user
	JWT string
} 
type GetUserGlobalMetadataResponse struct {
	Email       string
	PhoneNumber string
} 
type GetUserMetadataRequest struct {
	PublicKeyBase58Check string
} 
type GetUserMetadataResponse struct {
	HasPhoneNumber   bool
	CanCreateProfile bool
	BlockedPubKeys   map[string]struct{}
	HasEmail         bool
	EmailVerified    bool
	// JumioFinishedTime = Time user completed flow in Jumio
	JumioFinishedTime uint64
	// JumioVerified = user was verified from Jumio flow
	JumioVerified bool
	// JumioReturned = jumio webhook called
	JumioReturned bool
} 
type GetUsersResponse struct {
	UserList                 []*User
	DefaultFeeRateNanosPerKB uint64
	ParamUpdaters            map[string]bool
} 
type GetUsersStatelessRequest struct {
	PublicKeysBase58Check []string `safeForLogging:"true"`
	SkipForLeaderboard    bool     `safeForLogging:"true"`
	GetUnminedBalance     bool     `safeForLogging:"true"`
} 
type IsFolllowingPublicKeyResponse struct {
	IsFollowing bool
} 
type IsFollowingPublicKeyRequest struct {
	PublicKeyBase58Check            string
	IsFollowingPublicKeyBase58Check string
} 
type IsHodlingPublicKeyRequest struct {
	PublicKeyBase58Check          string
	IsHodlingPublicKeyBase58Check string
	IsDAOCoin                     bool
} 
type IsHodlingPublicKeyResponse struct {
	IsHodling    bool
	BalanceEntry *BalanceEntryResponse
} 
type ProfileEntryResponse struct {
	// PublicKey is the key used by the user to sign for things and generally
	// verify her identity.
	PublicKeyBase58Check string
	Username             string
	Description          string
	IsHidden             bool
	IsReserved           bool
	IsVerified           bool
	Comments             []*PostEntryResponse
	Posts                []*PostEntryResponse
	// Creator coin fields
	CoinEntry *CoinEntryResponse

	// DAO Coin fields
	DAOCoinEntry *DAOCoinEntryResponse

	// Include current price for the frontend to display.
	CoinPriceDeSoNanos     uint64
	CoinPriceBitCloutNanos uint64 // Deprecated

	// Profiles of users that hold the coin + their balances.
	UsersThatHODL []*BalanceEntryResponse

	// If user is featured as a well known creator in the tutorial.
	IsFeaturedTutorialWellKnownCreator bool
	// If user is featured as an up and coming creator in the tutorial.
	// Note: a user should not be both featured as well known and up and coming
	IsFeaturedTutorialUpAndComingCreator bool
} 
type SetNotificationMetadataRequest struct {
	PublicKeyBase58Check string
	// The last notification index the user has seen
	LastSeenIndex int64
	// The last notification index that has been scanned
	LastUnreadNotificationIndex int64
	// The total count of unread notifications
	UnreadNotifications int64
	// JWT token
	JWT string
} 
type TransactionMetadataResponse struct {
	Metadata           TransactionMetadata
	TxnOutputResponses []*OutputResponse
	Txn                *TransactionResponse
	Index              int64
} 
type UpdateUserGlobalMetadataRequest struct {
	// The public key of the user who is trying to update their metadata.
	UserPublicKeyBase58Check string `safeForLogging:"true"`

	// JWT token authenticates the user
	JWT string

	// User's email for receiving notifications.
	Email string

	// A map of ContactPublicKeyBase58Check keys and number of read messages int values.
	MessageReadStateUpdatesByContact map[string]int
} 
type UserDerivedKey struct {
	// This is the public key of the owner.
	OwnerPublicKeyBase58Check string `safeForLogging:"true"`

	// This is the derived public key.
	DerivedPublicKeyBase58Check string `safeForLogging:"true"`

	// This is the expiration date of the derived key.
	ExpirationBlock uint64 `safeForLogging:"true"`

	// This is the current state of the derived key.
	IsValid bool `safeForLogging:"true"`
}