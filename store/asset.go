package store

import (
	"github.com/icheckteam/explorer/types"
	"gopkg.in/mgo.v2/bson"
)

func (s *Store) GetAsset(id string) (*types.Asset, error) {
	t := &types.Asset{}
	err := s.assetC.Find(bson.M{"id": id}).One(t)
	return t, err
}

func (s *Store) GetAssets(limit, skip int) ([]types.Asset, error) {
	txns := []types.Asset{}
	err := s.txnC.Find(bson.M{}).Limit(limit).Skip(skip).Sort("-_id").All(&txns)
	return txns, err
}

func (s *Store) GetAssetsOfIssuer(issuer string, limit, skip int) ([]types.Asset, error) {
	txns := []types.Asset{}
	err := s.txnC.Find(bson.M{"issuer": issuer}).Limit(limit).Skip(skip).Sort("-_id").All(&txns)
	return txns, err
}
