import { combineReducers } from 'redux'
import blocks from './blocks'
import transactions from './transactions'
import addresses from './addresses'
import assets from './assets';
export default combineReducers({
  blocks,
  transactions,
  addresses,
  assets,
})