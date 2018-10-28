import { HTTP } from '../utils/http.js'

export class ClassicModel extends HTTP {
  // 获取自己点赞过的期刊
  getMyFavor() {
    return this.request({url: 'classic/favor'}).then(res => res.msg)
  }

  // 获取指定 id 的期刊
  getById(cid, type) {
    return this.request({ url: `classic/detail/${type}/${cid}`}).then(res => res.msg)
  }

  // 获取最新一期
  getLatest() {
    return this.request({url: 'classic/latest'}).then((res) => {
      // 存储最新一期的 index
      this._setLatestIndex(res.msg.index)
      // 写入缓存，以便下次获取
      wx.setStorageSync(this._getKey(res.msg.index), res.msg)
      return res.msg
    })
  }

  // 获取前一期或者后一期
  // 缓存中有则从缓存中获取，否则发送请求获取
  getClassic(index, nextOrPrevious) {
    return new Promise((resolve, reject) => {
      if (nextOrPrevious !== 'next' && nextOrPrevious !== 'previous') {
        throw new Error('nextOrPrevious error: ' + nextOrPrevious)
      }

      // 1. 尝试从缓存中获取
      const key = this._getKey(index + (nextOrPrevious == 'next' ? 1 : -1))
      const cache_classic = wx.getStorageSync(key)

      if (cache_classic) {
        resolve(cache_classic)
      }

      // 2. 缓存中没有，则发送请求获取数据
      this.request({ url: `classic/${nextOrPrevious}/${index}`}).then(res => {
        // 写入缓存，以便下次获取
        wx.setStorageSync(this._getKey(res.msg.index), res.msg)
        resolve(res.msg)
      })
    })
  }

  // 是否为第一个
  isFirst(index) {
    return index === 1
  }

  // 是否为最新的一个
  isLatest(index) {
    const latestIndex = this._getLatestIndex()

    return latestIndex === index
  }

  // 存储最新一期的 index 到 storage 中
  _setLatestIndex(index) {
    wx.setStorageSync('laterst', index)
  }

  // 获取 storage 中存储的最新一期的 index
  _getLatestIndex() {
    return wx.getStorageSync('laterst')
  }

  // 生成缓存 key
  _getKey(index) {
    const key = 'classic-' + index

    return key
  }

}