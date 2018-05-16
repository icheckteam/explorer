package types

import "time"

type Block struct {
	Height int64     `json:"height"`
	Hash   string    `json:"hash"`
	NumTxs int64     `json:"num_txs"`
	Time   time.Time `json:"time"`

	PrevBlock string `json:"prev_block"`
	NextBlock string `json:"next_block"`
}
