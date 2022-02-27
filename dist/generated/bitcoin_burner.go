package types

import (
	"github.com/btcsuite/btcd/chaincfg/chainhash"
)
 
type BitcoinUtxo struct {
	TxID           *chainhash.Hash
	Index          int64
	AmountSatoshis int64
} 
type BlockCypherAPIFullAddressResponse struct {
	Address string `json:"address"`
	// Balance data
	ConfirmedBalance   int64 `json:"balance"`
	UnconfirmedBalance int64 `json:"unconfirmed_balance"`
	FinalBalance       int64 `json:"final_balance"`

	// Transaction data
	Txns []*BlockCypherAPITxnResponse `json:"txs"`

	HasMore bool `json:"hasMore"`

	Error string `json:"error"`
} 
type BlockCypherAPIInputResponse struct {
	PrevTxIDHex    string   `json:"prev_hash"`
	Index          int64    `json:"output_index"`
	ScriptHex      string   `json:"script"`
	AmountSatoshis int64    `json:"output_value"`
	Sequence       int64    `json:"sequence"`
	Addresses      []string `json:"addresses"`
	ScriptType     string   `json:"script_type"`
	Age            int64    `json:"age"`
} 
type BlockCypherAPIOutputResponse struct {
	AmountSatoshis int64    `json:"value"`
	ScriptHex      string   `json:"script"`
	Addresses      []string `json:"addresses"`
	ScriptType     string   `json:"script_type"`
	SpentBy        string   `json:"spent_by"`
} 
type BlockCypherAPITxnResponse struct {
	BlockHashHex  string                          `json:"block_hash"`
	BlockHeight   int64                           `json:"block_height"`
	LockTime      int64                           `json:"lock_time"`
	TxIDHex       string                          `json:"hash"`
	Inputs        []*BlockCypherAPIInputResponse  `json:"inputs"`
	Outputs       []*BlockCypherAPIOutputResponse `json:"outputs"`
	Confirmations int64                           `json:"confirmations"`
	DoubleSpend   bool                            `json:"double_spend"`
} 
type BlockchainInfoAPIResponse struct {
	DoubleSpend bool `json:"double_spend"`
} 
type BlockonomicsRBFResponse struct {
	RBF    int64  `json:"rbf"`
	Status string `json:"status"`
}