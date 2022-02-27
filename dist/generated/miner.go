package types

import (
	"github.com/btcsuite/btcd/btcec"
)
 
type DeSoMiner struct {
	PublicKeys    []*btcec.PublicKey
	numThreads    uint32
	BlockProducer *DeSoBlockProducer
	params        *DeSoParams

	stopping int32
} 