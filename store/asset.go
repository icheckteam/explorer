package store

import (
	"github.com/icheckteam/explorer/types"
	"github.com/icheckteam/ichain/x/asset"
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

func (s *Store) GetAssetTxs(assetID string, limit, skip int) ([]types.Asset, error) {
	txns := []types.Asset{}
	err := s.txnC.Find(bson.M{"tags": "asset/" + assetID}).Limit(limit).Skip(skip).Sort("-_id").All(&txns)
	return txns, err
}

func (s *Store) GetAssetAccounts(assetID string, limit, skip int) ([]types.AccountAsset, error) {
	txns := []types.AccountAsset{}
	err := s.accountAssetC.Find(bson.M{"asset_id": assetID}).Limit(limit).Skip(skip).Sort("-_id").All(&txns)
	return txns, err
}

func (s *Store) GetAssetTransfers(assetID string, limit, skip int) ([]types.Transfer, error) {
	inputs := []types.TransactionInput{}
	var transfers []types.Transfer
	err := s.txnInputC.Find(bson.M{"asset_id": assetID}).Limit(limit).Skip(skip).Sort("-_id").All(&inputs)
	if err != nil {
		return nil, err
	}

	for _, in := range inputs {
		out := types.TransactionOutput{}
		err := s.txnOutputC.Find(bson.M{"asset_id": assetID}).Limit(limit).Skip(skip).Sort("-_id").One(&out)
		if err != nil {
			return nil, err
		}

		transfers = append(transfers, types.Transfer{
			Amount:    in.Amount,
			AssetID:   in.AssetID,
			AssetName: in.AssetName,
			From:      in.Address,
			To:        out.AssetID,
		})
	}
	return transfers, nil
}

func (s *Store) insertPropertipes(msg asset.UpdateAttrMsg, info types.TxInfo) error {
	propertipes := []*types.Property{}
	for _, property := range msg.Attributes {
		p := &types.Property{
			AssetID:     msg.ID,
			Height:      info.Height,
			Reporter:    msg.Issuer.String(),
			CreateTime:  info.Time,
			Name:        property.Name,
			Type:        property.Type,
			StringValue: property.StringValue,
			BytesValue:  property.BytesValue,
			EnumValue:   property.EnumValue,
			Location: types.Location{
				Latitude:  property.Location.Latitude,
				Longitude: property.Location.Longitude,
			},
			NumberValue:  property.NumberValue,
			BooleanValue: property.BooleanValue,
		}

		propertipes = append(propertipes, p)
	}
	return s.propertyC.Insert(propertipes)
}

func (s *Store) GetAssetHistory(assetID string, name string) ([]types.Property, error) {
	propertipes := []types.Property{}
	err := s.propertyC.Find(bson.M{"asset_id": assetID, "name": name}).All(&propertipes)
	return propertipes, err
}
