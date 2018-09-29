package book

// 新增短评的请求参数
type AddShortCommentBody struct {
	BookId  int    `json:"book_id"`
	Content string `json:"content"`
}

// 书籍搜索的请求参数
type SearchQuery struct {
	Start   int    `json:"start"`
	Count   int    `json:"count"`
	Summary int    `json:"summary"`
	Q       string `json:"q"`
}
