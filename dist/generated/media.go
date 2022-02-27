package types

import (

	// We import this so that we can decode gifs.
	_ "image/gif"

	// We import this so that we can decode pngs.
	_ "image/png"
)
  
type CFVideoDetailsResponse struct {
	Result   map[string]interface{} `json:"result"`
	Success  bool                   `json:"success"`
	Errors   []interface{}          `json:"errors"`
	Messages []interface{}          `json:"messages"`
} 
type GetFullTikTokURLRequest struct {
	TikTokShortVideoID string
} 
type GetFullTikTokURLResponse struct {
	FullTikTokURL string
} 
type GetVideoStatusResponse struct {
	ReadyToStream bool
	Duration float64
	Dimensions    map[string]interface{}
} 
type UploadImageResponse struct {
	// Location of the image after upload
	ImageURL string
}