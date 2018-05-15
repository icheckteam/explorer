package service

import mgo "gopkg.in/mgo.v2"

// Service ...
type Store struct {
	blockC   *mgo.Collation
	txC      *mgo.Collation
	currentC *mgo.Collation
}
