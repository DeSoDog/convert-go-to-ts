package types

import (
	"github.com/dgraph-io/badger/v3"
)
  
type UtxoView struct {
	// Utxo data
	NumUtxoEntries              uint64
	UtxoKeyToUtxoEntry          map[UtxoKey]*UtxoEntry
	PublicKeyToDeSoBalanceNanos map[PublicKey]uint64

	// BitcoinExchange data
	NanosPurchased     uint64
	USDCentsPerBitcoin uint64
	GlobalParamsEntry  *GlobalParamsEntry
	BitcoinBurnTxIDs   map[BlockHash]bool

	// Forbidden block signature pubkeys
	ForbiddenPubKeyToForbiddenPubKeyEntry map[PkMapKey]*ForbiddenPubKeyEntry

	// Messages data
	MessageKeyToMessageEntry map[MessageKey]*MessageEntry

	// Messaging group entries.
	MessagingGroupKeyToMessagingGroupEntry map[MessagingGroupKey]*MessagingGroupEntry

	// Postgres stores message data slightly differently
	MessageMap map[BlockHash]*PGMessage

	// Follow data
	FollowKeyToFollowEntry map[FollowKey]*FollowEntry

	// NFT data
	NFTKeyToNFTEntry              map[NFTKey]*NFTEntry
	NFTBidKeyToNFTBidEntry        map[NFTBidKey]*NFTBidEntry
	NFTKeyToAcceptedNFTBidHistory map[NFTKey]*[]*NFTBidEntry

	// Diamond data
	DiamondKeyToDiamondEntry map[DiamondKey]*DiamondEntry

	// Like data
	LikeKeyToLikeEntry map[LikeKey]*LikeEntry

	// Repost data
	RepostKeyToRepostEntry map[RepostKey]*RepostEntry

	// Post data
	PostHashToPostEntry map[BlockHash]*PostEntry

	// Profile data
	PublicKeyToPKIDEntry map[PkMapKey]*PKIDEntry
	// The PKIDEntry is only used here to store the public key.
	PKIDToPublicKey               map[PKID]*PKIDEntry
	ProfilePKIDToProfileEntry     map[PKID]*ProfileEntry
	ProfileUsernameToProfileEntry map[UsernameMapKey]*ProfileEntry

	// Creator coin balance entries
	HODLerPKIDCreatorPKIDToBalanceEntry map[BalanceEntryMapKey]*BalanceEntry

	// DAO coin balance entries
	HODLerPKIDCreatorPKIDToDAOCoinBalanceEntry map[BalanceEntryMapKey]*BalanceEntry

	// Derived Key entries. Map key is a combination of owner and derived public keys.
	DerivedKeyToDerivedEntry map[DerivedKeyMapKey]*DerivedKeyEntry

	// The hash of the tip the view is currently referencing. Mainly used
	// for error-checking when doing a bulk operation on the view.
	TipHash *BlockHash

	Handle   *badger.DB
	Postgres *Postgres
	Params   *DeSoParams
}