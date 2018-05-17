package types

import "time"

type Asset struct {
	ID         string    `json:"id" bson:"height"`
	TxHash     string    `json:"tx_hash" bson:"tx_hash"`
	Name       string    `json:"name" bson:"name"`
	Quantity   int64     `json:"quantity" bson:"quantity"`
	Issuer     string    `json:"issuer" bson:"issuer"`
	Height     int64     `json:"height" bson:"height"`
	Email      string    `json:"email" bson:"email"`
	Company    string    `json:"company" bson:"company"`
	CreateTime time.Time `json:"create_time" bson:"create_time"`
}
