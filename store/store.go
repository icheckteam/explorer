package store

import (
	mgo "gopkg.in/mgo.v2"
)

// Service ...
type Store struct {
	blockC        *mgo.Collection
	txnC          *mgo.Collection
	assetC        *mgo.Collection
	txnInputC     *mgo.Collection
	txnOutputC    *mgo.Collection
	accountAssetC *mgo.Collection
	assetIssueC   *mgo.Collection
	assetSubtracC *mgo.Collection
}

// NewStore ...
func NewStore(db *mgo.Database) *Store {
	return &Store{
		blockC:        db.C("blocks"),
		assetC:        db.C("assets"),
		txnC:          db.C("transactions"),
		txnInputC:     db.C("transaction_inputs"),
		txnOutputC:    db.C("transaction_output"),
		accountAssetC: db.C("account_assets"),
		assetIssueC:   db.C("asset_issues"),
		assetSubtracC: db.C("asset_subtract"),
	}
}
