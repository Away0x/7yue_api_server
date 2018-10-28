import { ClassicModel } from '../../models/classic.js'
import { LikeModel } from '../../models/like.js'

const classicModel = new ClassicModel()
const likeModel = new LikeModel()

Component({

  properties: {
    cid: Number,
    type: Number
  },
  
  data: {
    // 当前的期刊
    currentClassic: null,
    // 是否为最新一期
    latest: true,
    // 是否为第一期
    first: false,
    // 点赞数
    likeCount: 0,
    // 当前用户的点赞状态
    likeStatus: false
  },

  attached(options) {
    const cid = this.properties.cid
    const _type = this.properties.type

    if (!cid) {
      classicModel.getLatest().then(res => {
        this.setData({
          currentClassic: res,
          likeCount: res.fav_nums,
          likeStatus: res.like_status === 1
        })
      })
    } else {
      classicModel.getById(cid, _type).then(res => {
        this._getLikeStatus(res.id, res.type)
        this.setData({
          currentClassic: res,
          latest: classicModel.isLatest(res.index),
          first: classicModel.isFirst(res.index)
        }) 
      })
    }
  },

  methods: {
    onLike(event) {
      const behavior = event.detail.behavior

      likeModel.like(behavior, this.data.currentClassic.id,
        this.data.currentClassic.type)
    },

    // 前一份期刊: 右边按钮
    onPrevious(event) {
      this._updateClassic('previous')
    },

    // 新一份期刊: 左边按钮
    onNext(event) {
      this._updateClassic('next')
    },

    // 更新 classic
    _updateClassic(nextOrPrevious) {
      const index = this.data.currentClassic.index

      classicModel.getClassic(index, nextOrPrevious).then(res => {
        this.setData({
          currentClassic: res,
          latest: classicModel.isLatest(res.index),
          first: classicModel.isFirst(res.index)
        })
        // 更新点赞状态
        this._getLikeStatus(res.id, res.type)
      })
    },

    // 更新点赞状态
    _getLikeStatus(id, type) {
      likeModel.getClassicLikeStatus(id, type).then(res => {
        this.setData({
          likeCount: res.msg.fav_nums,
          likeStatus: res.msg.like_status === 1
        })
      })
    } 
  }

})