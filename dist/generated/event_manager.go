package types


type BlockEvent struct {
	Block *MsgDeSoBlock

	// Optional
	UtxoView *UtxoView
	UtxoOps  [][]*UtxoOperation
} 
type BlockEventFunc func(event *BlockEvent) 
type EventManager struct {
	transactionConnectedHandlers []TransactionEventFunc
	blockConnectedHandlers       []BlockEventFunc
	blockDisconnectedHandlers    []BlockEventFunc
	blockAcceptedHandlers        []BlockEventFunc
} 
type TransactionEvent struct {
	Txn     *MsgDeSoTxn
	TxnHash *BlockHash

	// Optional
	UtxoView *UtxoView
	UtxoOps  []*UtxoOperation
} 
type TransactionEventFunc func(event *TransactionEvent)