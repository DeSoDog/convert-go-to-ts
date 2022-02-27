package types

import (
	"net"
	"time"

	"github.com/decred/dcrd/lru"

	"github.com/btcsuite/btcd/wire"
	"github.com/deso-protocol/go-deadlock"
)
 
type DeSoMessageMeta struct {
	DeSoMessage DeSoMessage
	Inbound     bool
} 
type ExpectedResponse struct {
	TimeExpected time.Time
	MessageType  MsgType
} 
type Peer struct {
	// These stats should be accessed atomically.
	bytesReceived uint64
	bytesSent     uint64
	totalMessages uint64
	lastRecv      int64
	lastSend      int64

	// Stats that should be accessed using the mutex below.
	StatsMtx       deadlock.RWMutex
	timeOffsetSecs int64
	timeConnected  time.Time
	startingHeight uint32
	ID             uint64
	// Ping-related fields.
	LastPingNonce  uint64
	LastPingTime   time.Time
	LastPingMicros int64

	// Connection info.
	cmgr                *ConnectionManager
	conn                net.Conn
	isOutbound          bool
	isPersistent        bool
	stallTimeoutSeconds uint64
	Params              *DeSoParams
	MessageChan         chan *ServerMessage
	// A hack to make it so that we can allow an API endpoint to manually
	// delete a peer.
	PeerManuallyRemovedFromConnectionManager bool

	// In order to complete a version negotiation successfully, the peer must
	// reply to the initial version message we send them with a verack message
	// containing the nonce from that initial version message. This ensures that
	// the peer's IP isn't being spoofed since the only way to actually produce
	// a verack with the appropriate response is to actually own the IP that
	// the peer claims it has. As such, we maintain the version nonce we sent
	// the peer and the version nonce they sent us here.
	//
	// TODO: The way we synchronize the version nonce is currently a bit
	// messy; ideally we could do it without keeping global state.
	versionNonceSent     uint64
	versionNonceReceived uint64

	// A pointer to the Server
	srv *Server

	// Basic state.
	PeerInfoMtx               deadlock.Mutex
	serviceFlags              ServiceFlag
	addrStr                   string
	netAddr                   *wire.NetAddress
	userAgent                 string
	advertisedProtocolVersion uint64
	negotiatedProtocolVersion uint64
	versionNegotiated         bool
	minTxFeeRateNanosPerKB    uint64
	// Messages for which we are expecting a reply within a fixed
	// amount of time. This list is always sorted by ExpectedTime,
	// with the item having the earliest time at the front.
	expectedResponses []*ExpectedResponse

	// The addresses this peer is aware of.
	knownAddressMapLock deadlock.RWMutex
	knownAddressesmap   map[string]bool

	// Output queue for messages that need to be sent to the peer.
	outputQueueChan chan DeSoMessage

	// Set to zero until Disconnect has been called on the Peer. Used to make it
	// so that the logic in Disconnect will only be executed once.
	disconnected int32
	// Signals that the peer is now in the stopped state.
	quit chan interface{}

	// Each Peer is only allowed to have certain number of blocks being sent
	// to them at any gven time. We use
	// this value to enforce that constraint. The reason we need to do this is without
	// it one peer could theoretically clog our Server by issuing many GetBlocks
	// requests that ultimately don't get delivered. This way the number of blocks
	// being sent is limited to a multiple of the number of Peers we have.
	blocksToSendMtx deadlock.Mutex
	blocksToSend    map[BlockHash]bool

	// Inventory stuff.
	// The inventory that we know the peer already has.
	knownInventory lru.Cache

	// Whether the peer is ready to receive INV messages. For a peer that
	// still needs a mempool download, this is false.
	canReceiveInvMessagess bool

	// We process GetTransaction requests in a separate loop. This allows us
	// to ensure that the responses are ordered.
	mtxMessageQueue deadlock.RWMutex
	messagQueue     []*DeSoMessageMeta

	requestedBlocks map[BlockHash]bool
}