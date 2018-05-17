package store

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/icheckteam/explorer/types"
	"github.com/icheckteam/ichain/x/asset"
	"github.com/icheckteam/ichain/x/bank"
	ttypes "github.com/tendermint/tendermint/types"
	"gopkg.in/mgo.v2/bson"
)

func (s *Store) InsertTxn(info types.TxInfo) error {
	msg := info.Tx.GetMsg()

	switch msg := msg.(type) {
	case asset.RegisterMsg:
		return s.insertRegisterAssetTxn(msg, info)
	case bank.MsgSend:
		return s.insertSendTx(msg, info)
	default:
		errMsg := fmt.Sprintf("Unrecognized trace Msg type: %v", reflect.TypeOf(msg).Name())
		return errors.New(errMsg)
	}
}

func (s *Store) insertSendTx(msg bank.MsgSend, info types.TxInfo) error {
	var err error
	inputs := []types.TransactionInput{}
	assets := map[string]*types.Asset{}
	for _, in := range msg.Inputs {
		for _, coin := range in.Coins {
			assets[coin.Denom], err = s.GetAsset(coin.Denom)
			if err != nil {
				return err
			}
			inputs = append(inputs, types.TransactionInput{
				Address:   in.Address.String(),
				Amount:    coin.Amount,
				AssetID:   coin.Denom,
				AssetName: assets[coin.Denom].Name,
				TxHash:    info.Hash,
			})
		}
	}
	outputs := []types.TransactionOutput{}
	for _, in := range msg.Inputs {
		for _, coin := range in.Coins {
			outputs = append(outputs, types.TransactionOutput{
				Address:   in.Address.String(),
				Amount:    coin.Amount,
				AssetID:   coin.Denom,
				AssetName: assets[coin.Denom].Name,
				TxHash:    info.Hash,
			})
		}
	}

	if err = s.txnInputC.Insert(inputs); err != nil {
		return err
	}
	if err = s.txnOutputC.Insert(outputs); err != nil {
		return err
	}
	return s.insertTxBasicInfo("bank/send", info)
}

func (s *Store) insertRegisterAssetTxn(msg asset.RegisterMsg, info types.TxInfo) error {
	asset := &types.Asset{
		ID:         msg.ID,
		Name:       msg.Name,
		Quantity:   msg.Quantity,
		Issuer:     msg.Issuer.String(),
		Email:      msg.Email,
		Company:    msg.Company,
		TxHash:     info.Hash,
		CreateTime: info.Time,
	}
	if err := s.assetC.Insert(asset); err != nil {
		return err
	}
	// insert tx basic info
	return s.insertTxBasicInfo("asset/register", info)
}

func (s *Store) insertTxBasicInfo(txtype string, info types.TxInfo) error {
	tx := types.Transaction{
		Hash:   info.Hash,
		Height: info.Height,
		Fee:    0,
		Index:  info.Index,
		Time:   info.Time,
		Type:   txtype,
	}
	return s.txnC.Insert(&tx)
}

// GetCurrentHeight ....
func (s *Store) GetCurrentHeight() (int64, error) {
	b := &types.Block{}
	err := s.blockC.Find(bson.M{}).Sort("-height").One(b)
	if err != nil {
		return 0, err
	}
	return b.Height, nil
}

// GetCurrentHeight ....
func (s *Store) UpdateNextBlockHash(hash string, height int64) error {
	return s.blockC.Update(bson.M{
		"height": height,
	}, bson.M{"$set": bson.M{"next_block": hash}})
}

// GetCurrentHeight ....
func (s *Store) InsertBlock(block *ttypes.Block) error {
	b := &types.Block{
		Hash:      block.Hash().String(),
		Height:    block.Height,
		NextBlock: "",
		NumTxs:    block.NumTxs,
		Time:      block.Time,
	}

	if !block.LastBlockID.IsZero() {
		b.PrevBlock = block.LastBlockID.String()
	}

	return s.blockC.Insert(b)
}
