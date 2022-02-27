package types

import (
	"sync"
	"time"

	"github.com/deso-protocol/go-deadlock"

	"github.com/btcsuite/btcd/btcec"
)
  
type BlockTemplateStats struct {
	// The number of txns in the block template.
	TxnCount uint32
	// The final txn we attempted to put in the block.
	FailingTxnHash string
	// The reason why the final txn failed to add.
	FailingTxnError string
	// The "Added" time on a transaction changes every time a block is mined so we record
	// the first time added val we are aware of for a specific txn hash here.
	FailingTxnOriginalTimeAdded time.Time
	// The time since the failing txn was added to the mempool.
	FailingTxnMinutesSinceAdded float64
} 
type DeSoBlockProducer struct {
	// The minimum amount of time we wait before trying to produce a new block
	// template. If this value is set low enough then we will produce a block template
	// continuously.
	minBlockUpdateIntervalSeconds uint64
	// The number of templates to cache so that we can accept headers for blocks
	// that are a bit stale.
	maxBlockTemplatesToCache uint64
	// A private key that is used to sign blocks produced by this block producer. Only
	// set if a blockProducerSeed is provided when constructing the BlockProducer.
	blockProducerPrivateKey *btcec.PrivateKey
	// A lock on the block templates produced to avoid concurrency issues.
	mtxRecentBlockTemplatesProduced deadlock.RWMutex
	// The most recent N blocks that we've produced indexed by their hash.
	// Keeping this list allows us to accept a valid header from a miner without
	// requiring them to download/send the whole block.
	recentBlockTemplatesProduced map[BlockHash]*MsgDeSoBlock
	latestBlockTemplateHash      *BlockHash
	currentDifficultyTarget      *BlockHash

	latestBlockTemplateStats *BlockTemplateStats

	mempool *DeSoMempool
	chain   *Blockchain
	params  *DeSoParams

	producerWaitGroup   sync.WaitGroup
	stopProducerChannel chan struct{}

	postgres *Postgres
}