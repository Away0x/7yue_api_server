import { config } from '../config.js'

export class HTTP {
  request({url, data = {}, method = 'GET'}) {
    return new Promise((resolve, reject) => {
      wx.request({
        url: config.api_base_url + url,
        method, data,
        header: {
          'content-type': 'application/json',
          'appkey': config.appkey
        },
        success: (res) => {
          const code = String(res.statusCode)

          if (code.startsWith('2')) {
            resolve(res.data)
          } else {
            reject(res)
            wx.showToast({ title: '网络错误', icon: 'none', duration: 2000 })
          }
        },
        fail: (err) => {
          reject(err)
          wx.showToast({ title: '网络错误', icon: 'none', duration: 2000 })
        }
      })
    })
  }
}