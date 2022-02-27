package types

import (
	"net"
	"time"

	"github.com/holiman/uint256"

	"github.com/btcsuite/btcd/btcec"
	"github.com/btcsuite/btcd/wire"
	merkletree "github.com/deso-protocol/go-merkle-tree"
)
  
type AcceptNFTBidMetadata struct {
	NFTPostHash    *BlockHash
	SerialNumber   uint64
	BidderPKID     *PKID
	BidAmountNanos uint64
	UnlockableText []byte

	// When an NFT owner accepts a bid, they must specify the bidder's UTXO inputs they will lock up
	// as payment for the purchase. This prevents the transaction from accidentally using UTXOs
	// that are used by future transactions.
	BidderInputs []*DeSoInput
} 
type AcceptNFTTransferMetadata struct {
	NFTPostHash  *BlockHash
	SerialNumber uint64
} 
type AuthorizeDerivedKeyMetadata struct {
	// DerivedPublicKey is the key that is authorized to sign transactions
	// on behalf of the public key owner.
	DerivedPublicKey []byte

	// ExpirationBlock is the block at which this authorization becomes invalid.
	ExpirationBlock uint64

	// OperationType determines if transaction validates or invalidates derived key.
	OperationType AuthorizeDerivedKeyOperationType

	// AccessSignature is the signed hash of (derivedPublicKey + expirationBlock)
	// made with the ownerPublicKey. Signature is in the DER format.
	AccessSignature []byte
} 
type AuthorizeDerivedKeyOperationType uint8
 
type BasicTransferMetadata struct {
	// Requires no extra information
} 
type BitcoinExchangeMetadata struct {
	// The Bitcoin transaction that sends Bitcoin to the designated burn address.
	BitcoinTransaction *wire.MsgTx
	// The hash of the Bitcoin block in which the Bitcoin transaction was mined.
	BitcoinBlockHash *BlockHash
	// The Bitcoin mekle root corresponding to the block in which the BitcoinTransaction
	// above was mined. Note that it is not strictly necessary to include this field
	// since we can look it up from the Bitcoin header if we know the BitcoinBlockHash.
	// However, having it here is convenient and allows us to do more validation prior
	// to looking up the header in the Bitcoin header chain.
	BitcoinMerkleRoot *BlockHash
	// This is a merkle proof that shows that the BitcoinTransaction above, with
	// hash equal to BitcoinTransactionHash, exists in the block with hash equal
	// to BitcoinBlockHash. This is effectively a path through a Merkle tree starting
	// from BitcoinTransactionHash as a leaf node and finishing with BitcoinMerkleRoot
	// as the root.
	BitcoinMerkleProof []*merkletree.ProofPart
} 
type BlockProducerInfo struct {
	PublicKey []byte
	Signature *btcec.Signature
} 
type BlockRewardMetadataa struct {
	// A block reward txn has an ExtraData field that can be between
	// zero and 100 bytes long. It can theoretically contain anything
	// but in practice it's likely that miners will use this field to
	// update the merkle root of the block, which may make the block
	// easier to mine (namely by allowing the Nonce in the header to
	// be shorter).
	ExtraData []byte
} 
type BurnNFTMetadata struct {
	NFTPostHash  *BlockHash
	SerialNumber uint64
} 
type CreateNFTMetadata struct {
	NFTPostHash                    *BlockHash
	NumCopies                      uint64
	HasUnlockable                  bool
	IsForSale                      bool
	MinBidAmountNanos              uint64
	NFTRoyaltyToCreatorBasisPoints uint64
	NFTRoyaltyToCoinBasisPoints    uint64
} 
type CreatorCoinMetadataa struct {
	// ProfilePublicKey is the public key of the profile that owns the
	// coin the person wants to operate on. Creator coins can only be
	// bought and sold if a valid profile exists.
	ProfilePublicKey []byte

	// OperationType specifies what the user wants to do with this
	// creator coin.
	OperationType CreatorCoinOperationType

	// Generally, only one of these will be used depending on the OperationType
	// set. In a Buy transaction, DeSoToSellNanos will be converted into
	// creator coin on behalf of the user. In a Sell transaction,
	// CreatorCoinToSellNanos will be converted into DeSo. In an AddDeSo
	// operation, DeSoToAddNanos will be aded for the user. This allows us to
	// support multiple transaction types with same meta field.
	DeSoToSellNanos        uint64
	CreatorCoinToSellNanos uint64
	DeSoToAddNanos         uint64

	// When a user converts DeSo into CreatorCoin, MinCreatorCoinExpectedNanos
	// specifies the minimum amount of creator coin that the user expects from their
	// transaction. And vice versa when a user is converting CreatorCoin for DeSo.
	// Specifying these fields prevents the front-running of users' buy/sell. Setting
	// them to zero turns off the check. Give it your best shot, Ivan.
	MinDeSoExpectedNanos        uint64
	MinCreatorCoinExpectedNanos uint64
} 
type CreatorCoinOperationType uint8
 
type CreatorCoinTransferMetadataa struct {
	// ProfilePublicKey is the public key of the profile that owns the
	// coin the person wants to transer. Creator coins can only be
	// transferred if a valid profile exists.
	ProfilePublicKey []byte

	CreatorCoinToTransferNanos uint64
	ReceiverPublicKey          []byte
} 
type DAOCoinMetadata struct {
	// ProfilePublicKey is the public key of the profile that owns the
	// coin the person wants to operate on.
	ProfilePublicKey []byte

	// OperationType specifies what the user wants to do with this
	// DAO coin.
	OperationType DAOCoinOperationType

	// TODO: Should we only have one field that tracks number of coins in operation to keep this struct small?
	// We will only ever need 1 of these fields.
	// Mint field
	CoinsToMintNanos uint256.Int

	// Burn Fields
	CoinsToBurnNanos uint256.Int

	// TransferRestrictionStatus to set if OperationType == DAOCoinOperatoinTypeUpdateTransferRestrictionStatus
	TransferRestrictionStatus
} 
type DAOCoinOperationType uint8
 
type DAOCoinTransferMetadata struct {
	// ProfilePublicKey is the public key of the profile that owns the
	// coin the person wants to transfer. DAO coins can only be
	// transferred if a valid profile exists.
	ProfilePublicKey []byte

	DAOCoinToTransferNanos uint256.Int
	ReceiverPublicKey      []byte
} 
type DeSoBodySchema struct {
	Body      string   `json:",omitempty"`
	ImageURLs []string `json:",omitempty"`
	VideoURLs []string `json:",omitempty"`
} 
type DeSoInput UtxoKey
 
type DeSoMessage interface {
	// The following methods allow one to convert a message struct into
	// a byte slice and back. Example usage:
	//
	//   params := &DeSoTestnetParams
	//   msgType := MsgTypeVersion
	//   byteSlice := []byte{0x00, ...}
	//
	// 	 msg := NewMessage(msgType)
	//   err := msg.FromBytes(byteSlice)
	//   newByteSlice, err := msg.ToBytes(false)
	//
	// The data format is intended to be compact while allowing for efficient
	// transmission over the wire and storage in a database.
	//
	// The preSignature field specifies whether the message should be fully
	// serialized or whether it should be serialized in such a way that it
	// can be signed (which involves, for example, not serializing signature
	// fields).
	ToBytes(preSignature bool) ([]byte, error)
	FromBytes(data []byte) error

	// Each Message has a particular type.
	GetMsgType() MsgType
} 
type DeSoOutput struct {
	// Outputs always compensate a specific public key.
	PublicKey []byte
	// The amount of DeSo to send to this public key.
	AmountNanos uint64
} 
type DeSoTxnMetadata interface {
	ToBytes(preSignature bool) ([]byte, error)
	FromBytes(data []byte) error
	New() DeSoTxnMetadata
	GetTxnType() TxnType
} 
type FollowMetadata struct {
	// The follower is assumed to be the originator of the
	// top-level transaction.

	// The public key to follow.
	FollowedPublicKey []byte

	// Set to true when a user is requesting to unfollow.
	IsUnfollow bool
} 
type InvType uint32
 
type InvVect struct {
	Type InvType   // Type of data
	Hash BlockHash // Hash of the data
} 
type LikeMetadata struct {
	// The user casting a "like" is assumed to be the originator of the
	// top-level transaction.

	// The post hash to like.
	LikedPostHash *BlockHash

	// Set to true when a user is requesting to unlike a post.
	IsUnlike bool
} 
type MessagingGroupMetadata struct {
	// This struct is very similar to the MessagingGroupEntry type.
	MessagingPublicKey    []byte
	MessagingGroupKeyName []byte
	// This value is the signature of the following using the private key
	// of the GroupOwnerPublicKey (aka txn.PublicKey):
	// - Sha256DoubleHash(MessagingPublicKey || MessagingGroupKeyName)
	//
	// This signature is only required when setting up a group where
	// - MessagingGroupKeyName = "default-key"
	// In this case, we want to make sure that people don't accidentally register
	// this group name with a derived key, and forcing this signature ensures that.
	// The reason is that if someone accidentally registers the default-key with
	// the wrong public key, then they won't be able to receive messages cross-device
	// anymore.
	//
	// This field is not critical and can be removed in the future.
	GroupOwnerSignature   []byte

	MessagingGroupMembers []*MessagingGroupMember
} 
type MsgDeSoAddr struct {
	// The definition of NetAddress as defined by the btcd guys works fine for
	// our purposes. The only difference is that for DeSo nodes, the Service
	// flag in the NetAddress is as we define it above in ServiceFlag.
	// Note that we also rewrite the serialization logic as well to avoid
	// relying on potentially crusty Bitcoin-related work-arounds going forward.
	AddrList []*SingleAddr
} 
type MsgDeSoBitcoinManagerUpdate struct {
	// Keep it simple for now. A BitcoinManagerUpdate just signals that
	// the BitcoinManager has added at least one block or done a reorg.
	// No serialization because we don't want this sent on the wire ever.
	TransactionsFound []*MsgDeSoTxn
} 
type MsgDeSoBlock struct {
	Header *MsgDeSoHeader
	Txns   []*MsgDeSoTxn

	// This field is optional and provides the producer of the block the ability to sign it
	// with their private key. Doing this proves that this block was produced by a particular
	// entity, which can be useful for nodes that want to restrict who they accept blocks
	// from.
	BlockProducerInfo *BlockProducerInfo
} 
type MsgDeSoDonePeer struct {
} 
type MsgDeSoGetAddr struct {
} 
type MsgDeSoGetBlocks struct {
	HashList []*BlockHash
} 
type MsgDeSoGetHeaders struct {
	StopHash     *BlockHash
	BlockLocator []*BlockHash
} 
type MsgDeSoGetTransactions struct {
	HashList []*BlockHash
} 
type MsgDeSoHeader struct {
	// Note this is encoded as a fixed-width uint32 rather than a
	// uvarint or a uint64.
	Version uint32

	// Hash of the previous block in the chain.
	PrevBlockHash *BlockHash

	// The merkle root of all the transactions contained within the block.
	TransactionMerkleRoot *BlockHash

	// The unix timestamp (in seconds) specifying when this block was
	// mined.
	TstampSecs uint64

	// The height of the block this header corresponds to.
	Height uint64

	// The nonce that is used by miners in order to produce valid blocks.
	//
	// Note: Before the upgrade from HeaderVersion0 to HeaderVersion1, miners would make
	// use of ExtraData in the BlockRewardMetadata to get extra nonces. However, this is
	// no longer needed since HeaderVersion1 upgraded the nonce to 64 bits from 32 bits.
	Nonce uint64

	// An extra nonce that can be used to provice *even more* entropy for miners, in the
	// event that ASICs become powerful enough to have birthday problems in the future.
	ExtraNonce uint64
} 
type MsgDeSoHeaderBundle struct {
	Headers   []*MsgDeSoHeader
	TipHash   *BlockHash
	TipHeight uint32
} 
type MsgDeSoInv struct {
	InvList []*InvVect
	// IsSyncResponse indicates that the inv was sent in response to a sync message.
	// This indicates that the node shouldn't relay it to peers because they likely
	// already have it.
	IsSyncResponse bool
} 
type MsgDeSoMempool struct {
} 
type MsgDeSoNewPeer struct {
} 
type MsgDeSoPing struct {
	Nonce uint64
} 
type MsgDeSoPong struct {
	Nonce uint64
} 
type MsgDeSoQuit struct {
} 
type MsgDeSoTransactionBundle struct {
	Transactions []*MsgDeSoTxn
} 
type MsgDeSoTxn struct {
	TxInputs  []*DeSoInput
	TxOutputs []*DeSoOutput

	// DeSoTxnMetadata is an interface type that will give us information on how
	// we should handle the transaction, including what type of transaction this
	// is.
	TxnMeta DeSoTxnMetadata

	// Transactions must generally explicitly include the key that is
	// spending the inputs to the transaction. The exception to this rule is that
	// BlockReward and BitcoinExchange transactions do not require the inclusion
	// of a public key since they have no inputs to spend.
	//
	// The public key should be a serialized compressed ECDSA public key on the
	// secp256k1 curve.
	PublicKey []byte

	// This is typically a JSON field that can be used to add extra information to
	// a transaction without causing a hard fork. It is useful in rare cases where we
	// realize that something needs to be added to a transaction but where we can't
	// afford a hard fork.
	ExtraData map[string][]byte

	// Transactions must generally be signed by the key that is spending the
	// inputs to the transaction. The exception to this rule is that
	// BLOCK_REWARD and CREATE_deso transactions do not require a signature
	// since they have no inputs.
	Signature *btcec.Signature

	// (!!) **DO_NOT_USE** (!!)
	//
	// Use txn.TxnMeta.GetTxnType() instead.
	//
	// We need this for JSON encoding/decoding. It isn't used for anything
	// else and it isn't actually serialized or de-serialized when sent
	// across the network using ToBytes/FromBytes because we prefer that
	// any use of the MsgDeSoTxn in Go code rely on TxnMeta.GetTxnType() rather
	// than checking this value, which, in Go context, is redundant and
	// therefore error-prone (e.g. someone might change TxnMeta while
	// forgetting to set it). We make it a uint64 explicitly to prevent
	// people from using it in Go code.
	TxnTypeJSON uint64
} 
type MsgDeSoVerack struct {
	// A verack message must contain the nonce the peer received in the
	// initial version message. This ensures the peer that is communicating
	// with us actually controls the address she says she does similar to
	// "SYN Cookie" DDOS protection.
	Nonce uint64
} 
type MsgDeSoVersion struct {
	// What is the current version we're on?
	Version uint64

	// What are the services offered by this node?
	Services ServiceFlag

	// The node's unix timestamp that we use to compute a
	// robust "network time" using NTP.
	TstampSecs int64

	// Used to detect when a node connects to itself, which
	// we generally want to prevent.
	Nonce uint64

	// Used as a "vanity plate" to identify different DeSo
	// clients. Mainly useful in analyzing the network at
	// a meta level, not in the protocol itself.
	UserAgent string

	// The height of the last block on the main chain for
	// this node.
	//
	// TODO: We need to update this to uint64
	StartBlockHeight uint32

	// MinFeeRateNanosPerKB is the minimum feerate that a peer will
	// accept from other peers when validating transactions.
	MinFeeRateNanosPerKB uint64
} 
type MsgType uint64
 
type NFTBidMetadata struct {
	NFTPostHash    *BlockHash
	SerialNumber   uint64
	BidAmountNanos uint64
} 
type NFTTransferMetadata struct {
	NFTPostHash       *BlockHash
	SerialNumber      uint64
	ReceiverPublicKey []byte
	UnlockableText    []byte
} 
type PrivateMessageMetadata struct {
	// The sender of the message is assumed to be the originator of the
	// top-level transaction.

	// The public key of the recipient of the message.
	RecipientPublicKey []byte

	// The content of the message. It is encrypted with the recipient's
	// public key using ECIES.
	EncryptedText []byte

	// A timestamp used for ordering messages when displaying them to
	// users. The timestamp must be unique. Note that we use a nanosecond
	// timestamp because it makes it easier to deal with the uniqueness
	// constraint technically (e.g. If one second spacing is required
	// as would be the case with a standard Unix timestamp then any code
	// that generates these transactions will need to potentially wait
	// or else risk a timestamp collision. This complexity is avoided
	// by just using a nanosecond timestamp). Note that the timestamp is
	// an unsigned int as opposed to a signed int, which means times
	// before the zero time are not represented which doesn't matter
	// for our purposes. Restricting the timestamp in this way makes
	// lexicographic sorting based on bytes easier in our database which
	// is one of the reasons we do it.
	TimestampNanos uint64
} 
type ServiceFlag uint64
 
type SingleAddr struct {
	// Last time the address was seen. Encoded as number UNIX seconds on the wire.
	Timestamp time.Time

	// Bitfield which identifies the services supported by the address.
	Services ServiceFlag

	// IP address of the peer. Must be 4 or 16 bytes for IPV4 or IPV6 respectively.
	IP net.IP

	// Port the peer is using.
	Port uint16
} 
type SubmitPostMetadata struct {
	// The creator of the post is assumed to be the originator of the
	// top-level transaction.

	// When set, this transaction is treated as modifying an existing
	// post rather than creating a new post.
	PostHashToModify []byte

	// When a ParentStakeID is set, the post is actually a comment on
	// another entity (either a post or a profile depending on the
	// type of StakeID provided).
	ParentStakeID []byte
	Body          []byte

	// The amount the creator of the post gets when someone stakes
	// to the post.
	CreatorBasisPoints uint64
	// The multiple of the payout when a user stakes to a post.
	// 2x multiple = 200% = 20,000bps
	StakeMultipleBasisPoints uint64

	// A timestamp used for ordering messages when displaying them to
	// users. The timestamp must be unique. Note that we use a nanosecond
	// timestamp because it makes it easier to deal with the uniqueness
	// constraint technically (e.g. If one second spacing is required
	// as would be the case with a standard Unix timestamp then any code
	// that generates these transactions will need to potentially wait
	// or else risk a timestamp collision. This complexity is avoided
	// by just using a nanosecond timestamp). Note that the timestamp is
	// an unsigned int as opposed to a signed int, which means times
	// before the zero time are not represented which doesn't matter
	// for our purposes. Restricting the timestamp in this way makes
	// lexicographic sorting based on bytes easier in our database which
	// is one of the reasons we do it.
	TimestampNanos uint64

	// When set to true, indicates that the post should be deleted. This
	// value is only considered when PostHashToModify is set to a valid
	// pre-existing post.
	IsHidden bool
} 
type SwapIdentityOperationType uint8
 
type SwapIdentityOperationType uint8

type SwapIdentityMetadataa struct {
	// TODO: This is currently only accessible by ParamUpdater. This avoids the
	// possibility that a user will stomp over another user's profile, and
	// simplifies the logic. In the long run, though, we should eliminate all
	// dependencies on ParamUpdater.

	// The public key that we are swapping *from*. Doesn't matter which public
	// key is *from* and which public key is *to* because it's just a swap.
	FromPublicKey []byte

	// The public key that we are swapping *to*. Doesn't matter which public
	// key is *from* and which public key is *to* because it's just a swap.
	ToPublicKey []byte
} 
type TxnString string 
type TxnType uint8
 
type UpdateBitcoinUSDExchangeRateMetadataa struct {
	// The new exchange rate to set.
	USDCentsPerBitcoin uint64
} 
type UpdateGlobalParamsMetadata struct {
	// The GlobalParamsMetadata struct is empty because all information is stored in the transaction's ExtraData
} 
type UpdateNFTMetadata struct {
	NFTPostHash       *BlockHash
	SerialNumber      uint64
	IsForSale         bool
	MinBidAmountNanos uint64
} 
type UpdateProfileMetadata struct {
	// The public key being updated is assumed to be the originator of the
	// top-level transaction.

	// The public key of the profile to update. When left unset, the public
	// key in the transaction is used.
	ProfilePublicKey []byte

	NewUsername    []byte
	NewDescription []byte
	NewProfilePic  []byte

	// This is the percentage of each "net buy" that a creator earns when
	// someone purchases her coin. For example, if this were set to 25%,
	// then every time their coin reaches a new high, they would get 25%
	// of the coins as they're being minted. More concretely, if someone
	// put in enough DeSo to buy 10 coins, the creator would get 2.5
	// and this person would get 7.5. However, if they sold 5 coins and
	// someone subsequently bought those same coins, the creator wouldn't
	// get any coins because no "net new" coins have been created.
	NewCreatorBasisPoints uint64

	// The multiple of the payout when a user stakes to this profile. If
	// unset, a sane default is set when the first person stakes to this
	// profile.
	// 2x multiple = 200% = 20,000bps
	//
	// TODO: This field is deprecated; delete it.
	NewStakeMultipleBasisPoints uint64

	// Profile is hidden from the UI when this field is true.
	// TODO: This field is deprecated; delete it.
	IsHidden bool
} 
type UtxoKey struct {
	// The 32-byte transaction id where the unspent output occurs.
	TxID BlockHash
	// The index within the txn where the unspent output occurs.
	Index uint32
}