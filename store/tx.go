package store

import (
	"github.com/icheckteam/explorer/types"
	"gopkg.in/mgo.v2/bson"
)

func (s *Store) GetTxHash(hash string) (*types.Transaction, error) {
	t := &types.Transaction{}
	err := s.txnC.Find(bson.M{"hash": hash}).One(t)
	return t, err
}

func (s *Store) GetTxs(limit, skip int) ([]types.Transaction, error) {
	txns := []types.Transaction{}
	err := s.txnC.Find(bson.M{}).Limit(limit).Skip(skip).Sort("-_id").All(&txns)
	return txns, err
}
