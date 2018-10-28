import { BookModel } from '../../models/book.js'
import { LikeModel } from '../../models/like.js'

const bookModel = new BookModel()
const likeModel = new LikeModel()

Page({
  data: {
    comments: [],
    book: null,
    likeStatus: false,
    likeCount: 0,
    posting: false
  },

  onLoad(options) {
    wx.showLoading()
    // 获取页面 query 参数
    const bid = options.bid
    const isbn = options.isbn
  
    const detail = bookModel.getDetail(isbn)
    const comments = bookModel.getComments(bid)
    const likeStatus = bookModel.getLikeStatus(bid)

    Promise.all([detail, comments, likeStatus])
      .then(res => {
        this.setData({
          book: res[0],
          comments: res[1].comments,
          likeStatus: res[2].like_status,
          likeCount: res[2].fav_nums
        })
        wx.hideLoading()
      })
  },
  // 点赞
  onLike(event) {
    const like_or_cancel = event.detail.behavior
    likeModel.like(like_or_cancel, this.data.book.id, 400)
  },
  // 新增短评
  onFakePost(event) {
    this.setData({ posting: true })
  },
  onCancel(event) {
    this.setData({ posting: false })
  },
  onPost(event) {
    // event.detail.value 是 input 提交的
    const comment = event.detail.text || event.detail.value

    if (!comment) { return }

    if (comment.length > 12) {
      wx.showToast({ title: '短评最多12个字', icon: 'none' })
      return
    }

    bookModel.postComment(this.data.book.id, comment)
      .then(res => {
        wx.showToast({ title: '+ 1', icon: "none" })

        this.data.comments.unshift({ content: comment, nums: 1 })
        this.setData({ comments: this.data.comments, posting: false })
      })
  },
})