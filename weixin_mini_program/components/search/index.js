import { KeywordModel } from '../../models/keyword.js'
import { BookModel } from '../../models/book.js'
import { paginationBev } from '../../mixins/pagination.js'

const keywordModel = new KeywordModel()
const bookModel = new BookModel()

Component({
  behaviors: [paginationBev],
  /**
   * 组件的属性列表
   */
  properties: {
    more: {
      type: String,
      observer: 'loadMore'
    }
  },

  /**
   * 组件的初始数据
   */
  data: {
    historyWords: [],
    hotWords: [],
    searching: false,
    q: '',
    loading: false,
    loadingCenter: false
  },

  attached() {
    this.setData({ historyWords: keywordModel.getHistory() })

    keywordModel.getHot().then(res => {
      this.setData({ hotWords: res.hot })
    })
  },

  /**
   * 组件的方法列表
   */
  methods: {
    loadMore() {
      if (!this.data.q) { return }
      if (this.isLocked()) { return }

      if (this.hasMore()) {
        this.locked()
        bookModel.search(this.getCurrentStart(), this.data.q)
          .then(res => {
            this.setMoreData(res.books)
            this.unLocked()
          }, () => { this.unLocked() })
      }
    },

    // 回到书籍列表页面
    onCancel(event) {
      this.initialize()
      this.triggerEvent('cancel', {}, {})
    },
    // 清除搜索结果
    onDelete(event) {
      this.initialize()
      this._closeResult()
    },
    onConfirm(event) {
      this._showResult()
      this._showLoadingCenter()
      this.initialize() 
      const q = event.detail.value || event.detail.text
      this.setData({ q })

      bookModel.search(0, q)
        .then(res => {
          this.setMoreData(res.books)
          this.setTotal(res.total)
          keywordModel.addToHistory(q)
          this._hideLoadingCenter()
        })
    },

    _showLoadingCenter() {
      this.setData({ loadingCenter: true })
    },
    _hideLoadingCenter() {
      this.setData({ loadingCenter: false })
    },
    _showResult() {
      this.setData({ searching: true })
    },
    _closeResult() {
      this.setData({ searching: false, q: '' })
    }
  }
})
