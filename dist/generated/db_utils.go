package types

import (
	"github.com/holiman/uint256"
)
  
type AcceptNFTBidTxindexMetadata struct {
	NFTPostHashHex string
	SerialNumber   uint64
	BidAmountNanos uint64
	NFTRoyaltiesMetadata
} 
type AcceptNFTTransferTxindexMetadata struct {
	NFTPostHashHex string
	SerialNumber   uint64
} 
type AffectedPublicKey struct {
	PublicKeyBase58Check string
	// Metadata about how this public key was affected by the transaction.
	Metadata string
} 
type BasicTransferTxindexMetadata struct {
	TotalInputNanos  uint64
	TotalOutputNanos uint64
	FeeNanos         uint64
	UtxoOpsDump      string
	UtxoOps          []*UtxoOperation
	DiamondLevel     int64
	PostHashHex      string
} 
type BitcoinExchangeTxindexMetadata struct {
	BitcoinSpendAddress string
	// DeSoOutputPubKeyBase58Check = TransactorPublicKeyBase58Check
	SatoshisBurned uint64
	// NanosCreated = 0 OR TotalOutputNanos+FeeNanos
	NanosCreated uint64
	// TotalNanosPurchasedBefore = TotalNanosPurchasedAfter - NanosCreated
	TotalNanosPurchasedBefore uint64
	TotalNanosPurchasedAfter  uint64
	BitcoinTxnHash            string
} 
type BurnNFTTxindexMetadata struct {
	NFTPostHashHex string
	SerialNumber   uint64
} 
type ChainType uint8
 
type CreateNFTTxindexMetadata struct {
	NFTPostHashHex             string
	AdditionalCoinRoyaltiesMap map[string]uint64 `json:",omitempty"`
	AdditionalDESORoyaltiesMap map[string]uint64 `json:",omitempty"`
} 
type CreatorCoinTransferTxindexMetadata struct {
	CreatorUsername            string
	CreatorCoinToTransferNanos uint64
	DiamondLevel               int64
	PostHashHex                string
} 
type CreatorCoinTxindexMetadata struct {
	OperationType string
	// TransactorPublicKeyBase58Check = TransactorPublicKeyBase58Check
	// CreatorPublicKeyBase58Check in AffectedPublicKeys

	// Differs depending on OperationType.
	DeSoToSellNanos        uint64
	CreatorCoinToSellNanos uint64
	DeSoToAddNanos         uint64

	// Rosetta needs to know how much DESO was added or removed so it can
	// model the change to the total deso locked in the creator coin
	DESOLockedNanosDiff int64
} 
type DAOCoinTransferTxindexMetadata struct {
	CreatorUsername        string
	DAOCoinToTransferNanos uint256.Int
} 
type DAOCoinTxindexMetadata struct {
	CreatorUsername           string
	OperationType             string
	CoinsToMintNanos          uint256.Int
	CoinsToBurnNanos          uint256.Int
	TransferRestrictionStatus string
} 
type FollowTxindexMetadata struct {
	// FollowerPublicKeyBase58Check = TransactorPublicKeyBase58Check
	// FollowedPublicKeyBase58Check in AffectedPublicKeys

	IsUnfollow bool
} 
type LikeTxindexMetadata struct {
	// LikerPublicKeyBase58Check = TransactorPublicKeyBase58Check
	IsUnlike bool

	PostHashHex string
	// PosterPublicKeyBase58Check in AffectedPublicKeys
} 
type NFTBidTxindexMetadata struct {
	NFTPostHashHex            string
	SerialNumber              uint64
	BidAmountNanos            uint64
	IsBuyNowBid               bool
	OwnerPublicKeyBase58Check string
	// We omit the empty object here as a bid that doesn't trigger a "buy now" operation will have no royalty metadata
	NFTRoyaltiesMetadata `json:",omitempty"`
} 
type NFTRoyaltiesMetadata struct {
	CreatorCoinRoyaltyNanos     uint64
	CreatorRoyaltyNanos         uint64
	CreatorPublicKeyBase58Check string
	// We omit the maps when empty to save some space.
	AdditionalCoinRoyaltiesMap map[string]uint64 `json:",omitempty"`
	AdditionalDESORoyaltiesMap map[string]uint64 `json:",omitempty"`
} 
type NFTTransferTxindexMetadata struct {
	NFTPostHashHex string
	SerialNumber   uint64
} 
type PrivateMessageTxindexMetadata struct {
	// SenderPublicKeyBase58Check = TransactorPublicKeyBase58Check
	// RecipientPublicKeyBase58Check in AffectedPublicKeys

	TimestampNanos uint64
} 
type SubmitPostTxindexMetadata struct {
	PostHashBeingModifiedHex string
	// PosterPublicKeyBase58Check = TransactorPublicKeyBase58Check

	// If this is a reply to an existing post, then the ParentPostHashHex
	ParentPostHashHex string
	// ParentPosterPublicKeyBase58Check in AffectedPublicKeys

	// The profiles that are mentioned are in the AffectedPublicKeys
	// MentionedPublicKeyBase58Check in AffectedPublicKeys
} 
type SwapIdentityTxindexMetadata struct {
	// ParamUpdater = TransactorPublicKeyBase58Check

	FromPublicKeyBase58Check string
	ToPublicKeyBase58Check   string

	// Rosetta needs this information to track creator coin balances
	FromDeSoLockedNanos uint64
	ToDeSoLockedNanos   uint64
} 
type TransactionMetadata struct {
	BlockHashHex    string
	TxnIndexInBlock uint64
	TxnType         string
	// All transactions have a public key who executed the transaction and some
	// public keys that are affected by the transaction. Notifications are created
	// for the affected public keys. _getPublicKeysForTxn uses this to set entries in the
	// database.
	TransactorPublicKeyBase58Check string
	AffectedPublicKeys             []*AffectedPublicKey

	// We store these outputs so we don't have to load the full transaction from disk
	// when looking up output amounts
	TxnOutputs []*DeSoOutput

	BasicTransferTxindexMetadata       *BasicTransferTxindexMetadata       `json:",omitempty"`
	BitcoinExchangeTxindexMetadata     *BitcoinExchangeTxindexMetadata     `json:",omitempty"`
	CreatorCoinTxindexMetadata         *CreatorCoinTxindexMetadata         `json:",omitempty"`
	CreatorCoinTransferTxindexMetadata *CreatorCoinTransferTxindexMetadata `json:",omitempty"`
	UpdateProfileTxindexMetadata       *UpdateProfileTxindexMetadata       `json:",omitempty"`
	SubmitPostTxindexMetadata          *SubmitPostTxindexMetadata          `json:",omitempty"`
	LikeTxindexMetadata                *LikeTxindexMetadata                `json:",omitempty"`
	FollowTxindexMetadata              *FollowTxindexMetadata              `json:",omitempty"`
	PrivateMessageTxindexMetadata      *PrivateMessageTxindexMetadata      `json:",omitempty"`
	SwapIdentityTxindexMetadata        *SwapIdentityTxindexMetadata        `json:",omitempty"`
	NFTBidTxindexMetadata              *NFTBidTxindexMetadata              `json:",omitempty"`
	AcceptNFTBidTxindexMetadata        *AcceptNFTBidTxindexMetadata        `json:",omitempty"`
	NFTTransferTxindexMetadata         *NFTTransferTxindexMetadata         `json:",omitempty"`
	AcceptNFTTransferTxindexMetadata   *AcceptNFTTransferTxindexMetadata   `json:",omitempty"`
	BurnNFTTxindexMetadata             *BurnNFTTxindexMetadata             `json:",omitempty"`
	DAOCoinTxindexMetadata             *DAOCoinTxindexMetadata             `json:",omitempty"`
	DAOCoinTransferTxindexMetadata     *DAOCoinTransferTxindexMetadata     `json:",omitempty"`
	CreateNFTTxindexMetadata           *CreateNFTTxindexMetadata           `json:",omitempty"`
	UpdateNFTTxindexMetadata           *UpdateNFTTxindexMetadata           `json:",omitempty"`
} 
type UpdateNFTTxindexMetadata struct {
	NFTPostHashHex string
	IsForSale      bool
} 
type UpdateProfileTxindexMetadata struct {
	ProfilePublicKeyBase58Check string

	NewUsername    string
	NewDescription string
	NewProfilePic  string

	NewCreatorBasisPoints uint64

	NewStakeMultipleBasisPoints uint64

	IsHidden bool
}