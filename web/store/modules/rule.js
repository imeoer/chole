import * as types from '../types'

const state = {
  list: []
}

const mutations = {
  [types.INIT](preState, data) {
    const list = data.data.sort((item1, item2) => {
      return item1.name > item2.name
    })
    state.list = list
  },
  [types.CONNECTIONS](preState, data) {
    for (const item of state.list) {
      if (item.id == data.id) {
        item.connections = data.data
      }
    }
  },
  [types.DATA](preState, data) {
    for (const item of state.list) {
      if (item.id == data.id) {
        item.flow = data.data
      }
    }
  }
}

export default {
  state,
  mutations
}
