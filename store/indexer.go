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
	case asset.AddQuantityMsg:
		return s.insertIssueAssetTxn(msg, info)
	case asset.SubtractQuantityMsg:
		return s.insertAssetSubtractTxn(msg, info)
	default:
		errMsg := fmt.Sprintf("Unrecognized trace Msg type: %v", reflect.TypeOf(msg).Name())
		return errors.New(errMsg)
	}
}

func (s *Store) insertSendTx(msg bank.MsgSend, info types.TxInfo) error {
	var err error
	var tags []string
	inputs := []types.TransactionInput{}
	assets := map[string]*types.Asset{}
	for _, in := range msg.Inputs {
		tags = append(tags, "account/"+in.Address.String())
		for _, coin := range in.Coins {
			assets[coin.Denom], err = s.GetAsset(coin.Denom)
			if err != nil {
				return err
			}
			if err := s.addCoin(in.Address.String(), coin.Denom, -coin.Amount); err != nil {
				return err
			}
			tags = append(tags, "asset/"+coin.Denom)
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
	for _, out := range msg.Outputs {
		tags = append(tags, "account/"+out.Address.String())
		for _, coin := range out.Coins {
			if err := s.addCoin(out.Address.String(), coin.Denom, coin.Amount); err != nil {
				return err
			}
			outputs = append(outputs, types.TransactionOutput{
				Address:   out.Address.String(),
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
	return s.insertTxBasicInfo("bank/send", tags, info)
}

func (s *Store) insertRegisterAssetTxn(msg asset.RegisterMsg, info types.TxInfo) error {
	asset := &types.Asset{
		ID:          msg.ID,
		Name:        msg.Name,
		Quantity:    msg.Quantity,
		Issuer:      msg.Issuer.String(),
		Email:       msg.Email,
		Company:     msg.Company,
		TxHash:      info.Hash,
		CreateTime:  info.Time,
		NumTxs:      1,
		NumAccounts: 1,
	}
	if err := s.assetC.Insert(asset); err != nil {
		return err
	}

	// add coin ...
	if err := s.addCoin(asset.Issuer, asset.ID, asset.Quantity); err != nil {
		return err
	}

	// insert tx basic info
	return s.insertTxBasicInfo("asset/register", []string{
		"asset/" + asset.ID,
		"account/" + asset.Issuer,
	}, info)
}

func (s *Store) insertTxBasicInfo(txtype string, tags []string, info types.TxInfo) error {
	tx := types.Transaction{
		Hash:   info.Hash,
		Height: info.Height,
		Fee:    0,
		Index:  info.Index,
		Time:   info.Time,
		Type:   txtype,
		Tags:   tags,
	}
	return s.txnC.Insert(&tx)
}

func (s *Store) insertIssueAssetTxn(msg asset.AddQuantityMsg, info types.TxInfo) error {
	issue := types.AssetIssue{
		Amount:     msg.Quantity,
		AssetID:    msg.ID,
		Issuer:     msg.Issuer.String(),
		CreateTime: info.Time,
	}
	// add coin ...issue
	if err := s.addCoin(issue.Issuer, issue.AssetID, issue.Amount); err != nil {
		return err
	}
	return s.assetIssueC.Insert(issue)
}

func (s *Store) insertAssetSubtractTxn(msg asset.SubtractQuantityMsg, info types.TxInfo) error {
	issue := types.AssetSubtract{
		Amount:     msg.Quantity,
		AssetID:    msg.ID,
		Issuer:     msg.Issuer.String(),
		CreateTime: info.Time,
	}
	// add coin ...issue
	if err := s.addCoin(issue.Issuer, issue.AssetID, issue.Amount); err != nil {
		return err
	}
	return s.assetSubtracC.Insert(issue)
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
