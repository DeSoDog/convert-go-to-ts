package types
 
type GetJumioStatusForPublicKeyRequest struct {
	JWT                  string
	PublicKeyBase58Check string
} 
type GetJumioStatusForPublicKeyResponse struct {
	JumioFinishedTime uint64
	JumioReturned     bool
	JumioVerified     bool

	BalanceNanos *uint64
} 
type JumioBeginRequest struct {
	PublicKey          string
	ReferralHashBase58 string
	SuccessURL         string
	ErrorURL           string
	JWT                string
} 
type JumioBeginResponse struct {
	URL string
} 
type JumioFlowFinishedRequest struct {
	PublicKey              string
	JumioInternalReference string
	JWT                    string
} 
type JumioIdentityVerification struct {
	Similarity string `json:"similarity"`
	Validity   bool   `json:"validity"`
	Reason     string `json:"reason"`
} 
type JumioInitRequest struct {
	CustomerInternalReference string `json:"customerInternalReference"`
	UserReference             string `json:"userReference"`
	SuccessURL                string `json:"successUrl"`
	ErrorURL                  string `json:"errorUrl"`
} 
type JumioInitResponse struct {
	RedirectURL          string `json:"redirectUrl"`
	TransactionReference string `json:"transactionReference"`
} 
type JumioRejectReason struct {
	RejectReasonCode        string      `json:"rejectReasonCode"`
	RejectReasonDescription string      `json:"rejectReasonDescription"`
	RejectReasonDetails     interface{} `json:"rejectReasonDetails"`
} 
type ResendVerifyEmailRequest struct {
	PublicKey string
	JWT       string
} 
type SendPhoneNumberVerificationTextRequest struct {
	PublicKeyBase58Check string `safeForLogging:"true"`
	PhoneNumber          string
	JWT                  string
} 
type SendPhoneNumberVerificationTextResponse struct {
} 
type SubmitPhoneNumberVerificationCodeRequest struct {
	JWT                  string
	PublicKeyBase58Check string
	PhoneNumber          string
	VerificationCode     string
} 
type SubmitPhoneNumberVerificationCodeResponse struct {
	TxnHashHex string
} 
type VerifyEmailRequest struct {
	PublicKey string
	EmailHash string
}