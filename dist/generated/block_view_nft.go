package types
  
type HelpConnectNFTSoldStruct struct {
	NFTPostHash     *BlockHash
	SerialNumber    uint64
	BidderPKID      *PKID
	BidAmountNanos  uint64
	UnlockableText  []byte
	PrevNFTBidEntry *NFTBidEntry

	// When an NFT owner accepts a bid, they must specify the bidder's UTXO inputs they will lock up
	// as payment for the purchase. This prevents the transaction from accidentally using UTXOs
	// that are used by future transactions.
	BidderInputs []*DeSoInput

	BlockHeight      uint32
	Txn              *MsgDeSoTxn
	TxHash           *BlockHash
	VerifySignatures bool
}