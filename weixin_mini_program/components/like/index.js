// components/like/index.js
Component({
  /**
   * 组件的属性列表
   */
  properties: {
    like: { type: Boolean, value: false },
    count: { type: Number, value: 0 }
  },

  /**
   * 组件的初始数据
   */
  data: {
    yesSrc: 'images/like.png',
    noSrc: 'images/like@dis.png'
  },

  /**
   * 组件的方法列表
   */
  methods: {
    onLike(event) {
      const { like, count } = this.properties

      this.setData({
        count: like ? count - 1 : count + 1,
        like: !like
      })

      let behavior = this.properties.like ? 'like' : 'cancel'

      this.triggerEvent('like', { behavior })
    }
  }
})
