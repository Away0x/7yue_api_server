// components/epsoide/index.js
Component({
  properties: {
    index: {
      type: Number,
      observer(newVal, oldVal, changedPath) {
        if (newVal < 10) {
          // 补零
          this.setData({
            _index: '0' + newVal
          })
        }
      }
    }
  },

  data: {
    // data 里面的属性不能和 properties 里的同名，否则会覆盖
    months: [
      '一月', '二月', '三月', '四月', '五月', '六月', '七月', '八月', '九月', '十月', '十一月',
      '十二月'
    ],
    year: 0,
    month: '',
    _index: ''
  },

  attached() {
    const date = new Date()
    const month = date.getMonth()
    const year = date.getFullYear()

    this.setData({
      month: this.data.months[month],
      year
    })
  },

  methods: {

  }
})
