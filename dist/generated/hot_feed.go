package types

import (
	"github.com/deso-protocol/core/lib"
) 
 
type AdminGetHotFeedAlgorithmResponse struct {
	InteractionCap  uint64
	TimeDecayBlocks uint64
} 
type AdminGetHotFeedUserMultiplierRequest struct {
	Username string `safeforlogging:"true"`
} 
type AdminGetHotFeedUserMultiplierResponse struct {
	InteractionMultiplier float64 `safeforlogging:"true"`
	PostsMultiplier       float64 `safeforlogging:"true"`
} 
type AdminUpdateHotFeedAlgorithmRequest struct {
	// Maximum score amount that any individual PKID can contribute to the hot feed score
	// before time decay. Ignored if set to zero.
	InteractionCap int
	// Number of blocks per halving for the hot feed score time decay. Ignored if set to zero.
	TimeDecayBlocks int
} 
type AdminUpdateHotFeedPostMultiplierRequest struct {
	PostHashHex string  `safeforlogging:"true"`
	Multiplier  float64 `safeforlogging:"true"`
} 
type AdminUpdateHotFeedUserMultiplierRequest struct {
	Username              string  `safeforlogging:"true"`
	InteractionMultiplier float64 `safeforlogging:"true"`
	PostsMultiplier       float64 `safeforlogging:"true"`
} 
type HotFeedEntry struct {
	PostHash     BlockHash
	PostHashHex  string
	HotnessScore uint64
} 
type HotFeedInteractionKey struct {
	InteractionPKID     lib.PKID
	InteractionPostHash lib.BlockHash
} 
type HotFeedPKIDMultiplier struct {
	// A multiplier applied to the score that each user interaction adds to a post.
	InteractionMultiplier float64
	// A multiplier applied to all posts from this specific PKID.
	PostsMultiplier float64
} 
type HotFeedPageRequest struct {
	ReaderPublicKeyBase58Check string
	// Since the hot feed is constantly changing, we pass a list of posts that have already
	// been seen in order to send a more accurate next page.
	SeenPosts []string
	// Number of post entry responses to return.
	ResponseLimit int
} 
type HotFeedPageResponse struct {
	HotFeedPage []PostEntryResponse
} 
type HotnessPostInfo struct {
	// How long ago the post was created in number of blocks
	PostBlockAge int
	HotnessScore uint64
}