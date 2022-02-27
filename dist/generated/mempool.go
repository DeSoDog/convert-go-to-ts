package types

import (
	"time"

	"github.com/deso-protocol/go-deadlock"
)
  
type DeSoMempool struct {
	// Stops the mempool's services.
	quit chan struct{}

	// A reference to a blockchain object that can be used to validate transactions before
	// adding them to the pool.
	bc *Blockchain

	// Transactions with a feerate below this threshold are outright rejected.
	minFeeRateNanosPerKB uint64

	// rateLimitFeeRateNanosPerKB defines the minimum transaction feerate in "nanos per KB"
	// before a transaction is considered for rate-limiting. Note that even if a
	// transaction with a feerate below this threshold is not rate-limited, it must
	// still have a high enough feerate to be considered as part of the mempool.
	rateLimitFeeRateNanosPerKB uint64

	mtx deadlock.RWMutex

	// poolMap contains all of the transactions that have been validated by the pool.
	// Transactions in poolMap should be directly consumable by a miner and formed into
	// a block by taking them in order of when they were Added.
	poolMap map[BlockHash]*MempoolTx
	// txFeeMinHeap organizes transactions stored in poolMap by their FeePerKB. It is used
	// in order to prevent the pool from exhausing memory due to having to store too
	// many low-fee transactions.
	txFeeMinheap MempoolTxFeeMinHeap
	// totalTxSizeBytes is the total size of all of the transactions stored in poolMap. We
	// use it to determine when the pool is nearing memory-exhaustion so we can start
	// evicting transactions.
	totalTxSizeBytes uint64
	// Stores the inputs for every transaction stored in poolMap. Used to quickly check
	// if a transaction is double-spending.
	outpoints map[UtxoKey]*MsgDeSoTxn
	// Unconnected contains transactions whose inputs reference UTXOs that are not yet
	// present in either our UTXO database or the transactions stored in pool.
	unconnectedTxns map[BlockHash]*UnconnectedTx
	// Organizes unconnectedTxns by their UTXOs. Used when adding a transaction to determine
	// which unconnectedTxns are no longer missing parents.
	unconnectedTxnsByPrev map[UtxoKey]map[BlockHash]*MsgDeSoTxn
	// An exponentially-decayed accumulator of "low-fee" transactions we've relayed.
	// This is used to prevent someone from flooding the network with low-fee
	// transactions.
	lowFeeTxSizeAccumulator float64
	// The UNIX time (in seconds) when the last "low-fee" transaction was relayed.
	lastLowFeeTxUnixTime int64

	// pubKeyToTxnMap stores a mapping from the public key of outputs added
	// to the mempool to the corresponding transaction that resulted in their
	// addition. It is useful for figuring out how much DeSo a particular public
	// key has available to spend.
	pubKeyToTxnMap map[PkMapKey]map[BlockHash]*MempoolTx

	// The next time the unconnectTxn pool will be scanned for expired unconnectedTxns.
	nextExpireScan time.Time

	// Optional. When set, we use the BlockCypher API to detect double-spends.
	blockCypherAPIKey string

	// These two views are used to check whether a transaction is valid before
	// adding it to the mempool. This is done by applying the transaction to the
	// backup view, and then restoring the backup view if there's an error. In
	// the future, if we can figure out an easy way to rollback bad transactions
	// on a single view, then we won't need the second view anymore.
	backupUniversalUtxoView  *UtxoView
	universalUtxoView        *UtxoView
	universalTransactionList []*MempoolTx

	// When set, transactions are initially read from this dir and dumped
	// to this dir.
	mempoolDir string

	// Whether or not we should be computing readOnlyUtxoViews.
	generateReadOnlyUtxoView bool
	// A view that contains a *near* up-to-date snapshot of the mempool. It is
	// updated periodically after N transactions OR after M  seconds, whichever
	// comes first. It's useful because it can be obtained without acquiring a
	// lock on the mempool.
	//
	// This field isn't reset with ResetPool. It requires an explicit call to
	// UpdateReadOnlyView.
	readOnlyUtxoView *UtxoView
	// Keep a list of all transactions in the mempool. This is useful for dumping
	// to the database periodically.
	readOnlyUniversalTransactionList []*MempoolTx
	readOnlyUniversalTransactionMap  map[BlockHash]*MempoolTx
	readOnlyOutpoints                map[UtxoKey]*MsgDeSoTxn
	// Every time the readOnlyUtxoView is updated, this is incremented. It can
	// be used by obtainers of the readOnlyUtxoView to wait until a particular
	// transaction has been run.
	//
	// This field isn't reset with ResetPool. It requires an explicit call to
	// UpdateReadOnlyView.
	readOnlyUtxoViewSequenceNumber int64
	// The total number of times we've called processTransaction. Used to
	// determine whether we should update the readOnlyUtxoView.
	//
	// This field isn't reset with ResetPool. It requires an explicit call to
	// UpdateReadOnlyView.
	totalProcessTransactionCalls int64

	// We pass a copy of the data dir flag to the tx pool so that we can instantiate
	// temp badger db instances and dump mempool txns to them.
	dataDir string
} 
type MempoolTxFeeMinHeap []*MempoolTx

type MempoolTx struct {
	Tx *MsgDeSoTxn

	// TxMeta is the transaction metadata
	TxMeta *TransactionMetadata

	// Hash is a hash of the transaction so we don't have to recompute
	// it all the time.
	Hash *BlockHash

	// TxSizeBytes is the cached size of the transaction.
	TxSizeBytes uint64

	// The time when the txn was added to the pool
	Added time.Time

	// The block height when the txn was added to the pool. It's generally set
	// to tip+1.
	Height uint32

	// The total fee the txn pays. Cached for efficiency reasons.
	Fee uint64

	// The fee rate of the transaction in nanos per KB.
	FeePerKB uint64

	// index is used by the heap logic to allow for modification in-place.
	index int
} 
type SummaryStats struct {
	// Number of transactions of this type in the mempool.
	Count uint32

	// Number of bytes for transactions of this type in the mempool.
	TotalBytes uint64
} 
type UnconnectedTx struct {
	tx *MsgDeSoTxn
	// The ID of the Peer who initially sent the unconnected txn. Useful for
	// removing unconnected transactions when a Peer disconnects.
	peerID     uint64
	expiration time.Time
}