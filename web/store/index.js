import Vue from 'vue'
import Vuex from 'vuex'
import createLogger from 'vuex/logger'
import msgpack from 'msgpack-lite'
import app from './modules/app'
import rule from './modules/rule'
import * as types from './types'

Vue.use(Vuex)

const debug = process.env.NODE_ENV !== 'production'

const connect = (store) => {
  let reTimer = null
  const host = `${location.origin.replace(/^http/, 'ws')}/push`

  // const conn = new WebSocket('ws://' + location.host + '/push')
  const conn = new WebSocket('ws://192.168.1.9:8081/push')
  conn.binaryType = 'arraybuffer'

  conn.onopen = () => {
    store.dispatch(types.TOGGLE_STATUS, true)
  }

  conn.onmessage = (event) => {
    const data = msgpack.decode(new Uint8Array(event.data))
    console.log(data)
    store.dispatch(types[data.event], {
      id: data.id,
      data: data.data
    })
  }

  conn.onclose = ()  => {
    store.dispatch(types.TOGGLE_STATUS, false)
    if (reTimer) clearTimeout(reTimer)
    reTimer = setTimeout(() => {
      connect(store)
    }, 1000)
  }
}

const store = new Vuex.Store({
  modules: {
    app,
    rule
  },
  strict: debug,
  plugins: debug ? [createLogger()] : []
})

connect(store)

export default store
