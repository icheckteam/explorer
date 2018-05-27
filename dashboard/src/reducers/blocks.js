import {
  RECEIVE_BLOCKS,
} from '../common/ActionTypes'

const initialState = {
  blocks: [],
}


export default function blocks(state = initialState, action) {
  switch(action.type) {
    case RECEIVE_BLOCKS:
      return {
        ...state,
        blocks: action.blocks,
      }
    default:
      return state;
  }
}


