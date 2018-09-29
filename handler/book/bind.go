package book

// 新增短评的请求参数
type AddShortCommentBody struct {
	BookId  int    `json:"book_id" binding:"required"`
	Content string `json:"content" binding:"required"`
}
