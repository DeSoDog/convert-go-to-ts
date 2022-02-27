package types
 
type GetTutorialCreatorResponse struct {
	UpAndComingProfileEntryResponses []ProfileEntryResponse
	WellKnownProfileEntryResponses   []ProfileEntryResponse
} 
type GetTutorialCreatorsRequest struct {
	ResponseLimit int
} 
type StartOrSkipTutorialRequest struct {
	PublicKeyBase58Check string
	JWT                  string
	IsSkip               bool
} 
type UpdateTutorialStatusRequest struct {
	PublicKeyBase58Check                string
	TutorialStatus                      TutorialStatus
	CreatorPurchasedInTutorialPublicKey string
	ClearCreatorCoinPurchasedInTutorial bool
	JWT                                 string
}