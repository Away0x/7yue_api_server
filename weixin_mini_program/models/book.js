import { HTTP } from '../utils/http.js'

export class BookModel extends HTTP {
  // 热门书籍列表
  getHotList() {
    return this.request({ url: 'book/hot_list'}).then(res => res.msg)
  }

  // 获取喜欢书籍数量
  getMyBookCount() {
    return this.request({ url: 'book/favor_count' }).then(res => res.msg)
  }

  // 获取书籍详细信息
  getDetail(isbn) {
    return this.request({ url: `book/detail/${isbn}` }).then(res => res.msg)
  }

  // 获取书籍点赞情况
  getLikeStatus(bid) {
    return this.request({ url: `book/favor/${bid}` }).then(res => res.msg)
  }

  // 获取书籍短评
  getComments(bid) {
    return this.request({ url: `book/short_comment/${bid}` }).then(res => res.msg)
  }

  // 添加书籍短评
  postComment(bid, comment) {
    return this.request({
      url: 'book/add/short_comment',
      method: 'POST',
      data: { book_id: Number(bid), content: comment }
    })
  }

  // 搜索书籍
  search(start, q) {
    return this.request({
      url: 'book/search?summary=1',
      data: { q: q, start: start }
    }).then(data => data.msg)
  }
}