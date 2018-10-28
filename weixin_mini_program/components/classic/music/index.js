import { classicBehavior } from '../behavior.js'

// 音乐管理对象
const mMgr = wx.getBackgroundAudioManager()

Component({
  behaviors: [classicBehavior],
  
  properties: {
    src: String,
    title: String
  },
  
  // 老师说的 hidden 方案，其实可通过监听 hidden 值实现 wx:if 的 attached 功能
  attached() {
    this._recoverPayingStatus()
    this._musicManngerEventHandler()
  },

  data: {
    playing: false,
    waittingUrl: 'images/player@waitting.png',
    playingUrl: 'images/player@playing.png'
  },

  methods: {
    // 开始播放
    onPlay(event) {
      // 播放音乐
      if ( ! this.data.playing) {
        this.setData({ playing: true })
        mMgr.src = this.properties.src
      }
      // 暂停音乐
      else {
        this.setData({ playing: false })
        mMgr.pause()
      }
    },

    // 如果当前有正在播放的音乐，且播放的正式当前展示的音乐，那么设置其为播放状态
    // 否则设置为暂停状态
    _recoverPayingStatus() {
      // 处于暂停状态
      if (mMgr.paused) {
        this.setData({ playing: false })
        return
      }
      // 处于播放状态
      if (mMgr.src === this.properties.src) {
        this.setData({ playing: true })
      }
    },

    // 监听音乐播放事件 (有可能不是通过触发上面的 onPlay 事件暂停播放音乐的)
    // 这里对音乐管理对象的各个事件进行监听，并实时更新组件状态
    _musicManngerEventHandler() {
      // 音乐播放器播放
      mMgr.onPlay(() =>  { this._recoverPayingStatus() })
      // 音乐播放器暂停
      mMgr.onPause(() => { this._recoverPayingStatus() })
      // 音乐播放器关闭
      mMgr.onStop(() =>  { this._recoverPayingStatus() })
      // 音乐播放完成
      mMgr.onEnded(() => { this._recoverPayingStatus() })
    }
  }
})
