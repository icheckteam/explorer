package types

import "time"

type Transaction struct {
	Hash   string
	Height int64
	Type   string
	Time   time.Time
	Fee    int64
}
