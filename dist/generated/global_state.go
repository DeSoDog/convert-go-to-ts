package types

import (
	"github.com/dgraph-io/badger/v3"
)
  
type BatchGetRemoteRequest struct {
	KeyList [][]byte
} 
type BatchGetRemoteResponse struct {
	ValueList [][]byte
} 
type CountryLevelSignUpBonus struct {
	// If true, referee amount specified in referral code will be paid to users who sign up with IDs from this country.
	// If false, ReferralAmountOverrideUSDCents will be paid to users who sign up with IDs from this country.
	AllowCustomReferralAmount bool
	// Amount all referees will be paid when signing up from this country if AllowCustomReferralAmount is false.
	ReferralAmountOverrideUSDCents uint64
	// If true, referrer amount specified in referral code will be paid as a kickback to users who gave out referral
	// code that a user signed up with IDs from this country.
	// If false, KickbackAmountOverrideUSDCents will be paid as a kickback to referrers when a user signs up with an ID
	// from this country.
	AllowCustomKickbackAmount bool
	// Amount all referrers will be paid when a referee signs up from this country if AllowCustomKickbackAmount is
	// false.
	KickbackAmountOverrideUSDCents uint64
} 
type DeleteRemoteRequest struct {
	Key []byte
} 
type DeleteRemoteResponse struct {
} 
type GetRemoteRequest struct {
	Key []byte
} 
type GetRemoteResponse struct {
	Value []byte
} 
type GlobalState struct {
	GlobalStateRemoteNode   string
	GlobalStateRemoteSecret string
	GlobalStateDB           *badger.DB
} 
type HotFeedApprovedPostOp struct {
	IsRemoval  bool
	Multiplier float64 // Negatives are ignored when updating the ApprovedPosts map.
} 
type HotFeedPKIDMultiplierOp struct {
	InteractionMultiplier float64 // Negatives are ignored when updating the PKIDMultiplier map.
	PostsMultiplier       float64 // Negatives are ignored when updating the PKIDMultiplier map.
} 
type NFTDropEntry struct {
	IsActive        bool
	DropNumber      uint64
	DropTstampNanos uint64
	NFTHashes       []BlockHash
} 
type PhoneNumberMetadata struct {
	// The PublicKey of the user that this phone number belongs to.
	PublicKey []byte

	// E.164 format phone number for a user to receive text notifications at.
	PhoneNumber string

	// Country code associated with the user's phone number.
	PhoneNumberCountryCode string

	// if true, when the public key associated with this metadata tries to create a profile, we will comp their fee.
	ShouldCompProfileCreation bool

	// True if user deleted PII. Since users can
	PublicKeyDeleted bool
} 
type PutRemoteRequest struct {
	Key   []byte
	Value []byte
} 
type PutRemoteResponse struct {
} 
type ReferralInfo struct {
	ReferralHashBase58     string
	ReferrerPKID           PKID
	ReferrerAmountUSDCents uint64
	RefereeAmountUSDCents  uint64
	MaxReferrals           uint64 // If set to zero, there is no cap on referrals.
	RequiresJumio          bool

	// Stats
	NumJumioAttempts       uint64
	NumJumioSuccesses      uint64
	TotalReferrals         uint64
	TotalReferrerDeSoNanos uint64
	TotalRefereeDeSoNanos  uint64
	DateCreatedTStampNanos uint64
} 
type SeekRemoteRequest struct {
	StartPrefix    []byte
	ValidForPrefix []byte
	MaxKeyLen      int
	NumToFetch     int
	Reverse        bool
	FetchValues    bool
} 
type SeekRemoteResponse struct {
	KeysFound [][]byte
	ValsFound [][]byte
} 
type SimpleReferralInfo struct {
	ReferralHashBase58    string
	RefereeAmountUSDCents uint64
	MaxReferrals          uint64 // If set to zero, there is no cap on referrals.
	TotalReferrals        uint64
} 
type TutorialStatus string 
type UserMetadata struct {
	// The PublicKey of the user this metadata is associated with.
	PublicKey []byte

	// True if this user should be hidden from all data returned to the app.
	RemoveEverywhere bool

	// True if this user should be hidden from the creator leaderboard.
	RemoveFromLeaderboard bool

	// Email address for a user to receive email notifications at.
	Email string

	// Has the email been verified
	EmailVerified bool

	// E.164 format phone number for a user to receive text notifications at.
	PhoneNumber string

	// Country code associated with the user's phone number. This is a string like "US"
	PhoneNumberCountryCode string

	// This map stores the number of messages that a user has read from a specific contact.
	// The map is indexed with the contact's PublicKeyBase58Check and maps to an integer
	// number of messages that the user has read.
	MessageReadStateByContact map[string]int

	// Store the index of the last notification that the user saw
	NotificationLastSeenIndex int64

	// Amount of Bitcoin that users have burned so far via the Buy DeSo UI
	//
	// We track this so that, if the user does multiple burns,
	// we can set HasBurnedEnoughSatoshisToCreateProfile based on the total
	//
	// This tracks the "total input satoshis" (i.e. it includes fees the user spends).
	// Including fees makes it less expensive for a user to make a profile. We're cutting
	// users a break, but we could change this later.
	SatoshisBurnedSoFar uint64

	// True if the user has burned enough satoshis to create a profile. This can be
	// set to true from the BurnBitcoinStateless endpoint or canUserCreateProfile.
	//
	// We store this (instead of computing it when the user loads the page) to avoid issues
	// where the user burns the required amount, and then we reboot the node and change the min
	// satoshis required, and then the user hasn't burned enough. Once a user has burned enough,
	// we want him to be allowed to create a profile forever.
	HasBurnedEnoughSatoshisToCreateProfile bool

	// Map of public keys of profiles this user has blocked.  The map here functions as a hashset to make look ups more
	// efficient.  Values are empty structs to keep memory usage down.
	BlockedPublicKeys map[string]struct{}

	// If true, this user's posts will automatically be added to the global whitelist (max 5 per day).
	WhitelistPosts bool

	// JumioInternalReference = internal tracking reference for user's experience in Jumio
	JumioInternalReference string
	// JumioFinishedTime = has user completed flow in Jumio
	JumioFinishedTime uint64
	// JumioVerified = user was verified from Jumio flow
	JumioVerified bool
	// JumioReturned = jumio webhook called
	JumioReturned bool
	// JumioTransactionID = jumio's tracking number for the transaction in which this user was verified.
	JumioTransactionID string
	// JumioDocumentKey = Country - Document Type - Document SubType - Document Number. Helps uniquely identify users
	// and allows us to reset Jumio for a given user.
	// DEPRECATED
	JumioDocumentKey []byte
	// RedoJumio = boolean which allows user to skip the duplicate ID check in JumioCallback
	RedoJumio bool
	// JumioStarterDeSoTxnHashHex = Txn hash hex of the transaction in which the user was paid for
	// going through the Jumio flow
	JumioStarterDeSoTxnHashHex string
	// JumioShouldCompProfileCreation = True if we should comp the create profile fee because the user went through the
	// Jumio flow.
	JumioShouldCompProfileCreation bool

	// User must complete tutorial if they have been jumio verified.
	MustCompleteTutorial bool

	// If user is featured as a well known creator in the tutorial.
	IsFeaturedTutorialWellKnownCreator bool
	// If user is featured as an up and coming creator in the tutorial.
	// Note: a user should not be both featured as well known and up and coming
	IsFeaturedTutorialUpAndComingCreator bool

	TutorialStatus                  TutorialStatus
	CreatorPurchasedInTutorialPKID  PKID
	CreatorCoinsPurchasedInTutorial uint64

	// ReferralHashBase58Check with which user signed up
	ReferralHashBase58Check string

	// Txn hash in which the referrer was paid
	ReferrerDeSoTxnHash string

	// The number of unread notifications stored in the db.
	UnreadNotifications uint64
	// The most recently scanned notification transaction index in the database. Stored in order to prevent unnecessary re-scanning.
	LatestUnreadNotificationIndex int64
} 
type WyreWalletOrderMetadata struct {
	// Last payload received from Wyre webhook
	LatestWyreWalletOrderWebhookPayload WyreWalletOrderWebhookPayload

	// Track Wallet Order response received based on the last payload received from Wyre Webhook
	LatestWyreTrackWalletOrderResponse *WyreTrackOrderResponse

	// Amount of DeSo that was sent for this WyreWalletOrder
	DeSoPurchasedNanos uint64

	// BlockHash of the transaction for sending the DeSo
	BasicTransferTxnBlockHash BlockHash
}