package store

import (
	mgo "gopkg.in/mgo.v2"
)

// Service ...
type Store struct {
	blockC     *mgo.Collection
	txnC       *mgo.Collection
	assetC     *mgo.Collection
	txnInputC  *mgo.Collection
	txnOutputC *mgo.Collection
}

// NewStore ...
func NewStore(db *mgo.Database) *Store {
	return &Store{
		blockC:     db.C("blocks"),
		assetC:     db.C("assets"),
		txnC:       db.C("transactions"),
		txnInputC:  db.C("transaction_inputs"),
		txnOutputC: db.C("transaction_output"),
	}
}
