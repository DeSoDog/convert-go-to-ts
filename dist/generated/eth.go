package types
  
type AdminProcessETHTxRequest struct {
	ETHTxHash string
} 
type AdminProcessETHTxResponse struct {
	DESOTxHash string
} 
type ETHTx struct {
	Nonce   string `json:"nonce"`
	Value   string `json:"value"`
	ChainId string `json:"chainId"`
	To      string `json:"to"`
	R       string `json:"r"`
	S       string `json:"s"`
} 
type ETHTxLog struct {
	PublicKey  []byte
	DESOTxHash string
} 
type InfuraRequest struct {
	JSONRPC string        `json:"jsonrpc"`
	Method  string        `json:"method"`
	Params  []interface{} `json:"params"`
	Id      uint64        `json:"id"`
} 
type InfuraResponse struct {
	Id      uint64      `json:"id"`
	JSONRPC string      `json:"jsonrpc"`
	Result  interface{} `json:"result"`
	Error   struct {
		Code    float64 `json:"code"`
		Message string  `json:"message"`
	} `json:"error"`
} 
type InfuraTx struct {
	BlockHash        *string `json:"blockHash"`
	BlockNumber      *string `json:"blockNumber"`
	From             string  `json:"from"`
	Gas              string  `json:"gas"`
	GasPrice         string  `json:"gasPrice"`
	Hash             string  `json:"hash"`
	Input            string  `json:"input"`
	Nonce            string  `json:"nonce"`
	To               *string `json:"to"`
	TransactionIndex *string `json:"transactionIndex"`
	Value            string  `json:"value"`
	V                string  `json:"v"`
	R                string  `json:"r"`
	S                string  `json:"s"`
} 
type QueryETHRPCRequest struct {
	Method               string
	Params               []interface{}
} 
type SubmitETHTxRequest struct {
	PublicKeyBase58Check string
	Tx                   ETHTx
	TxBytes              string
	ToSign               []string
	SignedHashes         []string
} 
type SubmitETHTxResponse struct {
	DESOTxHash string
}