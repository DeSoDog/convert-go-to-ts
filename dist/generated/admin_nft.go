package types
 
type AdminGetNFTDropRequest struct {
	// "-1" is used to request the next planned drop.
	DropNumber int `safeForLogging:"true"`
} 
type AdminGetNFTDropResponse struct {
	DropEntry *NFTDropEntry
	Posts     []*PostEntryResponse
} 
type AdminUpdateNFTDropRequest struct {
	DropNumber         int    `safeForLogging:"true"`
	DropTstampNanos    int    `safeForLogging:"true"`
	IsActive           bool   `safeForLogging:"true"`
	NFTHashHexToAdd    string `safeForLogging:"true"`
	NFTHashHexToRemove string `safeForLogging:"true"`
} 
type AdminUpdateNFTDropResponse struct {
	DropEntry *NFTDropEntry
	Posts     []*PostEntryResponse
}