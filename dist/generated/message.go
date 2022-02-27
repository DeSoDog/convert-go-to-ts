package types
 
type CheckPartyMessagingKeysRequest struct {
	// SenderPublicKeyBase58Check is the main public key of the sender in base58check format.
	SenderPublicKeyBase58Check    string

	// SenderMessagingKeyName is the sender's key name the existence of which we want to verify.
	SenderMessagingKeyName        string

	// RecipientPublicKeyBase58Check is the public key of the recipient in base58check format.
	RecipientPublicKeyBase58Check string

	// RecipientMessagingKeyName is the recipient's key name the existence of we want to verify.
	RecipientMessagingKeyName     string
} 
type CheckPartyMessagingKeysResponse struct {
	// SenderMessagingPublicKeyBase58Check is the group messaging public key of the sender corresponding to the provided
	// SenderMessagingKeyName. This field will be an empty string if the key name doesn't exist.
	SenderMessagingPublicKeyBase58Check    string

	// SenderMessagingKeyName is the key name that was passed in the initial request. It's added to the response for
	// convenience.
	SenderMessagingKeyName                 string

	// IsSenderMessagingKey determines if the SenderMessagingKeyName existed for the sender.
	IsSenderMessagingKey                   bool

	// RecipientMessagingPublicKeyBase58Check is the group messaging public key of the recipient corresponding to the provided
	// RecipientMessagingKeyName. This field will be an empty string if the key name doesn't exist.
	RecipientMessagingPublicKeyBase58Check string

	// RecipientMessagingKeyName is the key name that was passed in the initial request. It's added to the response for
	// convenience.
	RecipientMessagingKeyName              string

	// IsRecipientMessagingKey determines if the RecipientMessagingKeyName existed for the sender.
	IsRecipientMessagingKey                bool
} 
type GetAllMessagingGroupKeysRequest struct {
	// OwnerPublicKeyBase58Check is the public key in base58check of the account whose group messaging keys we want to fetch.
	OwnerPublicKeyBase58Check string
} 
type GetAllMessagingGroupKeysResponse struct {
	// MessagingGroupEntries is the list of all user's group messaging keys.
	MessagingGroupEntries []*MessagingGroupEntryResponse
} 
type GetMessagesResponse struct {
	// PublicKeyToProfileEntry is a map of profile entries of the message parties. Keys are base58check public keys.
	PublicKeyToProfileEntry     map[string]*ProfileEntryResponse

	// OrderedContactsWithMessages is a list of message contacts. Each entry in the list corresponds to a messaging
	// thread and contains the public key and profile entry of the other party in the thread. Entries also contain a
	// list of encrypted messages for the threads.
	OrderedContactsWithMessages []*MessageContactResponse

	// UnreadStateByContact is a map indexed by public key base58check of contacts and with boolean values corresponding
	// to whether the thread has any unread messages. True means there are unread messages.
	UnreadStateByContact        map[string]bool

	// NumberOfUnreadThreads is a counter of how many unread threads are there.
	NumberOfUnreadThreads       int

	// MessagingGroups are all user's registered messaging keys and group chats that the user is a member of.
	MessagingGroups               []*MessagingGroupEntryResponse
} 
type GetMessagesStatelessRequest struct {
	PublicKeyBase58Check string `safeForLogging:"true"`

	// FetchAfterPublicKeyBase58Check specifies where to start
	// in the messages to begin fetching new messages. If set empty,
	// we start fetching threads from the most recent message.
	FetchAfterPublicKeyBase58Check string `safeForLogging:"true"`

	// NumToFetch specifies the number of message threads to return. Defaults to 20
	// unless otherwise specified.
	NumToFetch uint64 `safeForLogging:"true"`

	// There are four filters: HoldersOnly, HoldingsOnly, FollowersOnly, FollowedOnly
	// If all filters are false, we return all messages. Otherwise we include
	// messages from the sets set true.

	// HoldersOnly when set true includes messages from holders.
	HoldersOnly bool `safeForLogging:"true"`

	// HoldingsOnly when set true includes messages from the user's holdings.
	HoldingsOnly bool `safeForLogging:"true"`

	// FollowersOnly when set true includes messages from the user's followers.
	FollowersOnly bool `safeForLogging:"true"`

	// FollowedOnly when set true includes messages from who the user follows.
	FollowingOnly bool `safeForLogging:"true"`

	// SortAlgorithm determines how the messages should be returned. Currently
	// it support time, deso, and followers based sorting.
	SortAlgorithm string `safeForLogging:"true"`
} 
type MarkAllMessagesReadRequest struct {
	JWT                      string
	UserPublicKeyBase58Check string
} 
type MarkContactMessagesReadRequest struct {
	JWT                         string
	UserPublicKeyBase58Check    string
	ContactPublicKeyBase58Check string
} 
type RegisterMessagingGroupKeyRequest struct {
	// OwnerPublicKeyBase58Check is the public key in base58check of the account we want to register the messaging key for.
	OwnerPublicKeyBase58Check     string

	// MessagingPublicKeyBase58Check is the public key in base58check of the messaging group we want to register.
	MessagingPublicKeyBase58Check string

	// MessagingGroupKeyName is the name of the group key.
	MessagingGroupKeyName         string

	// MessagingKeySignatureHex is the signature of sha256x2(MessagingPublicKey + MessagingGroupKeyName). Currently,
	// the signature is only needed to register the default key.
	MessagingKeySignatureHex      string

	MinFeeRateNanosPerKB          uint64 `safeForLogging:"true"`

	// No need to specify ProfileEntryResponse in each TransactionFee
	TransactionFees []TransactionFee `safeForLogging:"true"`
} 
type RegisterMessagingGroupKeyResponse struct {
	TotalInputNanos   uint64
	ChangeAmountNanos uint64
	FeeNanos          uint64
	Transaction       MsgDeSoTxn
	TransactionHex    string
	TxnHashHex        string
} 
type SendMessageStatelessRequest struct {
	// SenderPublicKeyBase58Check is the public key in base58check of the message sender.
	SenderPublicKeyBase58Check     string `safeForLogging:"true"`

	// RecipientPublicKeyBase58Check is the public key in base58check of the messaging recipient.
	RecipientPublicKeyBase58Check  string `safeForLogging:"true"`

	MessageText                    string // Deprecated

	// EncryptedMessageText is the intended message content. It is recommended to pass actual encrypted message here,
	// although unencrypted message can be passed as well.
	EncryptedMessageText           string

	MinFeeRateNanosPerKB           uint64 `safeForLogging:"true"`
	// No need to specify ProfileEntryResponse in each TransactionFee
	TransactionFees                []TransactionFee `safeForLogging:"true"`

	// ---------------------------------------------------------
	// DeSo V3 Messages Fields
	// ---------------------------------------------------------

	// SenderMessagingGroupKeyName is the messaging group key name of the sender. If left empty, this endpoint
	// will replace it with the base messaging key. If both SenderMessagingGroupKeyName and
	// RecipientMessagingGroupKeyName are left empty, a V2 message will be constructed.
	SenderMessagingGroupKeyName    string `safeForLogging:"true"`

	// RecipientMessagingGroupKeyName is the messaging group key name of the recipient. If left empty, this endpoint
	// will replace it with the base messaging key. If both SenderMessagingGroupKeyName and
	// RecipientMessagingGroupKeyName are left empty, a V2 message will be constructed.
	RecipientMessagingGroupKeyName string `safeForLogging:"true"`
} 
type SendMessageStatelessResponse struct {
	TstampNanos uint64

	TotalInputNanos   uint64
	ChangeAmountNanos uint64
	FeeNanos          uint64
	Transaction       MsgDeSoTxn
	TransactionHex    string
}