package models

type Presence struct {
	State          string `json:"state" binding:"required"`
	LargeImageKey  string `json:"large_image_key" binding:"required"`
	SmallImageKey  string `json:"small_image_key" binding:"required"`
	StartTime      int64  `json:"start_time"`
	EndTime        int64  `json:"end_time"`
	Details        string `json:"details" binding:"required"`
	LargeImageText string `json:"large_image_text" binding:"required"`
	SmallImageText string `json:"small_image_text" binding:"required"`
	PartyId        string `json:"party_id" binding:"required"`
}
