import { RECEIVE_BLOCKS } from '../common/ActionTypes'

const receiveBlocks = blocks => ({
  type: RECEIVE_BLOCKS ,
  blocks
})

let height = 0;
function createData(numtxs, size) {
  height += 1;
  var time = new Date()
  return { height, time, numtxs, size, validator: "APyEx5f4Zm4oCHwFWiSTaph1fPBxZacYVR" };
}


export const getAllBlocks = () => dispatch => {
  dispatch(receiveBlocks([
    createData(1, 2),
    createData(2, 5),
  ]))
}