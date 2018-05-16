package sdk

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type TxInfo struct {
	Height int64
	Tx     sdk.StdTx
	Index  int64
	Time   time.Time
}
