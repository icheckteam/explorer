package service

import "github.com/cosmos/cosmos-sdk/client/context"

// get the current blockchain height
func GetChainHeight(ctx context.CoreContext) (int64, error) {
	node, err := ctx.GetNode()
	if err != nil {
		return -1, err
	}
	status, err := node.Status()
	if err != nil {
		return -1, err
	}
	height := status.SyncInfo.LatestBlockHeight
	return height, nil
}
