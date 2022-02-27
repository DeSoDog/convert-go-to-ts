package types

import (
	_ "net/http/pprof"
	"testing"

	"github.com/dgraph-io/badger/v3"
) 
 
type TestMeta struct {
	t                      *testing.T
	chain                  *Blockchain
	db                     *badger.DB
	params                 *DeSoParams
	mempool                *DeSoMempool
	miner                  *DeSoMiner
	txnOps                 [][]*UtxoOperation
	txns                   []*MsgDeSoTxn
	expectedSenderBalances []uint64
	savedHeight            uint32
}