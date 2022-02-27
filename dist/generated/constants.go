package types

import (
	"math/big"
	"time"

	"github.com/btcsuite/btcd/chaincfg"
)
  
type DeSoParams struct {
	// The network type (mainnet, testnet, etc).
	NetworkType NetworkType
	// The current protocol version we're running.
	ProtocolVersion uint64
	// The minimum protocol version we'll allow a peer we connect to
	// to have.
	MinProtocolVersion uint64
	// Used as a "vanity plate" to identify different DeSo
	// clients. Mainly useful in analyzing the network at
	// a meta level, not in the protocol itself.
	UserAgent string
	// The list of DNS seed hosts to use during bootstrapping.
	DNSSeeds []string

	// A list of DNS seed prefixes and suffixes to use during bootstrapping.
	// These prefixes and suffixes will be scanned and all IPs found will be
	// incorporated into the address manager.
	DNSSeedGenerators [][]string

	// The network parameter for Bitcoin messages as defined by the btcd library.
	// Useful for certain function calls we make to this library.
	BitcoinBtcdParams *chaincfg.Params

	// Because we use the Bitcoin header chain only to process exchanges from
	// BTC to DeSo, we don't need to worry about Bitcoin blocks before a certain
	// point, which is specified by this node. This is basically used to make
	// header download more efficient but it's important to note that if for
	// some reason there becomes a different main chain that is stronger than
	// this one, then we will still switch to that one even with this parameter
	// set such as it is.
	BitcoinStartBlockNode *BlockNode

	// The base58Check-encoded Bitcoin address that users must send Bitcoin to in order
	// to purchase DeSo. Note that, unfortunately, simply using an all-zeros or
	// mostly-all-zeros address or public key doesn't work and, in fact, I found that
	// using almost any address other than this one also doesn't work.
	BitcoinBurnAddress string

	// This is a fee in basis points charged on BitcoinExchange transactions that gets
	// paid to the miners. Basically, if a user burned enough Satoshi to create 100 DeSo,
	// and if the BitcoinExchangeFeeBasisPoints was 1%, then 99 DeSo would be allocated to
	// the user's public key while 1 DeSo would be left as a transaction fee to the miner.
	BitcoinExchangeFeeBasisPoints uint64

	// The amount of time to wait for a Bitcoin txn to broadcast throughout the Bitcoin
	// network before checking for double-spends.
	BitcoinDoubleSpendWaitSeconds float64

	// This field allows us to set the amount purchased at genesis to a non-zero
	// value.
	DeSoNanosPurchasedAtGenesis uint64

	// Port used for network communications among full nodes.
	DefaultSocketPort uint16
	// Port used for the limited JSON API that supports light clients.
	DefaultJSONPort uint16

	// The amount of time we wait when connecting to a peer.
	DialTimeout time.Duration
	// The amount of time we wait to receive a version message from a peer.
	VersionNegotiationTimeout time.Duration

	// The genesis block to use as the base of our chain.
	GenesisBlock *MsgDeSoBlock
	// The expected hash of the genesis block. Should align with what one
	// would get from actually hashing the provided genesis block.
	GenesisBlockHashHex string
	// How often we target a single block to be generated.
	TimeBetweenBlocks time.Duration
	// How many blocks between difficulty retargets.
	TimeBetweenDifficultyRetargets time.Duration
	// Block hashes, when interpreted as big-endian big integers, must be
	// values less than or equal to the difficulty
	// target. The difficulty target is expressed below as a big-endian
	// big integer and is adjusted every TargetTimePerBlock
	// order to keep blocks generating at consistent intervals.
	MinDifficultyTargetHex string
	// We will reject chains that have less than this amount of total work,
	// expressed as a hexadecimal big-endian bigint. Useful for preventing
	// disk-fill attacks, among other things.
	MinChainWorkHex string

	// This is used for determining whether we are still in initial block download.
	// If our tip is older than this, we continue with IBD.
	MaxTipAge time.Duration

	// Do not allow the difficulty to change by more than a factor of this
	// variable during each adjustment period.
	MaxDifficultyRetargetFactor int64
	// Amount of time one must wait before a block reward can be spent.
	BlockRewardMaturity time.Duration
	// When shifting from v0 blocks to v1 blocks, we changed the hash function to
	// DeSoHash, which is technically easier. Thus we needed to apply an adjustment
	// factor in order to phase it in.
	V1DifficultyAdjustmentFactor int64

	// The maximum number of seconds in a future a block timestamp is allowed
	// to be before it is rejected.
	MaxTstampOffsetSeconds uint64

	// The maximum number of bytes that can be allocated to transactions in
	// a block.
	MaxBlockSizeBytes uint64

	// It's useful to set the miner maximum block size to a little lower than the
	// maximum block size in certain cases. For example, on initial launch, setting
	// it significantly lower is a good way to avoid getting hit by spam blocks.
	MinerMaxBlockSizeBytes uint64

	// In order to make public keys more human-readable, we convert
	// them to base58. When we do that, we use a prefix that makes
	// the public keys to become more identifiable. For example, all
	// mainnet public keys start with "X" because we do this.
	Base58PrefixPublicKey  [3]byte
	Base58PrefixPrivateKey [3]byte

	// MaxFetchBlocks is the maximum number of blocks that can be fetched from
	// a peer at one time.
	MaxFetchBlocks uint32

	MiningIterationsPerCycle uint64

	// deso
	MaxUsernameLengthBytes        uint64
	MaxUserDescriptionLengthBytes uint64
	MaxProfilePicLengthBytes      uint64
	MaxProfilePicDimensions       uint64
	MaxPrivateMessageLengthBytes  uint64

	StakeFeeBasisPoints         uint64
	MaxPostBodyLengthBytes      uint64
	MaxPostSubLengthBytes       uint64
	MaxStakeMultipleBasisPoints uint64
	MaxCreatorBasisPoints       uint64
	MaxNFTRoyaltyBasisPoints    uint64
	ParamUpdaterPublicKeys      map[PkMapKey]bool

	// A list of transactions to apply when initializing the chain. Useful in
	// cases where we want to hard fork or reboot the chain with specific
	// transactions applied.
	SeedTxns []string

	// A list of balances to initialize the blockchain with. This is useful for
	// testing and useful in the event that the devs need to hard fork the chain.
	SeedBalances []*DeSoOutput

	// This is a small fee charged on creator coin transactions. It helps
	// prevent issues related to floating point calculations.
	CreatorCoinTradeFeeBasisPoints uint64
	// These two params define the "curve" that we use when someone buys/sells
	// creator coins. Effectively, this curve amounts to a polynomial of the form:
	// - currentCreatorCoinPrice ~= slope * currentCreatorCoinSupply^(1/reserveRatio-1)
	// Buys and sells effectively take the integral of the curve in opposite directions.
	//
	// To better understand where this curve comes from and how it works, check out
	// the following links. They are all well written so don't be intimidated/afraid to
	// dig in and read them:
	// - Primer on bonding curves: https://medium.com/@simondlr/tokens-2-0-curved-token-bonding-in-curation-markets-1764a2e0bee5
	// - The Uniswap v2 white paper: https://whitepaper.io/document/600/uniswap-whitepaper
	// - The Bancor white paper: https://whitepaper.io/document/52/bancor-whitepaper
	// - Article relating Bancor curves to polynomial curves: https://medium.com/@aventus/token-bonding-curves-547f3a04914
	// - Derivations of the Bancor supply increase/decrease formulas: https://blog.relevant.community/bonding-curves-in-depth-intuition-parametrization-d3905a681e0a
	// - Implementations of Bancor equations in Solidity with code: https://yos.io/2018/11/10/bonding-curves/
	// - Bancor is flawed blog post discussing Bancor edge cases: https://hackingdistributed.com/2017/06/19/bancor-is-flawed/
	// - A mathematica equation sheet with tests that walks through all the
	//   equations. You will need to copy this into a Mathematica notebook to
	//   run it: https://pastebin.com/raw/M4a1femY
	CreatorCoinSlope        *big.Float
	CreatorCoinReserveRatio *big.Float

	// CreatorCoinAutoSellThresholdNanos defines two things. The first is the minimum amount
	// of creator coins a user must purchase in order for a transaction to be valid. Secondly
	// it defines the point at which a sell operation will auto liquidate all remaining holdings.
	// For example if I hold 1000 nanos of creator coins and sell x nanos such that
	// 1000 - x < CreatorCoinAutoSellThresholdNanos, we auto liquidate the remaining holdings.
	// It does this to prevent issues with floating point rounding that can arise.
	// This value should be chosen such that the chain is resistant to "phantom nanos." Phantom nanos
	// are tiny amounts of CreatorCoinsInCirculation/DeSoLocked which can cause
	// the effective reserve ratio to deviate from the expected reserve ratio of the bancor curve.
	// A higher CreatorCoinAutoSellThresholdNanos makes it prohibitively expensive for someone to
	// attack the bancor curve to any meaningful measure.
	CreatorCoinAutoSellThresholdNanos uint64

	ForkHeights ForkHeights
} 
type ForkHeights struct {
	// Global Block Heights:
	// The block height at which various forks occurred including an
	// explanation as to why they're necessary.

	// The most deflationary event in DeSo history has yet to come...
	DeflationBombBlockHeight uint64

	// SalomonFixBlockHeight defines a block height where the protocol implements
	// two changes:
	// 	(1) The protocol now prints founder reward for all buy transactions instead
	//		of just when creators reach a new all time high.
	//		This was decided in order to provide lasting incentive for creators
	//		to utilize the protocol.
	//	(2) A fix was created to deal with a bug accidentally triggered by @salomon.
	//		After a series of buys and sells @salomon was left with a single creator coin
	//		nano in circulation and a single DeSo nano locked. This caused a detach
	//		between @salomon's bonding curve and others on the protocol. As more buys and sells
	//		continued, @salomon's bonding curve continued to detach further and further from its peers.
	// 		At its core, @salomon had too few creator coins in circulation. This fix introduces
	//		this missing supply back into circulation as well as prevented detached Bancor bonding
	//		curves from coming into existence.
	//		^ It was later decided to leave Salomon's coin circulation alone. A fix was introduced
	//		to prevent similar cases from occurring again, but @salomon is left alone.
	SalomonFixBlockHeight uint32

	// DeSoFounderRewardBlockHeight defines a block height where the protocol switches from
	// paying the founder reward in the founder's own creator coin to paying in DeSo instead.
	DeSoFounderRewardBlockHeight uint32

	// BuyCreatorCoinAfterDeletedBalanceEntryFixBlockHeight defines a block height after which the protocol will create
	// a new BalanceEntry when a user purchases a Creator Coin and their current BalanceEntry is deleted.
	// The situation in which a BalanceEntry reaches a deleted state occurs when a user transfers all their holdings
	// of a certain creator to another public key and subsequently purchases that same creator within the same block.
	// This resolves a bug in which users would purchase creator coins after transferring all holdings within the same
	// block and then the creator coins would be added to a deleted balance.  When the Balance Entries are flushed to
	// the database, the user would lose the creator coins they purchased.
	BuyCreatorCoinAfterDeletedBalanceEntryFixBlockHeight uint32

	// ParamUpdaterProfileUpdateFixBlockHeight defines a block height after which the protocol uses the update profile
	// txMeta's ProfilePublicKey when the Param Updater is creating a profile for ProfilePublicKey.
	ParamUpdaterProfileUpdateFixBlockHeight uint32

	// UpdateProfileFixBlockHeight defines the height at which a patch was added to prevent user from
	// updating the profile entry for arbitrary public keys that do not have existing profile entries.
	UpdateProfileFixBlockHeight uint32

	// BrokenNFTBidsFixBlockHeight defines the height at which the deso balance index takes effect
	// for accepting NFT bids.  This is used to fix a fork that was created by nodes running with a corrupted
	// deso balance index, allowing bids to be submitted that were greater than the user's deso balance.
	BrokenNFTBidsFixBlockHeight uint32

	// DeSoDiamondsBlockHeight defines the height at which diamonds will be given in DESO
	// rather than in creator coin.
	// Triggers: 3pm PT on 8/16/2021
	DeSoDiamondsBlockHeight uint32

	// NFTTransfersBlockHeight defines the height at which NFT transfer txns, accept NFT
	// transfer txns, NFT burn txns, and AuthorizeDerivedKey txns will be accepted.
	// Triggers: 12PM PT on 9/15/2021
	NFTTransferOrBurnAndDerivedKeysBlockHeight uint32

	// DeSoV3MessagesBlockHeight defines the height at which messaging key and messsage party
	// entries will be accepted by consensus.
	DeSoV3MessagesBlockHeight uint32

	// BuyNowAndNFTSplitsBlockHeight defines the height at which NFTs can be sold at a fixed price instead of an
	// auction style and allows splitting of NFT royalties to user's other than the post's creator.
	BuyNowAndNFTSplitsBlockHeight uint32

	// DAOCoinBlockHeight defines the height at which DAO Coin and DAO Coin Transfer
	// transactions will be accepted.
	DAOCoinBlockHeight uint32
} 
type NetworkType uint64
