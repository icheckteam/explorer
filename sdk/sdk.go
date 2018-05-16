package sdk

import (
	"github.com/cosmos/cosmos-sdk/client/context"
	ctypes "github.com/tendermint/tendermint/rpc/core/types"
)

type IchainSdk struct {
	ctx context.CoreContext
}

func NewIchainSdk(ctx context.CoreContext) *IchainSdk {
	return &IchainSdk{ctx: ctx}
}

func (sdk *IchainSdk) GetBlock(height int64) (*ctypes.ResultBlock, error) {
	// get the node
	node, err := sdk.ctx.GetNode()
	if err != nil {
		return nil, err
	}

	// TODO: actually honor the --select flag!
	// header -> BlockchainInfo
	// header, tx -> Block
	// results -> BlockResults
	return node.Block(&height)
}
