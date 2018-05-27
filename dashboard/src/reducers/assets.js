import {
  RECEIVE_ASSETS
} from '../common/ActionTypes'

const initialState = {
  assets: [],
}


export default function assets(state = initialState, action) {
  switch(action.type) {
    case RECEIVE_ASSETS:
      return {
        ...state,
        assets: action.assets,
      }
    default:
      return state;
  }
}


