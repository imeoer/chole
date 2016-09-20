import * as types from '../types'

const state = {
  editor: false
}

const mutations = {
  [types.TOGGLE_EDITOR](preState, flag) {
    state.editor = !state.editor
  }
}

export default {
  state,
  mutations
}
