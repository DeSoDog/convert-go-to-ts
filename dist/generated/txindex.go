package types

import (
	"sync"

	"github.com/deso-protocol/go-deadlock"
)
 
type TXIndex struct {
	// TXIndexLock protects the transaction index.
	TXIndexLock deadlock.RWMutex

	// The txindex has it s own separate Blockchain object. This allows us to
	// capture more metadata when collecting transactions without interfering
	// with the goings-on of the main chain.
	TXIndexChain *Blockchain

	// Core objects from Server
	CoreChain *Blockchain

	// Core params object
	Params *DeSoParams

	// Update wait group
	updateWaitGroup sync.WaitGroup

	// Shutdown channel
	stopUpdateChannel chan struct{}
}