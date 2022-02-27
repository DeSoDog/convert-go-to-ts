package types
  
type AcceptNFTBidRequest struct {
	UpdaterPublicKeyBase58Check string `safeForLogging:"true"`
	NFTPostHashHex              string `safeForLogging:"true"`
	SerialNumber                int    `safeForLogging:"true"`
	BidderPublicKeyBase58Check  string `safeForLogging:"true"`
	BidAmountNanos              int    `safeForLogging:"true"`
	EncryptedUnlockableText     string `safeForLogging:"true"`

	MinFeeRateNanosPerKB uint64 `safeForLogging:"true"`

	// No need to specify ProfileEntryResponse in each TransactionFee
	TransactionFees []TransactionFee `safeForLogging:"true"`
} 
type AcceptNFTBidResponse struct {
	BidderPublicKeyBase58Check string `safeForLogging:"true"`
	NFTPostHashHex             string `safeForLogging:"true"`
	SerialNumber               int    `safeForLogging:"true"`
	BidAmountNanos             int    `safeForLogging:"true"`

	TotalInputNanos   uint64
	ChangeAmountNanos uint64
	FeeNanos          uint64
	Transaction       MsgDeSoTxn
	TransactionHex    string
} 
type AcceptNFTTransferRequest struct {
	UpdaterPublicKeyBase58Check string `safeForLogging:"true"`
	NFTPostHashHex              string `safeForLogging:"true"`
	SerialNumber                int    `safeForLogging:"true"`

	MinFeeRateNanosPerKB uint64 `safeForLogging:"true"`

	// No need to specify ProfileEntryResponse in each TransactionFee
	TransactionFees []TransactionFee `safeForLogging:"true"`
} 
type AcceptNFTTransferResponse struct {
	UpdaterPublicKeyBase58Check string `safeForLogging:"true"`
	NFTPostHashHex              string `safeForLogging:"true"`
	SerialNumber                int    `safeForLogging:"true"`

	TotalInputNanos   uint64
	ChangeAmountNanos uint64
	FeeNanos          uint64
	Transaction       MsgDeSoTxn
	TransactionHex    string
} 
type BurnNFTRequest struct {
	UpdaterPublicKeyBase58Check string `safeForLogging:"true"`
	NFTPostHashHex              string `safeForLogging:"true"`
	SerialNumber                int    `safeForLogging:"true"`

	MinFeeRateNanosPerKB uint64 `safeForLogging:"true"`

	// No need to specify ProfileEntryResponse in each TransactionFee
	TransactionFees []TransactionFee `safeForLogging:"true"`
} 
type BurnNFTResponse struct {
	UpdaterPublicKeyBase58Check string `safeForLogging:"true"`
	NFTPostHashHex              string `safeForLogging:"true"`
	SerialNumber                int    `safeForLogging:"true"`

	TotalInputNanos   uint64
	ChangeAmountNanos uint64
	FeeNanos          uint64
	Transaction       MsgDeSoTxn
	TransactionHex    string
} 
type CreateNFTBidRequest struct {
	UpdaterPublicKeyBase58Check string `safeForLogging:"true"`
	NFTPostHashHex              string `safeForLogging:"true"`
	SerialNumber                int    `safeForLogging:"true"`
	BidAmountNanos              int    `safeForLogging:"true"`

	MinFeeRateNanosPerKB uint64 `safeForLogging:"true"`

	// No need to specify ProfileEntryResponse in each TransactionFee
	TransactionFees []TransactionFee `safeForLogging:"true"`
} 
type CreateNFTBidResponse struct {
	UpdaterPublicKeyBase58Check string `safeForLogging:"true"`
	NFTPostHashHex              string `safeForLogging:"true"`
	SerialNumber                int    `safeForLogging:"true"`
	BidAmountNanos              int    `safeForLogging:"true"`

	TotalInputNanos   uint64
	ChangeAmountNanos uint64
	FeeNanos          uint64
	Transaction       MsgDeSoTxn
	TransactionHex    string
} 
type CreateNFTRequest struct {
	UpdaterPublicKeyBase58Check    string            `safeForLogging:"true"`
	NFTPostHashHex                 string            `safeForLogging:"true"`
	NumCopies                      int               `safeForLogging:"true"`
	NFTRoyaltyToCreatorBasisPoints int               `safeForLogging:"true"`
	NFTRoyaltyToCoinBasisPoints    int               `safeForLogging:"true"`
	HasUnlockable                  bool              `safeForLogging:"true"`
	IsForSale                      bool              `safeForLogging:"true"`
	MinBidAmountNanos              int               `safeForLogging:"true"`
	IsBuyNow                       bool              `safeForLogging:"true"`
	BuyNowPriceNanos               uint64            `safeForLogging:"true"`
	AdditionalDESORoyaltiesMap     map[string]uint64 `safeForLogging:"true"`
	AdditionalCoinRoyaltiesMap     map[string]uint64 `safeForLogging:"true"`

	MinFeeRateNanosPerKB uint64 `safeForLogging:"true"`

	// No need to specify ProfileEntryResponse in each TransactionFee
	TransactionFees []TransactionFee `safeForLogging:"true"`
} 
type CreateNFTResponse struct {
	NFTPostHashHex string `safeForLogging:"true"`

	TotalInputNanos   uint64
	ChangeAmountNanos uint64
	FeeNanos          uint64
	Transaction       MsgDeSoTxn
	TransactionHex    string
} 
type GetAcceptedBidHistoryResponse struct {
	AcceptedBidHistoryMap map[uint64][]*NFTBidEntryResponse
} 
type GetNFTBidsForNFTPostRequest struct {
	ReaderPublicKeyBase58Check string
	PostHashHex                string
} 
type GetNFTBidsForNFTPostResponse struct {
	PostEntryResponse *PostEntryResponse
	NFTEntryResponses []*NFTEntryResponse
	BidEntryResponses []*NFTBidEntryResponse
} 
type GetNFTBidsForUserRequest struct {
	UserPublicKeyBase58Check   string `safeForLogging:"true"`
	ReaderPublicKeyBase58Check string `safeForLogging:"true"`
} 
type GetNFTBidsForUserResponse struct {
	NFTBidEntries                              []*NFTBidEntryResponse
	PublicKeyBase58CheckToProfileEntryResponse map[string]*ProfileEntryResponse
	PostHashHexToPostEntryResponse             map[string]*PostEntryResponse
} 
type GetNFTCollectionSummaryRequest struct {
	PostHashHex                string
	ReaderPublicKeyBase58Check string
} 
type GetNFTCollectionSummaryResponse struct {
	NFTCollectionResponse          *NFTCollectionResponse
	SerialNumberToNFTEntryResponse map[uint64]*NFTEntryResponse
} 
type GetNFTEntriesForPostHashRequest struct {
	PostHashHex                string
	ReaderPublicKeyBase58Check string
} 
type GetNFTEntriesForPostHashResponse struct {
	NFTEntryResponses []*NFTEntryResponse
} 
type GetNFTShowcaseRequest struct {
	ReaderPublicKeyBase58Check string `safeForLogging:"true"`
} 
type GetNFTShowcaseResponse struct {
	NFTCollections []*NFTCollectionResponse
} 
type GetNFTsCreatedByPublicKeyRequest struct {
	// Either PublicKeyBase58Check or Username can be set by the client to specify
	// which user we're obtaining NFTs for
	// If both are specified, PublicKeyBase58Check will supercede
	PublicKeyBase58Check string `safeForLogging:"true"`
	Username             string `safeForLogging:"true"`

	ReaderPublicKeyBase58Check string `safeForLogging:"true"`
	// PostHashHex of the last NFT from the previous page
	LastPostHashHex string `safeForLogging:"true"`
	// Number of records to fetch
	NumToFetch uint64 `safeForLogging:"true"`
} 
type GetNFTsCreatedByPublicKeyResponse struct {
	NFTs            []NFTDetails `safeForLogging:"true"`
	LastPostHashHex string       `safeForLogging:"true"`
} 
type GetNFTsForUserRequest struct {
	UserPublicKeyBase58Check   string `safeForLogging:"true"`
	ReaderPublicKeyBase58Check string `safeForLogging:"true"`
	IsForSale                  *bool  `safeForLogging:"true"`
	// Ignored if IsForSale is provided
	IsPending *bool `safeForLogging:"true"`
} 
type GetNFTsForUserResponse struct {
	NFTsMap map[string]*NFTEntryAndPostEntryResponse
} 
type GetNextNFTShowcaseResponse struct {
	NextNFTShowcaseTstamp uint64
} 
type NFTBidEntryResponse struct {
	PublicKeyBase58Check string
	ProfileEntryResponse *ProfileEntryResponse `json:",omitempty"`
	PostHashHex          *string               `json:",omitempty"`
	// likely nil if included in a list of NFTBidEntryResponses for a single NFT
	PostEntryResponse *PostEntryResponse `json:",omitempty"`
	SerialNumber      uint64             `safeForLogging:"true"`
	BidAmountNanos    uint64             `safeForLogging:"true"`

	// What is the highest bid and the lowest bid on this serial number
	HighestBidAmountNanos *uint64 `json:",omitempty"`
	LowestBidAmountNanos  *uint64 `json:",omitempty"`

	// If we fetched the accepted bid history, include the accepted block height.
	AcceptedBlockHeight *uint32 `json:",omitempty"`

	// Current balance of this bidder.
	BidderBalanceNanos uint64
} 
type NFTCollectionResponse struct {
	ProfileEntryResponse    *ProfileEntryResponse `json:",omitempty"`
	PostEntryResponse       *PostEntryResponse    `json:",omitempty"`
	HighestBidAmountNanos   uint64                `safeForLogging:"true"`
	LowestBidAmountNanos    uint64                `safeForLogging:"true"`
	HighestBuyNowPriceNanos *uint64               `safeForLogging:"true"`
	LowestBuyNowPriceNanos  *uint64               `safeForLogging:"true"`
	NumCopiesForSale        uint64                `safeForLogging:"true"`
	NumCopiesBuyNow         uint64                `safeForLogging:"true"`
	AvailableSerialNumbers  []uint64              `safeForLogging:"true"`
} 
type NFTDetails struct {
	NFTEntryResponses     []*NFTEntryResponse
	NFTCollectionResponse *NFTCollectionResponse
} 
type NFTEntryAndPostEntryResponse struct {
	PostEntryResponse *PostEntryResponse
	NFTEntryResponses []*NFTEntryResponse
} 
type NFTEntryResponse struct {
	OwnerPublicKeyBase58Check  string                `safeForLogging:"true"`
	ProfileEntryResponse       *ProfileEntryResponse `json:",omitempty"`
	PostEntryResponse          *PostEntryResponse    `json:",omitempty"`
	SerialNumber               uint64                `safeForLogging:"true"`
	IsForSale                  bool                  `safeForLogging:"true"`
	IsPending                  bool                  `safeForLogging:"true"`
	IsBuyNow                   bool                  `safeForLogging:"true"`
	BuyNowPriceNanos           uint64                `safeForLogging:"true"`
	MinBidAmountNanos          uint64                `safeForLogging:"true"`
	LastAcceptedBidAmountNanos uint64                `safeForLogging:"true"`

	HighestBidAmountNanos uint64 `safeForLogging:"true"`
	LowestBidAmountNanos  uint64 `safeForLogging:"true"`
	// These fields are only populated when the reader is the owner.
	LastOwnerPublicKeyBase58Check *string `json:",omitempty"`
	EncryptedUnlockableText       *string `json:",omitempty"`
} 
type TransferNFTRequest struct {
	SenderPublicKeyBase58Check   string `safeForLogging:"true"`
	ReceiverPublicKeyBase58Check string `safeForLogging:"true"`
	NFTPostHashHex               string `safeForLogging:"true"`
	SerialNumber                 int    `safeForLogging:"true"`
	EncryptedUnlockableText      string `safeForLogging:"true"`

	MinFeeRateNanosPerKB uint64 `safeForLogging:"true"`

	// No need to specify ProfileEntryResponse in each TransactionFee
	TransactionFees []TransactionFee `safeForLogging:"true"`
} 
type TransferNFTResponse struct {
	SenderPublicKeyBase58Check   string `safeForLogging:"true"`
	ReceiverPublicKeyBase58Check string `safeForLogging:"true"`
	NFTPostHashHex               string `safeForLogging:"true"`
	SerialNumber                 int    `safeForLogging:"true"`

	TotalInputNanos   uint64
	ChangeAmountNanos uint64
	FeeNanos          uint64
	Transaction       MsgDeSoTxn
	TransactionHex    string
} 
type UpdateNFTRequest struct {
	UpdaterPublicKeyBase58Check string `safeForLogging:"true"`
	NFTPostHashHex              string `safeForLogging:"true"`
	SerialNumber                int    `safeForLogging:"true"`
	IsForSale                   bool   `safeForLogging:"true"`
	MinBidAmountNanos           int    `safeForLogging:"true"`
	IsBuyNow                    bool   `safeForLogging:"true"`
	BuyNowPriceNanos            uint64 `safeForLogging:"true"`

	MinFeeRateNanosPerKB uint64 `safeForLogging:"true"`

	// No need to specify ProfileEntryResponse in each TransactionFee
	TransactionFees []TransactionFee `safeForLogging:"true"`
} 
type UpdateNFTResponse struct {
	NFTPostHashHex string `safeForLogging:"true"`
	SerialNumber   int    `safeForLogging:"true"`

	TotalInputNanos   uint64
	ChangeAmountNanos uint64
	FeeNanos          uint64
	Transaction       MsgDeSoTxn
	TransactionHex    string
}