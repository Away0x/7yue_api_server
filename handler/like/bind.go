package like

// 点赞的请求参数
type LikeBody struct {
	ArtId int `json:"art_id"`
	Type int `json:"type"`
}
