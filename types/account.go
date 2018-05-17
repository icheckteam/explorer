package types

type AccountAsset struct {
	Address string `json:"address" bson:"address"`
	AssetID string `json:"asset_id" bson:"asset_id"`
	Amount  int64  `json:"amount" bson:"amount"`
}
