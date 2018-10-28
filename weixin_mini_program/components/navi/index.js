// components/navi/index.js
Component({
  /**
   * 组件的属性列表
   */
  properties: {
    title: {
      type: String,
      value: '...'
    },
    latest: {
      type: Boolean,
      value: false
    },
    first: {
      type: Boolean,
      value: false
    }
  },

  /**
   * 组件的初始数据
   */
  data: {
    disLeftSrc: 'images/triangle.dis@left.png',
    highLeftSrc: 'images/triangle@left.png'
  },

  /**
   * 组件的方法列表
   */
  methods: {
    onLeft: function () {
      if ( ! this.properties.latest) {
        this.triggerEvent('left')
      }
    },
    onRight: function () {
      if ( ! this.properties.first) {
        this.triggerEvent('right')
      }
    }
  }
})
