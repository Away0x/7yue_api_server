Component({

  properties: {
    book: Object,
    showLike: {
      type: Boolean,
      value: true
    }
  },

  data: {
    title: '',
    author: '',
    img: ''
  },

  methods: {
    onTap(event) {
      const bid = this.properties.book.id
      const isbn = this.properties.book.isbn
      
      // this.triggerEvent('booktap', { bid })
      wx.navigateTo({
        url: `/pages/book-detail/book-detail?bid=${bid}&isbn=${isbn}`,
      })
    }
  }
})
