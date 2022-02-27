package types
  
type APIBalanceRequest struct {
	PublicKeyBase58Check string
	Confirmations        uint32
} 
type APIBalanceResponse struct {
	// Blank if successful. Otherwise, contains a description of the
	// error that occurred.
	Error string
	// The balance of the public key queried in “nanos.” Note
	// there are 1e9 “nanos” per DeSo, so if the balance were “1 DeSo” then
	// this value would be set to 1e9.
	ConfirmedBalanceNanos int64
	// The unconfirmed balance of the public key queried in “nanos.” This field
	// is set to zero if Confirmations is set to a value greater than zero.
	UnconfirmedBalanceNanos int64
	// DeSo uses a UTXO model similar to Bitcoin. As such, querying
	// the balance returns all of the UTXOs for a particular public key for
	// convenience. Note that a UTXO is simply a reference to a particular
	// output index in a previous transaction
	UTXOs []*UTXOEntryResponse
} 
type APIBaseResponse struct {
	// Blank if successful. Otherwise, contains a description of the
	// error that occurred.
	Error string

	// The information contained in the block’s header.
	Header *HeaderResponse

	Transactions []*TransactionResponse
} 
type APIBlockRequest struct {
	// Block height. 0 corresponds to the genesis block. An error will be
	// returned if the height exceeds the tip. This field is ignored if HashHex is
	// set.
	Height int64
	// Hash of the block to return. Height is ignored if this is set.
	HashHex string
	// When set to false, only returns the header of the block requested
	// not the full block. Otherwise, returns the full block.
	FullBlock bool
} 
type APIBlockResponse struct {
	// Blank if successful. Otherwise, contains a description of the
	// error that occurred.
	Error string

	// The information contained in the block’s header.
	Header *HeaderResponse

	Transactions []*TransactionResponse
} 
type APIKeyPairRequest struct {
	// A BIP39 mnemonic and extra text. Mnemonic can be 12 words or
	// 24 words. ExtraText is optional.
	Mnemonic  string
	ExtraText string

	// The index of the public/private key pair to generate
	Index uint32
} 
type APIKeyPairResponse struct {
	// Blank if successful. Otherwise, contains a description of the
	// error that occurred.
	Error string
	// The DeSo public key encoded using base58 check encoding with
	// prefix = [3]byte{0x9, 0x7f, 0x0}
	// This public key can be passed in subsequent API calls to check
	// balance, among other things. All encoded DeSo public keys start
	// with the characters “BC”
	PublicKeyBase58Check string
	// The DeSo public key encoded as a plain hex string. This should
	// match the public key with the corresponding index generated by this tool.
	// This should not be passed to subsequent API calls, it is only provided
	// as a reference, mainly as a sanity-check.
	PublicKeyHex string
	// The DeSo private key encoded using base58 check encoding with
	// prefix = [3]byte{0x50, 0xd5, 0x0}
	// This private key can be passed in subsequent API calls to spend DeSo,
	// among other things. All DeSo private keys start with
	// the characters “bc”
	PrivateKeyBase58Check string
	// The DeSo private key encoded as a plain hex string. Note that
	// this will not directly match what is produced by the tool because the
	// tool shows the private key encoded using Bitcoin’s WIF format rather
	// than as raw hex. To convert this raw hex into Bitcoin’s WIF format you can
	// use this simple Python script. This should not be passed to subsequent
	// API calls, it is only provided as a reference, mainly as a sanity-check.
	PrivateKeyHex string
} 
type APINodeInfoRequest struct {
} 
type APINodeInfoResponse struct {
	// Blank if successful. Otherwise, contains a description of the
	// error that occurred.
	Error string
} 
type APITransactionInfoRequest struct {
	// When set to true, the response simply contains all transactions in the
	// mempool with no filtering.
	IsMempool bool

	// A string that uniquely identifies this transaction. E.g. from a previous
	// call to “transfer-deso”. Ignored when PublicKeyBase58Check is set.
	TransactionIDBase58Check string

	// An DeSo public key encoded using base58 check encoding (starts
	// with “BC”) to get transaction IDs for. When set,
	// TransactionIDBase58Check is ignored.
	PublicKeyBase58Check string

	// Only return transaction IDs
	IDsOnly bool

	// Offset from which a page should be fetched
	LastTransactionIDBase58Check string

	// The last index of a transaction for a public key seen. If less than 0, it means we are not looking at
	// transactions in the database yet.
	LastPublicKeyTransactionIndex int64

	// Number of transactions to be returned
	Limit uint64
} 
type APITransactionInfoResponse struct {
	// Blank if successful. Otherwise, contains a description of the
	// error that occurred.
	Error string

	// The info for all transactions this public key is associated with from oldest
	// to newest.
	Transactions []*TransactionResponse

	// The hash of the last transaction
	LastTransactionIDBase58Check string

	// The last index of a transaction for a public key seen.
	LastPublicKeyTransactionIndex int64

	BalanceNanos uint64
} 
type APITransferDeSoRequest struct {
	// An DeSo private key encoded using base58 check encoding (starts
	// with "bc").
	SenderPrivateKeyBase58Check string
	// An DeSo public key encoded using base58 check encoding (starts
	// with “BC”) that will receive the DeSo being sent. This field is required
	// whether sending using an explicit public/private key pair.
	RecipientPublicKeyBase58Check string
	// The amount of DeSo to send in “nanos.” Note that “1 DeSo” is equal to
	// 1e9 nanos, so to send 1 DeSo, this value would need to be set to 1e9.
	AmountNanos int64
	// The fee rate to use for this transaction. If left unset, a default fee rate
	// will be used. This can be checked using the “DryRun” parameter below.
	MinFeeRateNanosPerKB int64

	// No need to specify ProfileEntryResponse in each TransactionFee
	TransactionFees []TransactionFee `safeForLogging:"true"`

	// When set to true, the transaction is returned in the response but not
	// actually broadcast to the network. Useful for testing.
	DryRun bool
} 
type APITransferDeSoResponse struct {
	// Blank if successful. Otherwise, contains a description of the
	// error that occurred.
	Error string

	// The transaction we assembled.
	Transaction *TransactionResponse

	// Information about the transaction that we compute for
	// convenience.
	TransactionInfo *TransactionInfoResponse
} 
type HeaderResponse struct {
	// The hash of the block that was queried.
	BlockHashHex string
	// Generally set to zero
	Version uint32
	// Hash of the previous block in the chain.
	PrevBlockHashHex string
	// The merkle root of all the transactions contained within the block.
	TransactionMerkleRootHex string
	// The unix timestamp (in seconds) specifying when this block was
	// mined.
	TstampSecs uint64
	// The height of the block this header corresponds to.
	Height uint64

	// The Nonce and ExtraNonce combine to give miners 128 bits of entropy
	Nonce      uint64
	ExtraNonce uint64
} 
type InputResponse struct {
	TransactionIDBase58Check string
	Index                    int64
} 
type OutputResponse struct {
	PublicKeyBase58Check string
	AmountNanos          uint64
} 
type TransactionInfoResponse struct {
	// The sum of the inputs
	TotalInputNanos uint64
	// The amount being sent to the “RecipientPublicKeyBase58Check”
	SpendAmountNanos uint64
	// The amount being returned to the “SenderPublicKeyBase58Check”
	ChangeAmountNanos uint64
	// The total fee and the fee rate (in nanos per KB) that was used for this
	// transaction.
	FeeNanos          uint64
	FeeRateNanosPerKB uint64
	// Will match the public keys passed as params. Note that
	// SenderPublicKeyBase58Check receives the change from this transaction.
	SenderPublicKeyBase58Check    string
	RecipientPublicKeyBase58Check string
} 
type TransactionResponse struct {
	// A string that uniquely identifies this transaction. This is a sha256 hash
	// of the transaction’s data encoded using base58 check encoding.
	TransactionIDBase58Check string
	// The raw hex of the transaction data. This can be fully-constructed from
	// the human-readable portions of this object.
	RawTransactionHex string `json:",omitempty"`
	// The inputs and outputs for this transaction.
	Inputs  []*InputResponse  `json:",omitempty"`
	Outputs []*OutputResponse `json:",omitempty"`
	// The signature of the transaction in hex format.
	SignatureHex string `json:",omitempty"`
	// Will always be “0” for basic transfers
	TransactionType string `json:",omitempty"`
	// TODO: Create a TransactionMeta portion for the response.

	// The hash of the block in which this transaction was mined. If the
	// transaction is unconfirmed, this field will be empty. To look up
	// how many confirmations a transaction has, simply plug this value
	// into the "block" endpoint.
	BlockHashHex string `json:",omitempty"`

	TransactionMetadata TransactionMetadata `json:",omitempty"`

	// The ExtraData added to this transaction
	ExtraData map[string]string `json:",omitempty"`
} 
type UTXOEntryResponse struct {
	// A string that uniquely identifies a previous transaction. This is
	// a sha256 hash of the transaction’s information encoded using
	// base58 check encoding.
	TransactionIDBase58Check string
	// The index within this transaction that corresponds to an output
	// spendable by the passed-in public key.
	Index int64
	// The amount that is spendable by this UTXO in “nanos”.
	AmountNanos uint64
	// The pulic key entitled to spend the amount stored in this UTXO.
	PublicKeyBase58Check string
	// The number of confirmations this UTXO has. Set to zero if the
	// UTXO is unconfirmed.
	Confirmations int64
	// Whether or not this UTXO was a block reward.
	UtxoType string

	BlockHeight int64
}