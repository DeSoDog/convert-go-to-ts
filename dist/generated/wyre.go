package types

import (
	"time"
)
 
type GetWyreWalletOrderForPublicKeyRequest struct {
	PublicKeyBase58Check string
	Username             string

	AdminPublicKey string
} 
type GetWyreWalletOrderForPublicKeyResponse struct {
	WyreWalletOrderMetadataResponses []*WyreWalletOrderMetadataResponse
} 
type WalletOrderQuotationRequest struct {
	SourceAmount   float64
	Country        string
	SourceCurrency string
} 
type WalletOrderReservationRequest struct {
	SourceAmount   float64
	ReferenceId    string
	Country        string
	SourceCurrency string
} 
type WyreTrackOrderResponse struct {
	TransferId  string  `json:"transferId"`
	FeeCurrency string  `json:"feeCurrency"`
	Fee         float64 `json:"fee"`
	Fees        struct {
		BTC float64 `json:"BTC"`
		USD float64 `json:"USD"`
	} `json:"fees"`
	SourceCurrency           string      `json:"sourceCurrency"`
	DestCurrency             string      `json:"destCurrency"`
	SourceAmount             float64     `json:"sourceAmount"`
	DestAmount               float64     `json:"destAmount"`
	DestSrn                  string      `json:"destSrn"`
	From                     string      `json:"from"`
	To                       interface{} `json:"to"`
	Rate                     float64     `json:"rate"`
	CustomId                 interface{} `json:"customId"`
	Status                   interface{} `json:"status"`
	BlockchainNetworkTx      interface{} `json:"blockchainNetworkTx"`
	Message                  interface{} `json:"message"`
	TransferHistoryEntryType string      `json:"transferHistoryEntryType"`
	SuccessTimeline          []struct {
		StatusDetails string `json:"statusDetails"`
		State         string `json:"state"`
		CreatedAt     int64  `json:"createdAt"`
	} `json:"successTimeline"`
	FailedTimeline []interface{} `json:"failedTimeline"`
	FailureReason  interface{}   `json:"failureReason"`
	ReversalReason interface{}   `json:"reversalReason"`
} 
type WyreTransferDetails struct {
	Owner              string      `json:"owner"`
	ReversingSubStatus interface{} `json:"reversingSubStatus"`
	Source             string      `json:"source"`
	PendingSubStatus   interface{} `json:"pendingSubStatus"`
	Status             string      `json:"status"`
	ReversalReason     interface{} `json:"reversalReason"`
	CreatedAt          int64       `json:"createdAt"`
	SourceAmount       float64     `json:"sourceAmount"`
	DestCurrency       string      `json:"destCurrency"`
	SourceCurrency     string      `json:"sourceCurrency"`
	StatusHistories    []struct {
		Id           string      `json:"id"`
		TransferId   string      `json:"transferId"`
		CreatedAt    int64       `json:"createdAt"`
		Type         string      `json:"type"`
		StatusOrder  int         `json:"statusOrder"`
		StatusDetail string      `json:"statusDetail"`
		State        string      `json:"state"`
		FailedState  interface{} `json:"failedState"`
	} `json:"statusHistories"`
	BlockchainTx struct {
		Id            string      `json:"id"`
		NetworkTxId   string      `json:"networkTxId"`
		CreatedAt     int64       `json:"createdAt"`
		Confirmations int         `json:"confirmations"`
		TimeObserved  int64       `json:"timeObserved"`
		BlockTime     int64       `json:"blockTime"`
		Blockhash     string      `json:"blockhash"`
		Amount        float64     `json:"amount"`
		Direction     string      `json:"direction"`
		NetworkFee    float64     `json:"networkFee"`
		Address       string      `json:"address"`
		SourceAddress interface{} `json:"sourceAddress"`
		Currency      string      `json:"currency"`
		TwinTxId      interface{} `json:"twinTxId"`
	} `json:"blockchainTx"`
	ExpiresAt     int64       `json:"expiresAt"`
	CompletedAt   int64       `json:"completedAt"`
	CancelledAt   interface{} `json:"cancelledAt"`
	FailureReason interface{} `json:"failureReason"`
	UpdatedAt     int64       `json:"updatedAt"`
	ExchangeRate  float64     `json:"exchangeRate"`
	DestAmount    float64     `json:"destAmount"`
	Fees          struct {
		BTC int     `json:"BTC"`
		USD float64 `json:"USD"`
	} `json:"fees"`
	TotalFees float64     `json:"totalFees"`
	CustomId  string      `json:"customId"`
	Dest      string      `json:"dest"`
	Message   interface{} `json:"message"`
	Id        string      `json:"id"`
} 
type WyreWalletOrderFullDetails struct {
	Id                      string  `json:"id"`
	CreatedAt               uint64  `json:"createdAt"`
	Owner                   string  `json:"owner"`
	Status                  string  `json:"status"`
	OrderType               string  `json:"orderType"`
	SourceAmount            float64 `json:"sourceAmount"`
	PurchaseAmount          float64 `json:"purchaseAmount"`
	SourceCurrency          string  `json:"sourceCurrency"`
	DestCurrency            string  `json:"destCurrency"`
	TransferId              string  `json:"transferId"`
	Dest                    string  `json:"dest"`
	AuthCodesRequested      bool    `json:"authCodesRequested"`
	ErrorCategory           string  `json:"errorCategory"`
	ErrorCode               string  `json:"errorCode"`
	ErrorMessage            string  `json:"errorMessage"`
	FailureReason           string  `json:"failureReason"`
	AccountId               string  `json:"accountId"`
	PaymentNetworkErrorCode string  `json:"paymentNetworkErrorCode"`
	InternalErrorCode       string  `json:"internalErrorCode"`
} 
type WyreWalletOrderMetadataResponse struct {
	// Last payload received from Wyre webhook
	LatestWyreWalletOrderWebhookPayload WyreWalletOrderWebhookPayload

	// Track Wallet Order response received based on the last payload received from Wyre Webhook
	LatestWyreTrackWalletOrderResponse *WyreTrackOrderResponse

	// Amount of DeSo that was sent for this WyreWalletOrder
	DeSoPurchasedNanos     uint64
	BitCloutPurchasedNanos uint64 // Deprecated

	// BlockHash of the transaction for sending the DeSo
	BasicTransferTxnHash string

	Timestamp *time.Time
} 
type WyreWalletOrderQuotationPayload struct {
	SourceCurrency    string `json:"sourceCurrency"`
	Dest              string `json:"dest"`
	DestCurrency      string `json:"destCurrency"`
	AmountIncludeFees bool   `json:"amountIncludeFees"`
	Country           string `json:"country"`
	SourceAmount      string `json:"sourceAmount"`
	WalletType        string `json:"walletType"`
	AccountId         string `json:"accountId"`
} 
type WyreWalletOrderReservationPayload struct {
	SourceCurrency    string   `json:"sourceCurrency"`
	Dest              string   `json:"dest"`
	DestCurrency      string   `json:"destCurrency"`
	Country           string   `json:"country"`
	Amount            string   `json:"amount"`
	ReferrerAccountId string   `json:"referrerAccountId"`
	LockFields        []string `json:"lockFields"`
	RedirectUrl       string   `json:"redirectUrl"`
	ReferenceId       string   `json:"referenceId"`
} 
type WyreWalletOrderWebhookPayload struct {
	// referenceId holds the public key of the user who made initiated the wallet order
	ReferenceId  string `json:"referenceId"`
	AccountId    string `json:"accountId"`
	OrderId      string `json:"orderId"`
	OrderStatus  string `json:"orderStatus"`
	TransferId   string `json:"transferId"`
	FailedReason string `json:"failedReason"`
}