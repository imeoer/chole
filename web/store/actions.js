import * as types from './types'

export const toggleEditor = ({ dispatch }, flag) => {
  dispatch(types.TOGGLE_EDITOR, flag)
}
