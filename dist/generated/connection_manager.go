package types

import (
	"net"

	"github.com/btcsuite/btcd/addrmgr"
	chainlib "github.com/btcsuite/btcd/blockchain"
	"github.com/decred/dcrd/lru"
	"github.com/deso-protocol/go-deadlock"
)
  
type ConnectionManager struct {
	// Keep a reference to the Server.
	// TODO: I'm pretty sure we can make it so that the ConnectionManager and the Peer
	// doesn't need a reference to the Server object. But for now we keep things lazy.
	srv *Server

	// When --connectips is set, we don't connect to anything from the addrmgr.
	connectIps []string

	// The address manager keeps track of peer addresses we're aware of. When
	// we need to connect to a new outbound peer, it chooses one of the addresses
	// it's aware of at random and provides it to us.
	addrMgr *addrmgr.AddrManager
	// The interfaces we listen on for new incoming connections.
	listeners []net.Listener
	// The parameters we are initialized with.
	params *DeSoParams
	// The target number of outbound peers we want to have.
	targetOutboundPeers uint32
	// The maximum number of inbound peers we allow.
	maxInboundPeers uint32
	// When true, only one connection per IP is allowed. Prevents eclipse attacks
	// among other things.
	limitOneInboundConnectionPerIP bool

	// Keep track of the nonces we've sent in our version messages so
	// we can prevent connections to ourselves.
	sentNonces lru.Cache

	// This section defines the data structures for storing all the
	// peers we're aware of.
	//
	// A count of the number active connections we have for each IP group.
	// We use this to ensure we don't connect to more than one outbound
	// peer from the same IP group. We need a mutex on it because it's used
	// concurrently by many goroutines to figure out if outbound connections
	// should be made to particular addresses.

	mtxOutboundConnIPGroups deadlock.Mutex
	outboundConnIPGroups    map[string]int
	// The peer maps map peer ID to peers for various types of peer connections.
	//
	// A persistent peer is typically one we got through a commandline argument.
	// The reason it's called persistent is because we maintain a connection to
	// it, and retry the connection if it fails.
	mtxPeerMaps     deadlock.RWMutex
	persistentPeers map[uint64]*Peer
	outboundPeers   map[uint64]*Peer
	inboundPeers    map[uint64]*Peer
	// Track the number of outbound peers we have so that this value can
	// be accessed concurrently when deciding whether or not to add more
	// outbound peers.
	numOutboundPeers   uint32
	numInboundPeers    uint32
	numPersistentPeers uint32

	// We keep track of the addresses for the outbound peers so that we can
	// avoid choosing them in the address manager. We need a mutex on this
	// guy because many goroutines will be querying the address manager
	// at once.
	mtxConnectedOutboundAddrs deadlock.RWMutex
	connectedOutboundAddrs    map[string]bool

	// Used to set peer ids. Must be incremented atomically.
	peerIndex uint64

	serverMessageQueue chan *ServerMessage

	// Keeps track of the network time, which is the median of all of our
	// peers' time.
	timeSource chainlib.MedianTimeSource

	// Events that can happen to a peer.
	newPeerChan  chan *Peer
	donePeerChan chan *Peer

	// stallTimeoutSeconds is how long we wait to receive responses from Peers
	// for certain types of messages.
	stallTimeoutSeconds uint64

	minFeeRateNanosPerKB uint64

	// More chans we might want.	modifyRebroadcastInv chan interface{}
	shutdown int32
}