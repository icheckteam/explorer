package types

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Transaction struct {
	Hash   string    `json:"hash" bson:"hash"`
	Height int64     `json:"height" bson:"height"`
	Index  int64     `json:"index" bson:"index"`
	Type   string    `json:"type" bson:"type"`
	Time   time.Time `json:"time" bson:"time"`
	Fee    int64     `json:"fee" bson:"fee"`
	Tags   []string  `json:"tags" bson:"tags"`
}

type TransactionInput struct {
	TxHash    string `json:"tx_hash" bson:"tx_hash"`
	Address   string `json:"address" bson:"address"`
	AssetID   string `json:"asset_id" bson:"asset_id"`
	AssetName string `json:"asset_name" bson:"asset_name"`
	Amount    int64  `json:"amount" bson:"amount"`
}

type TransactionOutput struct {
	TxHash    string `json:"tx_hash" bson:"tx_hash"`
	Address   string `json:"address" bson:"address"`
	AssetID   string `json:"asset_id" bson:"asset_id"`
	AssetName string `json:"asset_name" bson:"asset_name"`
	Amount    int64  `json:"amount" bson:"amount"`
}

type TxInfo struct {
	Height int64
	Tx     sdk.Tx
	Hash   string
	Time   time.Time
	Index  int64
}

type Transfer struct {
	TxHash    string `json:"tx_hash" bson:"tx_hash"`
	From      string `json:"from" bson:"from"`
	To        string `json:"to" bson:"to"`
	AssetID   string `json:"asset_id" bson:"asset_id"`
	AssetName string `json:"asset_name" bson:"asset_name"`
	Amount    int64  `json:"amount" bson:"amount"`
}
