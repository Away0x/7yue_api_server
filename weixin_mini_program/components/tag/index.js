Component({
  
  options: {
    multipleSlots: true // 在组件定义时的选项中启用多 slot 支持
  },
  externalClasses: ['tag-class'], // 外部 class

  properties: {
    text: String
  },

  data: {

  },

  methods: {
    onTap(event) {
      this.triggerEvent('tapping', { text: this.properties.text })
    }
  }
})
