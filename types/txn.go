package types

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Transaction struct {
	Hash   string
	Height int64
	Index  int64
	Type   string
	Time   time.Time
	Fee    int64
}

type Txinfo struct {
	Height int64
	Tx     sdk.StdTx
	Hash   string
	Time   time.Time
}
