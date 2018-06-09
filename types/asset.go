package types

import "time"

type Asset struct {
	ID          string    `json:"id" bson:"height"`
	TxHash      string    `json:"tx_hash" bson:"tx_hash"`
	Name        string    `json:"name" bson:"name"`
	Quantity    int64     `json:"quantity" bson:"quantity"`
	Issuer      string    `json:"issuer" bson:"issuer"`
	Height      int64     `json:"height" bson:"height"`
	Email       string    `json:"email" bson:"email"`
	Company     string    `json:"company" bson:"company"`
	CreateTime  time.Time `json:"create_time" bson:"create_time"`
	NumAccounts int64     `json:"num_accounts" bson:"num_accounts"`
	NumTxs      int64     `json:"num_txs" bson:"num_txs"`
}

// Property ...
type Property struct {
	AssetID      string    `json:"asset_id" bson:"asset_id"`
	Height       int64     `json:"height" bson:"height"`
	Reporter     string    `json:"reporter" bson:"reporter"`
	Name         string    `json:"name" bson:"name"`
	Type         int       `json:"type" bson:"type"`
	BytesValue   []byte    `json:"bytes_value" bson:"bytes_value"`
	StringValue  string    `json:"string_value" bson:"string_value"`
	BooleanValue bool      `json:"boolean_value" bson:"boolean_value"`
	NumberValue  int64     `json:"number_value" bson:"number_value"`
	EnumValue    []string  `json:"enum_value" bson:"enum_value"`
	Location     Location  `json:"location_value" bson:"location_value"`
	CreateTime   time.Time `json:"create_time" bson:"create_time"`
}

type Location struct {
	Latitude  float64 `json:"latitude" amino:"unsafe"`
	Longitude float64 `json:"longitude" amino:"unsafe"`
}
