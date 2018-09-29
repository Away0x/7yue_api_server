package like

// 点赞的请求参数
type LikeBody struct {
	ArtId int `json:"art_id" binding:"required"`
	Type int `json:"type" binding:"required"`
}
