package types

import (
	"container/list"
	"math/big"

	chainlib "github.com/btcsuite/btcd/blockchain"
	"github.com/deso-protocol/go-deadlock"
	"github.com/dgraph-io/badger/v3"
)
 
type BlockNode struct {
	// Pointer to a node representing the block's parent.
	Parent *BlockNode

	// The hash computed on this block.
	Hash *BlockHash

	// Height is the position in the block chain.
	Height uint32

	// The difficulty target for this block. Used to compute the next
	// block's difficulty target so it can be validated.
	DifficultyTarget *BlockHash

	// A computation of the total amount of work that has been performed
	// on this chain, including the current node.
	CumWork *big.Int

	// The block header.
	Header *MsgDeSoHeader

	// Status holds the validation state for the block and whether or not
	// it's stored in the database.
	Status BlockStatus
} 
type BlockStatus uint32
 
type Blockchain struct {
	db                              *badger.DB
	postgres                        *Postgres
	timeSource                      chainlib.MedianTimeSource
	trustedBlockProducerPublicKeys  map[PkMapKey]bool
	trustedBlockProducerStartHeight uint64
	params                          *DeSoParams
	eventManager                    *EventManager
	// Returns true once all of the housekeeping in creating the
	// blockchain is complete. This includes setting up the genesis block.
	isInitialized bool

	// Protects most of the fields below this point.
	ChainLock deadlock.RWMutex

	// These should only be accessed after acquiring the ChainLock.
	//
	// An in-memory index of the "tree" of blocks we are currently aware of.
	// This index includes forks and side-chains but does not include unconnectedTxns.
	blockIndex map[BlockHash]*BlockNode
	// An in-memory slice of the blocks on the main chain only. The end of
	// this slice is the best known tip that we have at any given time.
	bestChain    []*BlockNode
	bestChainMap map[BlockHash]*BlockNode

	bestHeaderChain    []*BlockNode
	bestHeaderChainMap map[BlockHash]*BlockNode

	// We keep track of orphan blocks with the following data structures. Orphans
	// are not written to disk and are only cached in memory. Moreover we only keep
	// up to MaxOrphansInMemory of them in order to prevent memory exhaustion.
	orphanList *list.List
} 
type OrphanBlock struct {
	Block *MsgDeSoBlock
	Hash  *BlockHash
} 
type SyncState uint8
