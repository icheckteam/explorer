package types

import "time"

type Asset struct {
	ID         string
	TxHash     string
	Name       string
	Quantity   int64
	Issuer     string
	Height     int64
	Email      string
	Company    string
	CreateTime time.Time
}
