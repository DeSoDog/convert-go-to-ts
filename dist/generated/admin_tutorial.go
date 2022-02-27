package types
 
type AdminResetTutorialStatusRequest struct {
	PublicKeyBase58Check string
} 
type AdminUpdateTutorialCreatorRequest struct {
	PublicKeyBase58Check string
	IsRemoval            bool
	IsWellKnown          bool
	JWT                  string
} 