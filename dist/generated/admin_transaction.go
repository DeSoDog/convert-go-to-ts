package types
  
type GetGlobalParamsRequest struct {
} 
type GetGlobalParamsResponse struct {
	// The current exchange rate.
	USDCentsPerBitcoin uint64 `safeForLogging:"true"`

	// The current create profile fee
	CreateProfileFeeNanos uint64 `safeForLogging:"true"`

	// The current minimum fee the network will accept
	MinimumNetworkFeeNanosPerKB uint64 `safeForLogging:"true"`

	// The fee per copy of an NFT minted.
	CreateNFTFeeNanos uint64 `safeForLogging:"true"`

	// The maximum number of copies a single NFT can have.
	MaxCopiesPerNFT uint64 `safeForLogging:"true"`
} 
type SwapIdentityRequest struct {
	// This is currently paramUpdater only
	UpdaterPublicKeyBase58Check string `safeForLogging:"true"`

	// Either a username or a public key works. If it starts with BC and
	// is over the username limit it will be interpreted as a username.
	FromUsernameOrPublicKeyBase58Check string `safeForLogging:"true"`

	// Either a username or a public key works. If it starts with BC and
	//	// is over the username limit it will be interpreted as a username.
	ToUsernameOrPublicKeyBase58Check string `safeForLogging:"true"`

	MinFeeRateNanosPerKB uint64 `safeForLogging:"true"`

	// No need to specify ProfileEntryResponse in each TransactionFee
	TransactionFees []TransactionFee `safeForLogging:"true"`
} 
type SwapIdentityResponse struct {
	TotalInputNanos   uint64
	ChangeAmountNanos uint64
	FeeNanos          uint64
	Transaction       MsgDeSoTxn
	TransactionHex    string
} 
type TestSignTransactionWithDerivedKeyRequest struct {
	// Transaction hex.
	TransactionHex string `safeForLogging:"true"`

	// Derived private key in base58Check.
	DerivedKeySeedHex string `safeForLogging:"false"`
} 
type TestSignTransactionWithDerivedKeyResponse struct {
	// Signed Transaction hex.
	TransactionHex string `safeForLogging:"true"`
} 
type UpdateGlobalParamsRequest struct {
	UpdaterPublicKeyBase58Check string `safeForLogging:"true"`
	// The new exchange rate to set.
	USDCentsPerBitcoin int64 `safeForLogging:"true"`

	// The fee to create a profile.
	CreateProfileFeeNanos int64 `safeForLogging:"true"`

	// The fee per copy of an NFT minted.
	CreateNFTFeeNanos int64 `safeForLogging:"true"`

	// The maximum number of copies a single NFT can have.
	MaxCopiesPerNFT int64 `safeForLogging:"true"`

	// The new minimum fee the network will accept
	MinimumNetworkFeeNanosPerKB int64 `safeForLogging:"true"`

	MinFeeRateNanosPerKB uint64 `safeForLogging:"true"`

	// No need to specify ProfileEntryResponse in each TransactionFee
	TransactionFees []TransactionFee `safeForLogging:"true"`

	// Can be left unset when Signature is false or if the user legitimately
	// doesn't have a password. Can also be left unset if the user has logged
	// in recently as the password will be stored in memory.
	Password string
	// Whether or not we should sign the transaction after constructing it.
	// Setting this flag to false is useful in
	// cases where the caller just wants to construct the transaction
	// to see what the fees will be, for example.
	Sign bool `safeForLogging:"true"`
	// Whether or not we should fully validate the transaction.
	Validate bool `safeForLogging:"true"`
	// Whether or not we should broadcast the transaction after constructing
	// it. This will also validate the transaction if it's set.
	Broadcast bool `safeForLogging:"true"`
} 
type UpdateGlobalParamsResponse struct {
	TotalInputNanos   uint64
	ChangeAmountNanos uint64
	FeeNanos          uint64
	Transaction       MsgDeSoTxn
	TransactionHex    string
}