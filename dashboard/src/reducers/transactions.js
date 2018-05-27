import {
  RECEIVE_TRANSACTIONS,
} from '../common/ActionTypes'

const initialState = {
  transactions: [],
}


export default function blocks(state = initialState, action) {
  switch(action.type) {
    case RECEIVE_TRANSACTIONS:
      return {
        ...state,
        transactions: action.transactions,
      }
    default:
      return state;
  }
}


