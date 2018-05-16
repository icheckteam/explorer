package store

import (
	"time"

	"github.com/icheckteam/explorer/types"
	"github.com/icheckteam/ichain/x/asset"
	"github.com/icheckteam/ichain/x/bank"
)

func (s *Store) insertTxn(info txInfo) error {
	// TODO: can tx just implement message?
	msg := info.Tx.GetMsg()

	switch msg := msg.(type) {
	case bank.MsgSend:
		return s.insertSendTxn(msg, info)
	case asset.RegisterMsg:
		return s.insertRegisterAssetMsg(msg, info)

	}
	return nil
}

func (s *Store) insertSendTxn(msg bank.MsgIssue, info txInfo) error {
	return nil
}

// insert register asset msg
func (s *Store) insertRegisterAssetMsg(msg bank.MsgIssue, info txInfo) error {
	asset := types.Asset{
		ID:         tx.ID,
		Height:     height,
		Issuer:     asset.Issuer.String(),
		Name:       asset.Name,
		Quantity:   asset.Quantity,
		Company:    asset.Company,
		Email:      asset.Email,
		CreateTime: time.Time,
	}
	if err := s.assetC.Insert(asset); err != nil {
		return err
	}
	// insert txn basic info
	return s.insertTxnBasicInfo("asset/register", info)
}

func (s *Store) insertTxnBasicInfo(txType string, info txInfo) error {
	tx := types.Transaction{
		Hash:   info.Hash,
		Height: info.Height,
		Time:   info.Time,
		Fee:    0,
		Index:  info.Index,
		Type:   txType,
	}
	return s.txC.Insert(tx)
}
