import { HTTP } from '../utils/http.js'

export class LikeModel extends HTTP {
  // 获取最新一期
  like(behavior, artID, category) {
    return this.request({
      url: behavior === 'like' ? 'like' : 'like/cancel',
      method: 'POST',
      data: { art_id: artID, type: category },
    })
  }

  // 获取点赞信息
  getClassicLikeStatus(id, type) {
    return this.request({ url: `classic/favor/${type}/${id}` })
  }

  // 点赞/ 取消点赞
  like(like_or_cancel, art_id, category) {
    return this.request({
      url: like_or_cancel === 'like' ? `like` : 'like/cancel',
      method: 'POST',
      data: { art_id: Number(art_id), type: category }
    })
  }
}