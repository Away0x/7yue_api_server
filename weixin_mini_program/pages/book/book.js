import { BookModel } from '../../models/book.js'
import { random } from '../../utils/common.js'

const bookModel = new BookModel()

Page({
  data: {
    books: [],
    searching: false,
    more: ''
  },

  onLoad(options) {
    bookModel.getHotList().then(res => {
      this.setData({
        books: res
      })
    })
  },

  onSearching(event) {
    this.setData({
      searching: true
    })
  },

  onCancel(event) {
    this.setData({
      searching: false
    })
  },

  // 页面滚动到底部时会自动触发 (自带)
  onReachBottom() {
    this.setData({
      more: random(16) // 使得每次子组件对该属性的监听器都能执行到
    })
  }
  
})