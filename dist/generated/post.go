package types
 
type CommentsPostEntryResponse struct {
	PostEntryResponse *PostEntryResponse
	PosterPublicKeyBytes [] byte
} 
type DiamondSenderResponse struct {
	DiamondSenderProfile *ProfileEntryResponse
	DiamondLevel         int64
} 
type GetDiamondsForPostRequest struct {
	// PostHashHex to fetch.
	PostHashHex                string `safeForLogging:"true"`
	Offset                     uint32 `safeForLogging:"true"`
	Limit                      uint32 `safeForLogging:"true"`
	ReaderPublicKeyBase58Check string `safeForLogging:"true"`
} 
type GetDiamondsForPostResponse struct {
	DiamondSenders []*DiamondSenderResponse
} 
type GetLikesForPostRequest struct {
	// PostHashHex to fetch.
	PostHashHex                string `safeForLogging:"true"`
	Offset                     uint32 `safeForLogging:"true"`
	Limit                      uint32 `safeForLogging:"true"`
	ReaderPublicKeyBase58Check string `safeForLogging:"true"`
} 
type GetLikesForPostResponse struct {
	Likers []*ProfileEntryResponse
} 
type GetPostsDiamondedBySenderForReceiverRequest struct {
	// Public key of the poster who received diamonds from the sender
	ReceiverPublicKeyBase58Check string

	// Username of Receiver
	ReceiverUsername string

	// Public key of the sender who gave diamonds to receiver
	SenderPublicKeyBase58Check string

	// Username of Sender
	SenderUsername string

	// Public key of the reader to get the post entry reader state
	ReaderPublicKeyBase58Check string

	// Start Post Hash Hex
	StartPostHashHex string

	// NumToFetch
	NumToFetch uint64
} 
type GetPostsDiamondedBySenderForReceiverResponse struct {
	// Map of diamond level to a list of post entry responses ordered by timestamp
	DiamondedPosts []*PostEntryResponse

	// Sum of all diamonds sender gave to receiver
	TotalDiamondsGiven uint64

	ReceiverProfileEntryResponse *ProfileEntryResponse

	SenderProfileEntryResponse *ProfileEntryResponse
} 
type GetPostsForPublicKeyRequest struct {
	// Either PublicKeyBase58Check or Username can be set by the client to specify
	// which user we're obtaining posts for
	// If both are specified, PublicKeyBase58Check will supercede
	PublicKeyBase58Check string `safeForLogging:"true"`
	Username             string `safeForLogging:"true"`

	ReaderPublicKeyBase58Check string `safeForLogging:"true"`
	// PostHashHex of the last post from the previous page
	LastPostHashHex string `safeForLogging:"true"`
	// Number of records to fetch
	NumToFetch    uint64 `safeForLogging:"true"`
	MediaRequired bool   `safeForLogging:"true"`
} 
type GetPostsForPublicKeyResponse struct {
	Posts           []*PostEntryResponse `safeForLogging:"true"`
	LastPostHashHex string               `safeForLogging:"true"`
} 
type GetPostsStatelessRequest struct {
	// This is the PostHashHex of the post we want to start our paginated lookup at. We
	// will fetch up to "NumToFetch" posts after it, ordered by time stamp.  If no
	// PostHashHex is provided we will return the most recent posts.
	PostHashHex                string `safeForLogging:"true"`
	ReaderPublicKeyBase58Check string `safeForLogging:"true"`
	OrderBy                    string `safeForLogging:"true"`
	StartTstampSecs            uint64 `safeForLogging:"true"`
	PostContent                string `safeForLogging:"true"`
	NumToFetch                 int    `safeForLogging:"true"`

	// Note: if the GetPostsForFollowFeed option is passed, FetchSubcomments is currently ignored
	// (fetching comments / subcomments for the follow feed is currently unimplemented)
	FetchSubcomments bool `safeForLogging:"true"`

	// This gets posts by people that ReaderPublicKeyBase58Check follows.
	GetPostsForFollowFeed bool `safeForLogging:"true"`

	// This gets posts by people that ReaderPublicKeyBase58Check follows.
	GetPostsForGlobalWhitelist bool `safeForLogging:"true"`

	// This gets posts sorted by deso
	GetPostsByDESO  bool `safeForLogging:"true"`
	GetPostsByClout bool // Deprecated

	// This only gets posts that include media, like photos and videos
	MediaRequired bool `safeForLogging:"true"`

	PostsByDESOMinutesLookback uint64 `safeForLogging:"true"`

	// If set to true, then the posts in the response will contain a boolean about whether they're in the global feed
	AddGlobalFeedBool bool `safeForLogging:"true"`
} 
type GetPostsStatelessResponse struct {
	PostsFound []*PostEntryResponse
} 
type GetQuoteRepostsForPostRequest struct {
	// PostHashHex to fetch.
	PostHashHex                string `safeForLogging:"true"`
	Offset                     uint32 `safeForLogging:"true"`
	Limit                      uint32 `safeForLogging:"true"`
	ReaderPublicKeyBase58Check string `safeForLogging:"true"`
} 
type GetQuoteRepostsForPostResponse struct {
	QuoteReposts  []*PostEntryResponse
	QuoteReclouts []*PostEntryResponse // Deprecated
} 
type GetRepostsForPostRequest struct {
	// PostHashHex to fetch.
	PostHashHex                string `safeForLogging:"true"`
	Offset                     uint32 `safeForLogging:"true"`
	Limit                      uint32 `safeForLogging:"true"`
	ReaderPublicKeyBase58Check string `safeForLogging:"true"`
} 
type GetRepostsForPostResponse struct {
	Reposters  []*ProfileEntryResponse
	Reclouters []*ProfileEntryResponse // Deprecated
} 
type GetSinglePostRequest struct {
	// PostHashHex to fetch.
	PostHashHex                string `safeForLogging:"true"`
	FetchParents               bool   `safeForLogging:"true"`
	CommentOffset              uint32 `safeForLogging:"true"`
	CommentLimit               uint32 `safeForLogging:"true"`
	ReaderPublicKeyBase58Check string `safeForLogging:"true"`
	// How many levels of replies will be retrieved. If unset, will only retrieve the top-level replies.
	ThreadLevelLimit           uint32 `safeForLogging:"true"`
	// How many child replies of a parent comment will be considered when returning a comment thread. Setting this to -1 will include all child replies. This limit does not affect the top-level replies to a post.
	ThreadLeafLimit int32 `safeForLogging:"true"`
	// If the post contains a comment thread where all comments are created by the author, include that thread in the response.
	LoadAuthorThread           bool   `safeForLogging:"true"`

	// If set to true, then the posts in the response will contain a boolean about whether they're in the global feed.
	AddGlobalFeedBool bool `safeForLogging:"true"`
} 
type GetSinglePostResponse struct {
	PostFound *PostEntryResponse
} 
type PostEntryResponse struct {
	PostHashHex                string
	PosterPublicKeyBase58Check string
	ParentStakeID              string
	Body                       string
	ImageURLs                  []string
	VideoURLs                  []string
	RepostedPostEntryResponse  *PostEntryResponse
	CreatorBasisPoints         uint64
	StakeMultipleBasisPoints   uint64
	TimestampNanos             uint64
	IsHidden                   bool
	ConfirmationBlockHeight    uint32
	InMempool                  bool
	// The profile associated with this post.
	ProfileEntryResponse *ProfileEntryResponse
	// The comments associated with this post.
	Comments     []*PostEntryResponse
	LikeCount    uint64
	DiamondCount uint64
	// Information about the reader's state w/regard to this post (e.g. if they liked it).
	PostEntryReaderState PostEntryReaderState
	InGlobalFeed         *bool `json:",omitempty"`
	InHotFeed            *bool `json:",omitempty"`
	// True if this post hash hex is pinned to the global feed.
	IsPinned *bool `json:",omitempty"`
	// PostExtraData stores an arbitrary map of attributes of a PostEntry
	PostExtraData    map[string]string
	CommentCount     uint64
	RepostCount      uint64
	QuoteRepostCount uint64
	// A list of parent posts for this post (ordered: root -> closest parent post).
	ParentPosts []*PostEntryResponse

	// NFT info.
	IsNFT                          bool
	NumNFTCopies                   uint64
	NumNFTCopiesForSale            uint64
	NumNFTCopiesBurned             uint64
	HasUnlockable                  bool
	NFTRoyaltyToCreatorBasisPoints uint64
	NFTRoyaltyToCoinBasisPoints    uint64
	// This map specifies royalties that should go to user's  other than the creator
	AdditionalDESORoyaltiesMap map[string]uint64
	// This map specifies royalties that should be add to creator coins other than the creator's coin.
	AdditionalCoinRoyaltiesMap map[string]uint64

	// Number of diamonds the sender gave this post. Only set when getting diamond posts.
	DiamondsFromSender uint64

	// Score given to this post by the hot feed go routine. Not always populated.
	HotnessScore   uint64
	PostMultiplier float64

	RecloutCount               uint64             // Deprecated
	QuoteRecloutCount          uint64             // Deprecated
	RecloutedPostEntryResponse *PostEntryResponse // Deprecated
}