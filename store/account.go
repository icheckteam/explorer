package store

import (
	"github.com/icheckteam/explorer/types"
	"gopkg.in/mgo.v2/bson"
)

func (s *Store) addCoin(addr, assetID string, amount int64) error {
	_, err := s.accountAssetC.Upsert(bson.M{
		"asset_id": assetID,
		"address":  addr,
	}, bson.M{
		"$set": bson.M{
			"asset_id": assetID,
			"address":  addr,
		},
		"$inc": bson.M{
			"amount": amount,
		},
	})
	return err
}

func (s *Store) GetAccountAssets(addr string) ([]types.AccountAsset, error) {
	assets := []types.AccountAsset{}
	err := s.accountAssetC.Find(bson.M{"address": addr}).All(assets)
	return assets, err
}

func (s *Store) GetAccountTxs(addr string, limit, skip int) ([]types.AccountAsset, error) {
	assets := []types.AccountAsset{}
	err := s.txnC.Find(bson.M{"tags": "account/" + addr}).Skip(skip).Limit(limit).All(assets)
	return assets, err
}
