package types
 
type AdminPinPostRequest struct {
	// The post hash of the post to pin or unpin from the global feed
	PostHashHex string `safeForLogging:"true"`
	// If true, remove the given post hash hex from the list of pinned posts
	UnpinPost bool `safeForLogging:"true"`
} 
type AdminRemoveNilPostsRequest struct {
	// Number of posts to try to fetch from global state, starting from the most recent post
	// added to the global feed.
	NumPostsToSearch int `safeForLogging:"true"`
} 
type AdminUpdateGlobalFeedRequest struct {
	// The post hash of the post to add or remove from the global feed.
	PostHashHex string `safeForLogging:"true"`
	// If true, remove the given post hash hex from the global feed.
	RemoveFromGlobalFeed bool `safeForLogging:"true"`
}