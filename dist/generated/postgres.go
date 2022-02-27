package types

import (
	"github.com/go-pg/pg/v10"
)
 
type NotificationType uint8
 
type PGBalance struct {
	tableName struct{} `pg:"pg_balances"`

	PublicKey    *PublicKey `pg:",pk,type:bytea"`
	BalanceNanos uint64     `pg:",use_zero"`
} 
type PGBlock struct {
	tableName struct{} `pg:"pg_blocks"`

	// BlockNode and MsgDeSoHeader
	Hash       *BlockHash `pg:",pk,type:bytea"`
	ParentHash *BlockHash `pg:",type:bytea"`
	Height     uint64     `pg:",use_zero"`

	// BlockNode
	DifficultyTarget *BlockHash  `pg:",type:bytea"`
	CumWork          *BlockHash  `pg:",type:bytea"`
	Status           BlockStatus `pg:",use_zero"` // TODO: Refactor

	// MsgDeSoHeader
	TxMerkleRoot *BlockHash `pg:",type:bytea"`
	Version      uint32     `pg:",use_zero"`
	Timestamp    uint64     `pg:",use_zero"`
	Nonce        uint64     `pg:",use_zero"`
	ExtraNonce   uint64     `pg:",use_zero"`

	// Notifications
	Notified bool `pg:",use_zero"`
} 
type PGChain struct {
	tableName struct{} `pg:"pg_chains"`

	Name    string     `pg:",pk"`
	TipHash *BlockHash `pg:",type:bytea"`
} 
type PGCreatorCoinBalance struct {
	tableName struct{} `pg:"pg_creator_coin_balances"`

	HolderPKID   *PKID `pg:",pk,type:bytea"`
	CreatorPKID  *PKID `pg:",pk,type:bytea"`
	BalanceNanos uint64
	HasPurchased bool
} 
type PGDAOCoinBalance struct {
	tableName struct{} `pg:"pg_dao_coin_balances"`

	HolderPKID   *PKID `pg:",pk,type:bytea"`
	CreatorPKID  *PKID `pg:",pk,type:bytea"`
	BalanceNanos string
	HasPurchased bool
} 
type PGDerivedKey struct {
	tableName struct{} `pg:"pg_derived_keys"`

	OwnerPublicKey   PublicKey                        `pg:",pk,type:bytea"`
	DerivedPublicKey PublicKey                        `pg:",pk,type:bytea"`
	ExpirationBlock  uint64                           `pg:",use_zero"`
	OperationType    AuthorizeDerivedKeyOperationType `pg:",use_zero"`
} 
type PGDiamond struct {
	tableName struct{} `pg:"pg_diamonds"`

	SenderPKID      *PKID      `pg:",pk,type:bytea"`
	ReceiverPKID    *PKID      `pg:",pk,type:bytea"`
	DiamondPostHash *BlockHash `pg:",pk,type:bytea"`
	DiamondLevel    uint8
} 
type PGFollow struct {
	tableName struct{} `pg:"pg_follows"`

	FollowerPKID *PKID `pg:",pk,type:bytea"`
	FollowedPKID *PKID `pg:",pk,type:bytea"`
} 
type PGForbiddenKey struct {
	tableName struct{} `pg:"pg_forbidden_keys"`

	PublicKey *PublicKey `pg:",pk,type:bytea"`
} 
type PGGlobalParams struct {
	tableName struct{} `pg:"pg_global_params"`

	ID uint64

	USDCentsPerBitcoin      uint64 `pg:",use_zero"`
	CreateProfileFeeNanos   uint64 `pg:",use_zero"`
	CreateNFTFeeNanos       uint64 `pg:",use_zero"`
	MaxCopiesPerNFT         uint64 `pg:",use_zero"`
	MinNetworkFeeNanosPerKB uint64 `pg:",use_zero"`
} 
type PGLike struct {
	tableName struct{} `pg:"pg_likes"`

	LikerPublicKey []byte     `pg:",pk,type:bytea"`
	LikedPostHash  *BlockHash `pg:",pk,type:bytea"`
} 
type PGMessage struct {
	tableName struct{} `pg:"pg_messages"`

	MessageHash        *BlockHash `pg:",pk,type:bytea"`
	SenderPublicKey    []byte
	RecipientPublicKey []byte
	EncryptedText      []byte
	TimestampNanos     uint64
	// TODO: Version

	// Used to track deletions in the UtxoView
	isDeleted bool
} 
type PGMessagingGroup struct {
	tableName struct{} `pg:"pg_messaging_group"`

	GroupOwnerPublicKey   *PublicKey    `pg:",type:bytea"`
	MessagingPublicKey    *PublicKey    `pg:",type:bytea"`
	MessagingGroupKeyName *GroupKeyName `pg:",type:bytea"`
	MessagingGroupMembers []byte        `pg:",type:bytea"`
} 
type PGMetadataAcceptNFTBid struct {
	tableName struct{} `pg:"pg_metadata_accept_nft_bids"`

	TransactionHash *BlockHash            `pg:",pk,type:bytea"`
	NFTPostHash     *BlockHash            `pg:",type:bytea"`
	SerialNumber    uint64                `pg:",use_zero"`
	BidderPKID      *PKID                 `pg:",type:bytea"`
	BidAmountNanos  uint64                `pg:",use_zero"`
	UnlockableText  []byte                `pg:",type:bytea"`
	BidderInputs    []*PGMetadataBidInput `pg:"rel:has-many,join_fk:transaction_hash"`
} 
type PGMetadataAcceptNFTTransfer struct {
	tableName struct{} `pg:"pg_metadata_accept_nft_transfer"`

	TransactionHash *BlockHash `pg:",pk,type:bytea"`
	NFTPostHash     *BlockHash `pg:",pk,type:bytea"`
	SerialNumber    uint64     `pg:",use_zero"`
} 
type PGMetadataBidInput struct {
	tableName struct{} `pg:"pg_metadata_bid_inputs"`

	TransactionHash *BlockHash `pg:",pk,type:bytea"`
	InputHash       *BlockHash `pg:",pk,type:bytea"`
	InputIndex      uint32     `pg:",pk,use_zero"`
} 
type PGMetadataBitcoinExchange struct {
	tableName struct{} `pg:"pg_metadata_bitcoin_exchanges"`

	TransactionHash   *BlockHash `pg:",pk,type:bytea"`
	BitcoinBlockHash  *BlockHash `pg:",type:bytea"`
	BitcoinMerkleRoot *BlockHash `pg:",type:bytea"`
	// Not storing BitcoinTransaction *wire.MsgTx
	// Not storing BitcoinMerkleProof []*merkletree.ProofPart
} 
type PGMetadataBlockReward struct {
	tableName struct{} `pg:"pg_metadata_block_rewards"`

	TransactionHash *BlockHash `pg:",pk,type:bytea"`
	ExtraData       []byte     `pg:",type:bytea"`
} 
type PGMetadataBurnNFT struct {
	tableName struct{} `pg:"pg_metadata_burn_nft"`

	TransactionHash *BlockHash `pg:",pk,type:bytea"`
	NFTPostHash     *BlockHash `pg:",pk,type:bytea"`
	SerialNumber    uint64     `pg:",use_zero"`
} 
type PGMetadataCreateNFT struct {
	tableName struct{} `pg:"pg_metadata_create_nfts"`

	TransactionHash           *BlockHash `pg:",pk,type:bytea"`
	NFTPostHash               *BlockHash `pg:",type:bytea"`
	NumCopies                 uint64     `pg:",use_zero"`
	HasUnlockable             bool       `pg:",use_zero"`
	IsForSale                 bool       `pg:",use_zero"`
	MinBidAmountNanos         uint64     `pg:",use_zero"`
	CreatorRoyaltyBasisPoints uint64     `pg:",use_zero"`
	CoinRoyaltyBasisPoints    uint64     `pg:",use_zero"`
} 
type PGMetadataCreatorCoin struct {
	tableName struct{} `pg:"pg_metadata_creator_coins"`

	TransactionHash             *BlockHash               `pg:",pk,type:bytea"`
	ProfilePublicKey            []byte                   `pg:",type:bytea"`
	OperationType               CreatorCoinOperationType `pg:",use_zero"`
	DeSoToSellNanos             uint64                   `pg:",use_zero"`
	CreatorCoinToSellNanos      uint64                   `pg:",use_zero"`
	DeSoToAddNanos              uint64                   `pg:",use_zero"`
	MinDeSoExpectedNanos        uint64                   `pg:",use_zero"`
	MinCreatorCoinExpectedNanos uint64                   `pg:",use_zero"`
} 
type PGMetadataCreatorCoinTransfer struct {
	tableName struct{} `pg:"pg_metadata_creator_coin_transfers"`

	TransactionHash            *BlockHash `pg:",pk,type:bytea"`
	ProfilePublicKey           []byte     `pg:",type:bytea"`
	CreatorCoinToTransferNanos uint64     `pg:",use_zero"`
	ReceiverPublicKey          []byte     `pg:",type:bytea"`
} 
type PGMetadataDAOCoin struct {
	tableName struct{} `pg:"pg_metadata_dao_coins"`

	TransactionHash           *BlockHash           `pg:",pk,type:bytea"`
	ProfilePublicKey          []byte               `pg:",type:bytea"`
	OperationType             DAOCoinOperationType `pg:",use_zero"`
	CoinsToMintNanos          string
	CoinsToBurnNanos          string
	TransferRestrictionStatus `pg:",use_zero"`
} 
type PGMetadataDAOCoinTransfer struct {
	tableName struct{} `pg:"pg_metadata_dao_coin_transfers"`

	TransactionHash        *BlockHash `pg:",pk,type:bytea"`
	ProfilePublicKey       []byte     `pg:",type:bytea"`
	DAOCoinToTransferNanos string     `pg:"dao_coin_to_transfer_nanos,use_zero"`
	ReceiverPublicKey      []byte     `pg:",type:bytea"`
} 
type PGMetadataDerivedKey struct {
	tableName struct{} `pg:"pg_metadata_derived_keys"`

	TransactionHash  *BlockHash                       `pg:",pk,type:bytea"`
	DerivedPublicKey PublicKey                        `pg:",type:bytea"`
	ExpirationBlock  uint64                           `pg:",use_zero"`
	OperationType    AuthorizeDerivedKeyOperationType `pg:",use_zero"`
	AccessSignature  []byte                           `pg:",type:bytea"`
} 
type PGMetadataFollow struct {
	tableName struct{} `pg:"pg_metadata_follows"`

	TransactionHash   *BlockHash `pg:",pk,type:bytea"`
	FollowedPublicKey []byte     `pg:",type:bytea"`
	IsUnfollow        bool       `pg:",use_zero"`
} 
type PGMetadataLike struct {
	tableName struct{} `pg:"pg_metadata_likes"`

	TransactionHash *BlockHash `pg:",pk,type:bytea"`
	LikedPostHash   *BlockHash `pg:",type:bytea"`
	IsUnlike        bool       `pg:",use_zero"`
} 
type PGMetadataNFTBid struct {
	tableName struct{} `pg:"pg_metadata_nft_bids"`

	TransactionHash *BlockHash `pg:",pk,type:bytea"`
	NFTPostHash     *BlockHash `pg:",type:bytea"`
	SerialNumber    uint64     `pg:",use_zero"`
	BidAmountNanos  uint64     `pg:",use_zero"`
} 
type PGMetadataNFTTransfer struct {
	tableName struct{} `pg:"pg_metadata_nft_transfer"`

	TransactionHash   *BlockHash `pg:",pk,type:bytea"`
	NFTPostHash       *BlockHash `pg:",pk,type:bytea"`
	SerialNumber      uint64     `pg:",use_zero"`
	ReceiverPublicKey []byte     `pg:",pk,type:bytea"`
	UnlockableText    []byte     `pg:",type:bytea"`
} 
type PGMetadataPrivateMessage struct {
	tableName struct{} `pg:"pg_metadata_private_messages"`

	TransactionHash    *BlockHash `pg:",pk,type:bytea"`
	RecipientPublicKey []byte     `pg:",type:bytea"`
	EncryptedText      []byte     `pg:",type:bytea"`
	TimestampNanos     uint64
} 
type PGMetadataSubmitPost struct {
	tableName struct{} `pg:"pg_metadata_submit_posts"`

	TransactionHash  *BlockHash `pg:",pk,type:bytea"`
	PostHashToModify *BlockHash `pg:",type:bytea"`
	ParentStakeID    *BlockHash `pg:",type:bytea"`
	Body             []byte     `pg:",type:bytea"`
	TimestampNanos   uint64
	IsHidden         bool `pg:",use_zero"`
} 
type PGMetadataSwapIdentity struct {
	tableName struct{} `pg:"pg_metadata_swap_identities"`

	TransactionHash *BlockHash `pg:",pk,type:bytea"`
	FromPublicKey   []byte     `pg:",type:bytea"`
	ToPublicKey     []byte     `pg:",type:bytea"`
} 
type PGMetadataUpdateExchangeRate struct {
	tableName struct{} `pg:"pg_metadata_update_exchange_rates"`

	TransactionHash    *BlockHash `pg:",pk,type:bytea"`
	USDCentsPerBitcoin uint64     `pg:",use_zero"`
} 
type PGMetadataUpdateNFT struct {
	tableName struct{} `pg:"pg_metadata_update_nfts"`

	TransactionHash   *BlockHash `pg:",pk,type:bytea"`
	NFTPostHash       *BlockHash `pg:",type:bytea"`
	SerialNumber      uint64     `pg:",use_zero"`
	IsForSale         bool       `pg:",use_zero"`
	MinBidAmountNanos uint64     `pg:",use_zero"`
} 
type PGMetadataUpdateProfile struct {
	tableName struct{} `pg:"pg_metadata_update_profiles"`

	TransactionHash       *BlockHash `pg:",pk,type:bytea"`
	ProfilePublicKey      []byte     `pg:",type:bytea"`
	NewUsername           []byte     `pg:",type:bytea"`
	NewDescription        []byte     `pg:",type:bytea"`
	NewProfilePic         []byte     `pg:",type:bytea"`
	NewCreatorBasisPoints uint64     `pg:",use_zero"`
} 
type PGNFT struct {
	tableName struct{} `pg:"pg_nfts"`

	NFTPostHash  *BlockHash `pg:",pk,type:bytea"`
	SerialNumber uint64     `pg:",pk"`

	// This is needed to decrypt unlockable text.
	LastOwnerPKID              *PKID  `pg:",type:bytea"`
	OwnerPKID                  *PKID  `pg:",type:bytea"`
	ForSale                    bool   `pg:",use_zero"`
	MinBidAmountNanos          uint64 `pg:",use_zero"`
	UnlockableText             string
	LastAcceptedBidAmountNanos uint64 `pg:",use_zero"`
	IsPending                  bool   `pg:",use_zero"`
	IsBuyNow                   bool   `pg:",use_zero"`
	BuyNowPriceNanos           uint64 `pg:",use_zero"`
} 
type PGNFTBid struct {
	tableName struct{} `pg:"pg_nft_bids"`

	BidderPKID          *PKID      `pg:",pk,type:bytea"`
	NFTPostHash         *BlockHash `pg:",pk,type:bytea"`
	SerialNumber        uint64     `pg:",pk,use_zero"`
	BidAmountNanos      uint64     `pg:",use_zero"`
	Accepted            bool       `pg:",use_zero"`
	AcceptedBlockHeight *uint32    `pg:",use_zero"`
} 
type PGNotification struct {
	tableName struct{} `pg:"pg_notifications"`

	TransactionHash *BlockHash       `pg:",pk,type:bytea"`
	Mined           bool             `pg:",use_zero"`
	ToUser          []byte           `pg:",type:bytea"`
	FromUser        []byte           `pg:",type:bytea"`
	OtherUser       []byte           `pg:",type:bytea"`
	Type            NotificationType `pg:",use_zero"`
	Amount          uint64           `pg:",use_zero"`
	PostHash        *BlockHash       `pg:",type:bytea"`
	Timestamp       uint64           `pg:",use_zero"`
} 
type PGPost struct {
	tableName struct{} `pg:"pg_posts"`

	PostHash                                    *BlockHash `pg:",pk,type:bytea"`
	PosterPublicKey                             []byte
	ParentPostHash                              *BlockHash `pg:",type:bytea"`
	Body                                        string
	RepostedPostHash                            *BlockHash        `pg:",type:bytea"`
	QuotedRepost                                bool              `pg:",use_zero"`
	Timestamp                                   uint64            `pg:",use_zero"`
	Hidden                                      bool              `pg:",use_zero"`
	LikeCount                                   uint64            `pg:",use_zero"`
	RepostCount                                 uint64            `pg:",use_zero"`
	QuoteRepostCount                            uint64            `pg:",use_zero"`
	DiamondCount                                uint64            `pg:",use_zero"`
	CommentCount                                uint64            `pg:",use_zero"`
	Pinned                                      bool              `pg:",use_zero"`
	NFT                                         bool              `pg:",use_zero"`
	NumNFTCopies                                uint64            `pg:",use_zero"`
	NumNFTCopiesForSale                         uint64            `pg:",use_zero"`
	NumNFTCopiesBurned                          uint64            `pg:",use_zero"`
	Unlockable                                  bool              `pg:",use_zero"`
	CreatorRoyaltyBasisPoints                   uint64            `pg:",use_zero"`
	CoinRoyaltyBasisPoints                      uint64            `pg:",use_zero"`
	AdditionalNFTRoyaltiesToCoinsBasisPoints    map[string]uint64 `pg:"additional_nft_royalties_to_coins_basis_points"`
	AdditionalNFTRoyaltiesToCreatorsBasisPoints map[string]uint64 `pg:"additional_nft_royalties_to_creators_basis_points"`
	ExtraData                                   map[string][]byte
} 
type PGProfile struct {
	tableName struct{} `pg:"pg_profiles"`

	PKID               *PKID      `pg:",pk,type:bytea"`
	PublicKey          *PublicKey `pg:",type:bytea"`
	Username           string
	Description        string
	ProfilePic         []byte
	CreatorBasisPoints uint64
	DeSoLockedNanos    uint64
	NumberOfHolders    uint64
	// FIXME: Postgres will break when values exceed uint64
	// We don't use Postgres right now so going to plow ahead and set this as-is
	// to fix compile errors. CoinsInCirculationNanos will never exceed uint64
	CoinsInCirculationNanos          uint64
	CoinWatermarkNanos               uint64
	MintingDisabled                  bool
	DAOCoinNumberOfHolders           uint64                    `pg:"dao_coin_number_of_holders"`
	DAOCoinCoinsInCirculationNanos   string                    `pg:"dao_coin_coins_in_circulation_nanos"`
	DAOCoinMintingDisabled           bool                      `pg:"dao_coin_minting_disabled"`
	DAOCoinTransferRestrictionStatus TransferRestrictionStatus `pg:"dao_coin_transfer_restriction_status"`
} 
type PGRepost struct {
	tableName struct{} `pg:"pg_reposts"`

	ReposterPublickey *PublicKey `pg:",pk,type:bytea"`
	RepostedPostHash  *BlockHash `pg:",pk,type:bytea"`
	RepostPostHash    *BlockHash `pg:",type:bytea"`

	// Whether or not this entry is deleted in the view.
	isDeleted bool
} 
type PGTransaction struct {
	tableName struct{} `pg:"pg_transactions"`

	Hash      *BlockHash `pg:",pk,type:bytea"`
	BlockHash *BlockHash `pg:",type:bytea"`
	Type      TxnType    `pg:",use_zero"`
	PublicKey []byte     `pg:",type:bytea"`
	ExtraData map[string][]byte
	R         *BlockHash `pg:",type:bytea"`
	S         *BlockHash `pg:",type:bytea"`

	// Relationships
	Outputs                     []*PGTransactionOutput         `pg:"rel:has-many,join_fk:output_hash"`
	MetadataBlockReward         *PGMetadataBlockReward         `pg:"rel:belongs-to,join_fk:transaction_hash"`
	MetadataBitcoinExchange     *PGMetadataBitcoinExchange     `pg:"rel:belongs-to,join_fk:transaction_hash"`
	MetadataPrivateMessage      *PGMetadataPrivateMessage      `pg:"rel:belongs-to,join_fk:transaction_hash"`
	MetadataSubmitPost          *PGMetadataSubmitPost          `pg:"rel:belongs-to,join_fk:transaction_hash"`
	MetadataUpdateExchangeRate  *PGMetadataUpdateExchangeRate  `pg:"rel:belongs-to,join_fk:transaction_hash"`
	MetadataUpdateProfile       *PGMetadataUpdateProfile       `pg:"rel:belongs-to,join_fk:transaction_hash"`
	MetadataFollow              *PGMetadataFollow              `pg:"rel:belongs-to,join_fk:transaction_hash"`
	MetadataLike                *PGMetadataLike                `pg:"rel:belongs-to,join_fk:transaction_hash"`
	MetadataCreatorCoin         *PGMetadataCreatorCoin         `pg:"rel:belongs-to,join_fk:transaction_hash"`
	MetadataCreatorCoinTransfer *PGMetadataCreatorCoinTransfer `pg:"rel:belongs-to,join_fk:transaction_hash"`
	MetadataSwapIdentity        *PGMetadataSwapIdentity        `pg:"rel:belongs-to,join_fk:transaction_hash"`
	MetadataCreateNFT           *PGMetadataCreateNFT           `pg:"rel:belongs-to,join_fk:transaction_hash"`
	MetadataUpdateNFT           *PGMetadataUpdateNFT           `pg:"rel:belongs-to,join_fk:transaction_hash"`
	MetadataAcceptNFTBid        *PGMetadataAcceptNFTBid        `pg:"rel:belongs-to,join_fk:transaction_hash"`
	MetadataNFTBid              *PGMetadataNFTBid              `pg:"rel:belongs-to,join_fk:transaction_hash"`
	MetadataNFTTransfer         *PGMetadataNFTTransfer         `pg:"rel:belongs-to,join_fk:transaction_hash"`
	MetadataAcceptNFTTransfer   *PGMetadataAcceptNFTTransfer   `pg:"rel:belongs-to,join_fk:transaction_hash"`
	MetadataBurnNFT             *PGMetadataBurnNFT             `pg:"rel:belongs-to,join_fk:transaction_hash"`
	MetadataDerivedKey          *PGMetadataDerivedKey          `pg:"rel:belongs-to,join_fk:transaction_hash"`
	MetadataDAOCoin             *PGMetadataDAOCoin             `pg:"rel:belongs-to,join_fk:transaction_hash"`
	MetadataDAOCoinTransfer     *PGMetadataDAOCoinTransfer     `pg:"rel:belongs-to,join_fk:transaction_hash"`
} 
type PGTransactionOutput struct {
	tableName struct{} `pg:"pg_transaction_outputs"`

	OutputHash  *BlockHash `pg:",pk"`
	OutputIndex uint32     `pg:",pk,use_zero"`
	OutputType  UtxoType   `pg:",use_zero"`
	Height      uint32     `pg:",use_zero"`
	PublicKey   []byte
	AmountNanos uint64 `pg:",use_zero"`
	Spent       bool   `pg:",use_zero"`
	InputHash   *BlockHash
	InputIndex  uint32 `pg:",pk,use_zero"`
} 
type Postgres struct {
	db *pg.DB
}