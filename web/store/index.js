import Vue from 'vue'
import Vuex from 'vuex'
import createLogger from 'vuex/logger'
import config from './modules/config'

Vue.use(Vuex)

const debug = process.env.NODE_ENV !== 'production'

const connect = () => {
  let reTimer = null
  const host = `${location.origin.replace(/^http/, 'ws')}/push`

  // const conn = new WebSocket('ws://' + location.host + '/push')
  const conn = new WebSocket('ws://localhost:8081/push')

  conn.onmessage = function(event) {
    console.log(event)
  }

  conn.onclose = function() {
    alert('CLOSE')
    // if (reTimer) clearTimeout(reTimer)
    // reTimer = setTimeout(() => {
    //   connect()
    // }, 1000)
  }
}

export default new Vuex.Store({
  modules: {
    config
  },
  strict: debug,
  plugins: debug ? [createLogger()] : []
})

connect()