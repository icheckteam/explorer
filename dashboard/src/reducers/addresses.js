import {
  RECEIVE_ADDRESSES
} from '../common/ActionTypes'

const initialState = {
  addresses: [],
}


export default function addresses(state = initialState, action) {
  switch(action.type) {
    case RECEIVE_ADDRESSES:
      return {
        ...state,
        addresses: action.addresses,
      }
    default:
      return state;
  }
}


