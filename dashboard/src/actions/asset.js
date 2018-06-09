import { RECEIVE_ASSETS } from '../common/ActionTypes'

const receiveAssets = assets => ({
  type: RECEIVE_ASSETS ,
  assets
})

export const getAllAssets = () => dispatch => {
  dispatch(receiveAssets([
    {id: "12121", name: "ABC", created: new Date(), numaddrs: 1000, numtxs: 100},
  ]))
}