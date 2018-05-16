package types

import "time"

type Block struct {
	Height int64     `json:"height" bson:"height"`
	Hash   string    `json:"hash" bson:"hash"`
	NumTxs int64     `json:"num_txs" bson:"num_txs"`
	Time   time.Time `json:"time" bson:"time"`

	PrevBlock string `json:"prev_block" bson:"prev_block"`
	NextBlock string `json:"next_block" bson:"next_block"`
}
