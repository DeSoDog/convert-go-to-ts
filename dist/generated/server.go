package types

import (
	"net/http"
	"time"

	"github.com/decred/dcrd/lru"

	"github.com/DataDog/datadog-go/statsd"

	"github.com/deso-protocol/go-deadlock"
)
 
type GetDataRequestInfo struct {
	PeerWhoSentInv *Peer
	TimeRequested  time.Time
} 
type Server struct {
	cmgr          *ConnectionManager
	blockchain    *Blockchain
	mempool       *DeSoMempool
	miner         *DeSoMiner
	blockProducer *DeSoBlockProducer
	eventManager  *EventManager

	// All messages received from peers get sent from the ConnectionManager to the
	// Server through this channel.
	//
	// Generally, the
	// ConnectionManager is responsible for managing the connections to all the peers,
	// but when it receives a message from one of them, it forwards it to the Server
	// on this channel to actually process (acting as a router in that way).
	//
	// In addition to messages from peers, the ConnectionManager will also send control
	// messages to notify the Server e.g. when a Peer connects or disconnects so that
	// the Server can take action appropriately.
	incomingMessages chan *ServerMessage
	// inventoryBeingProcessed keeps track of the inventory (hashes of blocks and
	// transactions) that we've recently processed from peers. It is useful for
	// avoiding situations in which we re-fetch the same data from many peers.
	// For example, if we get the same Block inv message from multiple peers,
	// adding it to this map and checking this map before replying will make it
	// so that we only send a reply to the first peer that sent us the inv, which
	// is more efficient.
	inventoryBeingProcessed lru.Cache
	// hasRequestedSync indicates whether we've bootstrapped our mempool
	// by requesting all mempool transactions from a
	// peer. It's initially false
	// when the server boots up but gets set to true after we make a Mempool
	// request once we're fully synced.
	// The waitGroup is used to manage the cleanup of the Server.
	waitGroup deadlock.WaitGroup

	// During initial block download, we request headers and blocks from a single
	// peer. Note: These fields should only be accessed from the messageHandler thread.
	//
	// TODO: This could be much faster if we were to download blocks in parallel
	// rather than from a single peer but it won't be a problem until later, at which
	// point we can make the optimization.
	SyncPeer *Peer
	// How long we wait on a transaction we're fetching before giving
	// up on it. Note this doesn't apply to blocks because they have their own
	// process for retrying that differs from transactions, which are
	// more best-effort than blocks.
	requestTimeoutSeconds uint32

	// dataLock protects requestedTxns and requestedBlocks
	dataLock deadlock.Mutex

	// requestedTransactions contains hashes of transactions for which we have
	// requested data but have not yet received a response.
	requestedTransactionsMap map[BlockHash]*GetDataRequestInfo

	// addrsToBroadcast is a list of all the addresses we've received from valid addr
	// messages that we intend to broadcast to our peers. It is organized as:
	// <recipient address> -> <list of addresses we received from that recipient>.
	//
	// It is organized in this way so that we can limit the number of addresses we
	// are distributing for a single peer to avoid a DOS attack.
	addrsToBroadcastLock deadlock.RWMutex
	addrsToBroadcastt    map[string][]*SingleAddr

	// When set to true, we disable the ConnectionManager
	disableNetworking bool

	// When set to true, transactions created on this node will be ignored.
	readOnlyMode                 bool
	ignoreInboundPeerInvMessages bool

	// Becomes true after the node has processed its first transaction bundle from
	// any peer. This is useful in a deployment setting because it makes it so that
	// a health check can wait until this value becomes true.
	hasProcessedFirstTransactionBundle bool

	statsdClient *statsd.Client

	Notifier *Notifier
} 
type ServerMessage struct {
	Peer      *Peer
	Msg       DeSoMessage
	ReplyChan chan *ServerReply
} 
type ServerReply struct {
	// GraylistedResponseMap is a map of PKIDs converted to base58-encoded string to a byte slice. This is computed
	// from the GraylistedPKIDMap above and is a JSON-encodable version of that map. This map is only used when
	// responding to requests for this node's graylist. A JSON-encoded response is easier for any language to digest
	// than a gob-encoded one.
	GraylistedResponseMap map[string][]byte
	// GlobalFeedPostHashes is a slice of BlockHashes representing the state of posts on the global feed on this node.
	GlobalFeedPostHashes []BlockHash

	// Cache of Total Supply and Rich List
	TotalSupplyNanos  uint64
	TotalSupplyDESO   float64
	CountKeysWithDESO uint64

	// map of country name to sign up bonus data
	AllCountryLevelSignUpBonuses map[string]CountrySignUpBonusResponse

	// Frequently accessed data from global state
	USDCentsToDESOReserveExchangeRate uint64
	BuyDESOFeeBasisPoints uint64
	JumioUSDCents uint64
	JumioKickbackUSDCents uint64

	// Signals that the frontend server is in a stopped state
	quit chan struct{}
} 
type AccessLevel int
 
type AdminRequest struct {
	JWT            string
	AdminPublicKey string
} 
type AmplitudeEvent struct {
	UserId          string                 `json:"user_id"`
	EventType       string                 `json:"event_type"`
	EventProperties map[string]interface{} `json:"event_properties"`
} 
type AmplitudeUploadRequestBody struct {
	ApiKey string           `json:"api_key"`
	Events []AmplitudeEvent `json:"events"`
} 
type LastTradePriceHistoryItem struct {
	LastTradePrice uint64
	Timestamp      uint64
} 
type Route struct {
	Name        string
	Method      []string
	Pattern     string
	HandlerFunc http.HandlerFunc
	AccessLevel AccessLevel
}