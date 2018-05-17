package store

import "gopkg.in/mgo.v2/bson"

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
