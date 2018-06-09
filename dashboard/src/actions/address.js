import { RECEIVE_ADDRESSES } from '../common/ActionTypes'

const receiveAddresses = addresses => ({
  type: RECEIVE_ADDRESSES ,
  addresses
})

export const getAllAddresses = () => dispatch => {
  dispatch(receiveAddresses([
    {address: "12121", height: 100, created: new Date(), last_tx: new Date(), numtxs: 100},
  ]))
}