package types

import (
	"github.com/deso-protocol/core/lib"
) 
 
type AdminAddExemptPublicKey struct {
	// PublicKeyBase58Check is the public key for which we are adding or removing an exemption from node fees.
	PublicKeyBase58Check string
	// IsRemoval is a boolean that when true means we should remove the exemption from a public key, when false means we
	// should add an exemption.
	IsRemoval bool
} 
type AdminGetExemptPublicKeysResponse struct {
	// ExemptPublicKeyMap is a map of PublicKeyBase58Check to ProfileEntryResponse. These public keys do not have to pay
	// node fees.
	ExemptPublicKeyMap map[string]*ProfileEntryResponse
} 
type AdminGetTransactionFeeMapResponse struct {
	// TransactionFeeMap is the current state of Transaction fees on this node.
	TransactionFeeMap map[string][]TransactionFee
} 
type AdminSetAllTransactionFeesRequest struct {
	// NewTransactionFees is a slice of TransactionFees that should be applied to all transaction types.
	// This overwrites all transaction types.
	NewTransactionFees []TransactionFee
} 
type AdminSetAllTransactionFeesResponse struct {
	// TransactionFeeMap is the current state of Transaction fees on this node after the fees defined in
	// AdminSetAllTransactionFeesRequest have been set.
	TransactionFeeMap map[string][]TransactionFee
} 
type AdminSetTransactionFeeForTransactionTypeRequest struct {
	// TransactionType is the type of transaction for which we are setting the fees.
	TransactionType lib.TxnString
	// NewTransactionFees is a slice of TransactionFee structs that tells us who should receive a fee and how much
	// when a transaction of TransactionType is performed.
	NewTransactionFees []TransactionFee
} 
type AdminSetTransactionFeeForTransactionTypeResponse struct {
	// TransactionFeeMap is the current state of Transaction fees on this node after the fees defined in
	// AdminSetTransactionFeeForTransactionTypeRequest have been set.
	TransactionFeeMap map[string][]TransactionFee
} 
type TransactionFee struct {
	// PublicKeyBase58Check is the public key of the user who receives the fee.
	PublicKeyBase58Check string
	// ProfileEntryResponse is only non-nil when TransactionFees are retrieved through admin endpoints.
	// The ProfileEntryResponse is only used to display usernames and avatars in the admin dashboard and thus is
	// excluded in other places to reduce payload sizes and improve performance.
	ProfileEntryResponse *ProfileEntryResponse
	// AmountNanos is the amount PublicKeyBase58Check receives when this fee is incurred.
	AmountNanos uint64
}