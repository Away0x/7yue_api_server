import { HTTP } from '../utils/http.js'

export class KeywordModel extends HTTP {
  key = 'q'
  maxLength = 10 // 缓存中最大的历史搜索关键字长度

  // 获取热门搜索关键字
  getHot() {
    return this.request({ url: 'book/hot_keyword' }).then(res => res.msg)
  }

  // 获取历史搜索关键字 (缓存中获取)
  getHistory() {
    const words = wx.getStorageSync(this.key)
    if ( ! words) { return [] }
    return words
  }

  addToHistory(keyword) {
    let words = this.getHistory()
    const has = words.includes(keyword)
    // 队列 栈
    if (!has) {
      // 数组末尾 删除，keyword 数组第一位
      const length = words.length
      if (length >= this.maxLength) {
        words.pop()
      }
      words.unshift(keyword)
      wx.setStorageSync(this.key, words)
    }
  }
}