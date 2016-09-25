import * as types from '../types'

const state = {
  status: false,
  config: ''
}

const mutations = {
  [types.TOGGLE_STATUS](preState, flag) {
    state.status = flag
  }
}

export default {
  state,
  mutations
}
