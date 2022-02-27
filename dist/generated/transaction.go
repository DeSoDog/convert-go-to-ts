package types

import (
	"github.com/holiman/uint256"

	"github.com/btcsuite/btcd/wire"
)
  
type AppendExtraDataRequest struct {
	// Transaction hex.
	TransactionHex string `safeForLogging:"true"`

	// ExtraData object.
	ExtraData map[string]string `safeForLogging:"true"`
} 
type AppendExtraDataResponse struct {
	// Final Transaction hex.
	TransactionHex string `safeForLogging:"true"`
} 
type AuthorizeDerivedKeyRequest struct {
	// The original public key of the derived key owner.
	OwnerPublicKeyBase58Check string `safeForLogging:"true"`

	// The derived public key
	DerivedPublicKeyBase58Check string `safeForLogging:"true"`

	// The expiration block of the derived key pair.
	ExpirationBlock uint64 `safeForLogging:"true"`

	// The signature of hash(derived key + expiration block) made by the owner.
	AccessSignature string `safeForLogging:"true"`

	// The intended operation on the derived key.
	DeleteKey bool `safeForLogging:"true"`

	// If we intend to sign this transaction with a derived key.
	DerivedKeySignature bool `safeForLogging:"true"`

	// No need to specify ProfileEntryResponse in each TransactionFee
	TransactionFees []TransactionFee `safeForLogging:"true"`

	MinFeeRateNanosPerKB uint64 `safeForLogging:"true"`
} 
type AuthorizeDerivedKeyResponse struct {
	SpendAmountNanos  uint64
	TotalInputNanos   uint64
	ChangeAmountNanos uint64
	FeeNanos          uint64
	Transaction       MsgDeSoTxn
	TransactionHex    string
	TxnHashHex        string
} 
type BuyOrSellCreatorCoinRequest struct {
	// The public key of the user who is making the buy/sell.
	UpdaterPublicKeyBase58Check string `safeForLogging:"true"`

	// The public key of the profile that the purchaser is trying
	// to buy.
	CreatorPublicKeyBase58Check string `safeForLogging:"true"`

	// Whether this is a "buy" or "sell"
	OperationType string `safeForLogging:"true"`

	// Generally, only one of these will be used depending on the OperationType
	// set. In a Buy transaction, DeSoToSellNanos will be converted into
	// creator coin on behalf of the user. In a Sell transaction,
	// CreatorCoinToSellNanos will be converted into DeSo. In an AddDeSo
	// operation, DeSoToAddNanos will be aded for the user. This allows us to
	// support multiple transaction types with same meta field.
	DeSoToSellNanos        uint64 `safeForLogging:"true"`
	CreatorCoinToSellNanos uint64 `safeForLogging:"true"`
	DeSoToAddNanos         uint64 `safeForLogging:"true"`

	// When a user converts DeSo into CreatorCoin, MinCreatorCoinExpectedNanos
	// specifies the minimum amount of creator coin that the user expects from their
	// transaction. And vice versa when a user is converting CreatorCoin for DeSo.
	// Specifying these fields prevents the front-running of users' buy/sell. Setting
	// them to zero turns off the check. Give it your best shot, Ivan.
	MinDeSoExpectedNanos        uint64 `safeForLogging:"true"`
	MinCreatorCoinExpectedNanos uint64 `safeForLogging:"true"`

	MinFeeRateNanosPerKB uint64 `safeForLogging:"true"`

	// No need to specify ProfileEntryResponse in each TransactionFee
	TransactionFees []TransactionFee `safeForLogging:"true"`

	InTutorial bool `safeForLogging:"true"`

	BitCloutToSellNanos      uint64 `safeForLogging:"true"` // Deprecated
	BitCloutToAddNanos       uint64 `safeForLogging:"true"` // Deprecated
	MinBitCloutExpectedNanos uint64 `safeForLogging:"true"` // Deprecated
} 
type BuyOrSellCreatorCoinResponse struct {
	// The amount of DeSo
	ExpectedDeSoReturnedNanos        uint64
	ExpectedCreatorCoinReturnedNanos uint64
	FounderRewardGeneratedNanos      uint64

	// Spend is defined as DeSo that's specified as input that winds up as "output not
	// belonging to you." In the case of a creator coin sell, your input is creator coin (not
	// DeSo), so this ends up being 0. In the case of a creator coin buy,
	// it should equal the amount of DeSo you put in to buy the creator coin
	SpendAmountNanos  uint64
	TotalInputNanos   uint64
	ChangeAmountNanos uint64
	FeeNanos          uint64
	Transaction       MsgDeSoTxn
	TransactionHex    string
	TxnHashHex        string
} 
type CreateFollowTxnStatelessRequest struct {
	FollowerPublicKeyBase58Check string `safeForLogging:"true"`
	FollowedPublicKeyBase58Check string `safeForLogging:"true"`
	IsUnfollow                   bool   `safeForLogging:"true"`
	MinFeeRateNanosPerKB         uint64 `safeForLogging:"true"`

	// No need to specify ProfileEntryResponse in each TransactionFee
	TransactionFees []TransactionFee `safeForLogging:"true"`
} 
type CreateFollowTxnStatelessResponse struct {
	TotalInputNanos   uint64
	ChangeAmountNanos uint64
	FeeNanos          uint64
	Transaction       MsgDeSoTxn
	TransactionHex    string
} 
type CreateLikeStatelessRequest struct {
	ReaderPublicKeyBase58Check string `safeForLogging:"true"`
	LikedPostHashHex           string `safeForLogging:"true"`
	IsUnlike                   bool   `safeForLogging:"true"`
	MinFeeRateNanosPerKB       uint64 `safeForLogging:"true"`

	// No need to specify ProfileEntryResponse in each TransactionFee
	TransactionFees []TransactionFee `safeForLogging:"true"`
} 
type CreateLikeStatelessResponse struct {
	TotalInputNanos   uint64
	ChangeAmountNanos uint64
	FeeNanos          uint64
	Transaction       MsgDeSoTxn
	TransactionHex    string
} 
type DAOCoinOperationTypeString string 
type DAOCoinRequest struct {
	// The public key of the user who is performing the DAOCoin Txn
	UpdaterPublicKeyBase58Check string `safeForLogging:"true"`

	// The public key or username of the profile whose DAO coin the transactor is trying to transact with.
	ProfilePublicKeyBase58CheckOrUsername string `safeForLogging:"true"`

	// Whether this is a "mint", "burn" or "disable_minting" transaction
	OperationType DAOCoinOperationTypeString `safeForLogging:"true"`

	// Coins
	CoinsToMintNanos uint256.Int `safeForLogging:"true"`

	CoinsToBurnNanos uint256.Int `safeForLogging:"true"`

	// Transfer Restriction Status
	TransferRestrictionStatus TransferRestrictionStatusString `safeForLogging:"true"`

	MinFeeRateNanosPerKB uint64 `safeForLogging:"true"`

	// No need to specify ProfileEntryResponse in each TransactionFee
	TransactionFees []TransactionFee `safeForLogging:"true"`
} 
type DAOCoinResponse struct {
	TotalInputNanos   uint64
	ChangeAmountNanos uint64
	FeeNanos          uint64
	Transaction       MsgDeSoTxn
	TransactionHex    string
	TxnHashHex        string
} 
type ExchangeBitcoinRequest struct {
	// The public key of the user who we're creating the burn for.
	PublicKeyBase58Check string `safeForLogging:"true"`
	// Note: When BurnAmountSatoshis is negative, we assume that the user wants
	// to burn the maximum amount of satoshi she has available.
	BurnAmountSatoshis   int64 `safeForLogging:"true"`
	FeeRateSatoshisPerKB int64 `safeForLogging:"true"`

	// We rely on the frontend to query the API and give us the response.
	// Doing it this way makes it so that we don't exhaust our quota on the
	// free tier.
	LatestBitcionAPIResponse BlockCypherAPIFullAddressResponse
	// The Bitcoin address we will be processing this transaction for.
	BTCDepositAddress string `safeForLogging:"true"`

	// Whether or not we should broadcast the transaction after constructing
	// it. This will also validate the transaction if it's set.
	// The client must provide SignedHashes which it calculates by signing
	// all the UnsignedHashes in the identity service
	Broadcast bool `safeForLogging:"true"`

	// Signed hashes from the identity service
	// One for each transaction input
	SignedHashes []string
} 
type ExchangeBitcoinResponse struct {
	TotalInputSatoshis   uint64
	BurnAmountSatoshis   uint64
	ChangeAmountSatoshis uint64
	FeeSatoshis          uint64
	BitcoinTransaction   *wire.MsgTx

	SerializedTxnHex string
	TxnHashHex       string
	DeSoTxnHashHex   string

	UnsignedHashes []string
} 
type GetTransactionSpendingRequest struct {
	// Transaction hex.
	TransactionHex string `safeForLogging:"true"`
} 
type GetTransactionSpendingResponse struct {
	// Total transaction spending in nanos.
	TotalSpendingNanos uint64 `safeForLogging:"true"`
} 
type GetTxnRequest struct {
	// TxnHash to fetch.
	TxnHashHex string `safeForLogging:"true"`
} 
type GetTxnResponse struct {
	TxnFound bool
} 
type SendDeSoRequest struct {
	SenderPublicKeyBase58Check   string `safeForLogging:"true"`
	RecipientPublicKeyOrUsername string `safeForLogging:"true"`
	AmountNanos                  int64  `safeForLogging:"true"`
	MinFeeRateNanosPerKB         uint64 `safeForLogging:"true"`

	// No need to specify ProfileEntryResponse in each TransactionFee
	TransactionFees []TransactionFee `safeForLogging:"true"`
} 
type SendDeSoResponse struct {
	TotalInputNanos          uint64
	SpendAmountNanos         uint64
	ChangeAmountNanos        uint64
	FeeNanos                 uint64
	TransactionIDBase58Check string
	Transaction              MsgDeSoTxn
	TransactionHex           string
	TxnHashHex               string
} 
type SendDiamondsRequest struct {
	// The public key of the user who is making the transfer.
	SenderPublicKeyBase58Check string `safeForLogging:"true"`

	// The public key or username of the user receiving the transferred creator coin.
	ReceiverPublicKeyBase58Check string `safeForLogging:"true"`

	// The number of diamonds to give the post.
	DiamondPostHashHex string `safeForLogging:"true"`

	// The number of diamonds to give the post.
	DiamondLevel int64 `safeForLogging:"true"`

	MinFeeRateNanosPerKB uint64 `safeForLogging:"true"`

	// No need to specify ProfileEntryResponse in each TransactionFee
	TransactionFees []TransactionFee `safeForLogging:"true"`

	InTutorial bool `safeForLogging:"true"`
} 
type SendDiamondsResponse struct {
	SpendAmountNanos  uint64
	TotalInputNanos   uint64
	ChangeAmountNanos uint64
	FeeNanos          uint64
	Transaction       MsgDeSoTxn
	TransactionHex    string
	TxnHashHex        string
} 
type SubmitPostRequest struct {
	// The public key of the user who made the post or the user
	// who is subsequently is modifying the post.
	UpdaterPublicKeyBase58Check string `safeForLogging:"true"`

	// Optional. Set when modifying a post as opposed to creating one
	// from scratch.
	PostHashHexToModify string `safeForLogging:"true"`

	// The parent post or profile. This is used for comments.
	ParentStakeID string `safeForLogging:"true"`
	// The body of this post.
	BodyObj DeSoBodySchema

	// The PostHashHex of the post being reposted
	RepostedPostHashHex string `safeForLogging:"true"`

	// ExtraData object to hold arbitrary attributes of a post.
	PostExtraData map[string]string `safeForLogging:"true"`

	// When set to true the post will be hidden.
	IsHidden bool `safeForLogging:"true"`

	MinFeeRateNanosPerKB uint64 `safeForLogging:"true"`

	// No need to specify ProfileEntryResponse in each TransactionFee
	TransactionFees []TransactionFee `safeForLogging:"true"`

	InTutorial bool `safeForLogging:"true"`
} 
type SubmitPostResponse struct {
	TstampNanos uint64 `safeForLogging:"true"`
	PostHashHex string `safeForLogging:"true"`

	TotalInputNanos   uint64
	ChangeAmountNanos uint64
	FeeNanos          uint64
	Transaction       MsgDeSoTxn
	TransactionHex    string
} 
type SubmitTransactionRequest struct {
	TransactionHex string `safeForLogging:"true"`
} 
type SubmitTransactionResponse struct {
	Transaction MsgDeSoTxn
	TxnHashHex  string

	// include the PostEntryResponse if a post was submitted
	PostEntryResponse *PostEntryResponse
} 
type TransferCreatorCoinRequest struct {
	// The public key of the user who is making the transfer.
	SenderPublicKeyBase58Check string `safeForLogging:"true"`

	// The public key of the profile for the creator coin that the user is transferring.
	CreatorPublicKeyBase58Check string `safeForLogging:"true"`

	// The public key or username of the user receiving the transferred creator coin.
	ReceiverUsernameOrPublicKeyBase58Check string `safeForLogging:"true"`

	// The amount of creator coins to transfer in nanos.
	CreatorCoinToTransferNanos uint64 `safeForLogging:"true"`

	MinFeeRateNanosPerKB uint64 `safeForLogging:"true"`

	// No need to specify ProfileEntryResponse in each TransactionFee
	TransactionFees []TransactionFee `safeForLogging:"true"`
} 
type TransferCreatorCoinResponse struct {
	SpendAmountNanos  uint64
	TotalInputNanos   uint64
	ChangeAmountNanos uint64
	FeeNanos          uint64
	Transaction       MsgDeSoTxn
	TransactionHex    string
	TxnHashHex        string
} 
type TransferDAOCoinRequest struct {
	// The public key of the user who is making the transfer.
	SenderPublicKeyBase58Check string `safeForLogging:"true"`

	// The public key/Username of the profile for the DAO coin that the user is transferring.
	ProfilePublicKeyBase58CheckOrUsername string `safeForLogging:"true"`

	// The public key/username of the user receiving the transferred creator coin.
	ReceiverPublicKeyBase58CheckOrUsername string `safeForLogging:"true"`

	// The amount of creator coins to transfer in nanos.
	DAOCoinToTransferNanos uint256.Int `safeForLogging:"true"`

	MinFeeRateNanosPerKB uint64 `safeForLogging:"true"`

	// No need to specify ProfileEntryResponse in each TransactionFee
	TransactionFees []TransactionFee `safeForLogging:"true"`
} 
type TransferDAOCoinResponse struct {
	SpendAmountNanos  uint64
	TotalInputNanos   uint64
	ChangeAmountNanos uint64
	FeeNanos          uint64
	Transaction       MsgDeSoTxn
	TransactionHex    string
	TxnHashHex        string
} 
type TransferRestrictionStatusString string 
type UpdateProfileRequest struct {
	// The public key of the user who is trying to update their profile.
	UpdaterPublicKeyBase58Check string `safeForLogging:"true"`

	// This is only set when the user wants to modify a profile
	// that isn't theirs. Otherwise, the UpdaterPublicKeyBase58Check is
	// assumed to own the profile being updated.
	ProfilePublicKeyBase58Check string `safeForLogging:"true"`

	NewUsername    string `safeForLogging:"true"`
	NewDescription string `safeForLogging:"true"`
	// The profile pic string encoded as a link e.g.
	// data:image/png;base64,<data in base64>
	NewProfilePic               string
	NewCreatorBasisPoints       uint64 `safeForLogging:"true"`
	NewStakeMultipleBasisPoints uint64 `safeForLogging:"true"`

	IsHidden bool `safeForLogging:"true"`

	MinFeeRateNanosPerKB uint64 `safeForLogging:"true"`

	// No need to specify ProfileEntryResponse in each TransactionFee
	TransactionFees []TransactionFee `safeForLogging:"true"`
} 
type UpdateProfileResponse struct {
	TotalInputNanos               uint64
	ChangeAmountNanos             uint64
	FeeNanos                      uint64
	Transaction                   MsgDeSoTxn
	TransactionHex                string
	TxnHashHex                    string
	CompProfileCreationTxnHashHex string
}