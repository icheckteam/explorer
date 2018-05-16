package store

import (
	"github.com/cosmos/cosmos-sdk/wire"
	mgo "gopkg.in/mgo.v2"
)

// Service ...
type Store struct {
	blockC   *mgo.Collection
	txC      *mgo.Collection
	currentC *mgo.Collection
	assetC   *mgo.Collection

	cdc *wire.Codec
}
