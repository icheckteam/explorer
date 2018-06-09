import { RECEIVE_TRANSACTIONS } from '../common/ActionTypes'

const receiveTransactions = transactions => ({
  type: RECEIVE_TRANSACTIONS ,
  transactions
})

export const getAllTransactions = () => dispatch => {
  dispatch(receiveTransactions([
    {height: 100, hash: "232321", fee: 100, type: "createAsset", time: new Date()},
    {height: 101, hash: "1212", fee: 100, type: "createAsset", time: new Date()},
  ]))
}