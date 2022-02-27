package types

import (
	"github.com/holiman/uint256"
)
 
type BalanceEntryResponse struct {
	// The public keys are provided for the frontend
	HODLerPublicKeyBase58Check string
	// The public keys are provided for the frontend
	CreatorPublicKeyBase58Check string

	// Has the hodler purchased this creator's coin
	HasPurchased bool

	// How much this HODLer owns of a particular creator coin.
	BalanceNanos uint64

	// For simplicity, we create a new field for the uint256 balance for DAO coins
	BalanceNanosUint256 uint256.Int

	// The net effect of transactions in the mempool on a given BalanceEntry's BalanceNanos.
	// This is used by the frontend to convey info about mining.
	NetBalanceInMempool int64

	ProfileEntryResponse *ProfileEntryResponse `json:",omitempty"`
} 
type MessageContactResponse struct {
	// PublicKeyBase58Check is the public key in base58check format of the message contact.
	PublicKeyBase58Check string

	// Messages is the list of messages within this contact.
	Messages             []*MessageEntryResponse

	// ProfileEntryResponse is the profile entry corresponding to the contact.
	ProfileEntryResponse *ProfileEntryResponse

	// The number of messages this user has read from this contact. This is
	// used to show a notification badge for unread messages.
	NumMessagesRead int64
} 
type MessageEntryResponse struct {
	// SenderPublicKeyBase58Check is the main public key of the sender in base58check.
	SenderPublicKeyBase58Check    string

	// RecipientPublicKeyBase58Check is the main public key of the recipient in base58check.
	RecipientPublicKeyBase58Check string

	// EncryptedText is the encrypted message in hex format.
	EncryptedText string
	// TstampNanos is the message's timestamp.
	TstampNanos   uint64

	// Whether or not the user is the sender of the message.
	IsSender bool

	// Indicate if message was encrypted using shared secret
	V2 bool // Deprecated

	// Indicate message version
	Version uint32

	// ---------------------------------------------------------
	// DeSo V3 Messages Fields
	// ---------------------------------------------------------

	// SenderMessagingPublicKey is the sender's messaging public key that was used
	// to encrypt the corresponding message.
	SenderMessagingPublicKey string

	// SenderMessagingGroupKeyName is the sender's group key name of SenderMessagingPublicKey
	SenderMessagingGroupKeyName string

	// RecipientMessagingPublicKey is the recipient's messaging public key that was
	// used to encrypt the corresponding message.
	RecipientMessagingPublicKey string

	// RecipientMessagingGroupKeyName is the recipient's group key name of RecipientMessagingPublicKey
	RecipientMessagingGroupKeyName string
} 
type MessagingGroupEntryResponse struct {
	// GroupOwnerPublicKeyBase58Check is the main public key of the group owner, or, equivalently, the public key that
	// registered the group.
	GroupOwnerPublicKeyBase58Check string

	// MessagingPublicKeyBase58Check is the group messaging public key in base58check.
	MessagingPublicKeyBase58Check string

	// MessagingGroupKeyName is the name of the group messaging key.
	MessagingGroupKeyName string

	// MessagingGroupMembers is the list of the members in the group chat.
	MessagingGroupMembers []*MessagingGroupMemberResponse

	// EncryptedKey is the hex string of the encrypted private corresponding with the MessagingPublicKeyBase58Check.
	EncryptedKey string
} 
type MessagingGroupMemberResponse struct {
	// GroupMemberPublicKeyBase58Check is the main public key of the group member.
	GroupMemberPublicKeyBase58Check string

	// GroupMemberKeyName is the key name of the member that we encrypt the group messaging public key to. The group
	// messaging public key should not be confused with the GroupMemberPublicKeyBase58Check, the former is the public
	// key of the whole group, while the latter is the public key of the group member.
	GroupMemberKeyName string

	// EncryptedKey is the encrypted private key corresponding to the group messaging public key that's encrypted
	// to the member's registered messaging key labeled with GroupMemberKeyName.
	EncryptedKey string
} 
type TransactionInfo struct {
	TotalInputNanos          uint64
	SpendAmountNanos         uint64
	ChangeAmountNanos        uint64
	FeeNanos                 uint64
	TransactionIDBase58Check string

	// These are Base58Check encoded
	RecipientPublicKeys   []string
	RecipientAmountsNanos []uint64

	TransactionHex string

	// TODO: Not including the transaction because it causes encoding to
	// fail due to the presence of an interface for TxnMeta.
	//Transaction    MsgDeSoTxn

	// Unix timestamp (seconds since epoch).
	TimeAdded int64
} 
type User struct {
	// The public key for the user is computed from the seed using the exact
	// parameters used to generate the BTC deposit address below. Because
	// of this, the DeSo private and public key pair is also the key
	// pair corresponding to the BTC address above. We store this same
	// key in base58 format above for convenience in communicating with
	// the FE.
	PublicKeyBase58Check string

	ProfileEntryResponse *ProfileEntryResponse

	Utxos               []*UTXOEntryResponse
	BalanceNanos        uint64
	UnminedBalanceNanos uint64

	PublicKeysBase58CheckFollowedByUser []string

	UsersYouHODL         []*BalanceEntryResponse
	UsersWhoHODLYouCount int

	// HasPhoneNumber is a computed boolean so we can avoid returning the phone number in the
	// API response, since phone numbers are sensitive PII.
	HasPhoneNumber   bool
	CanCreateProfile bool
	BlockedPubKeys   map[string]struct{}
	HasEmail         bool
	EmailVerified    bool

	// JumioStartTime = Time user requested to initiate Jumio flow
	JumioStartTime uint64
	// JumioFinishedTime = Time user completed flow in Jumio
	JumioFinishedTime uint64
	// JumioVerified = user was verified from Jumio flow
	JumioVerified bool
	// JumioReturned = jumio webhook called
	JumioReturned bool

	// Is this user an admin
	IsAdmin bool
	// Is th user a super admin
	IsSuperAdmin bool

	// Is this user blacklisted/graylisted
	IsBlacklisted bool
	IsGraylisted  bool

	// Where is the user in the tutorial flow
	TutorialStatus TutorialStatus

	// Username of creator purchased during onboarding flow - used in case a user changes devices in the middle of the flow.
	CreatorPurchasedInTutorialUsername *string `json:",omitempty"`

	// Amount of creator coins purchased in the tutorial
	CreatorCoinsPurchasedInTutorial uint64

	// Does this user need to complete the tutorial
	MustCompleteTutorial bool
}