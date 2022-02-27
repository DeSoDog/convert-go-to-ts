package types
  
type AdminGetMempoolStatsResponse struct {
	TransactionSummaryStats map[string]SummaryStats
} 
type NodeControlRequest struct {
	// An address in <IP>:<Port> format.
	Address string `safeForLogging:"true"`

	// A comma-separated list of miner public keys to use.
	MinerPublicKeys string `safeForLogging:"true"`

	// The type of operation to perform on the node.
	OperationType string `safeForLogging:"true"`

	JWT            string
	AdminPublicKey string
} 
type NodeControlResponse struct {
	// The current status the DeSo node is at in terms of syncing the DeSo
	// chain.
	DeSoStatus *NodeStatusResponse

	DeSoOutboundPeers    []*PeerResponse
	DeSoInboundPeers     []*PeerResponse
	DeSoUnconnectedPeers []*PeerResponse

	MinerPublicKeys []string
} 
type NodeStatusResponse struct {
	// A summary of what the node is currently doing.
	State string `safeForLogging:"true"`

	// We generally track the latest header we have and the latest block we have
	// separately since headers-first synchronization can cause the latest header
	// to diverge slightly from the latest block.
	LatestHeaderHeight     uint32 `safeForLogging:"true"`
	LatestHeaderHash       string `safeForLogging:"true"`
	LatestHeaderTstampSecs uint32 `safeForLogging:"true"`

	LatestBlockHeight     uint32 `safeForLogging:"true"`
	LatestBlockHash       string `safeForLogging:"true"`
	LatestBlockTstampSecs uint32 `safeForLogging:"true"`
	LatestTxIndexHeight   uint32 `safeForLogging:"true"`

	// This is non-zero unless the main header chain is fully current. It can be
	// an estimate in cases where we don't know exactly what the tstamp of the
	// current main chain is.
	HeadersRemaining uint32 `safeForLogging:"true"`
	// This is non-zero unless the main header chain is fully current and all
	// the corresponding blocks have been downloaded.
	BlocksRemaining uint32 `safeForLogging:"true"`
} 
type PeerResponse struct {
	IP           string
	ProtocolPort uint16
	IsSyncPeer   bool
}