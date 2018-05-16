package store

import (
	"time"
	"github.com/icheckteam/explorer/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/icheckteam/ichain/x/bank"
	"github.com/icheckteam/ichain/x/asset"
)

func (s *Store) insertTxn(tx sdk.Tx, height int64, index int64, t time.Time) error {
	// TODO: can tx just implement message?
	msg := tx.GetMsg()

	switch msg := msg.(type) {
	case bank.MsgSend:
		return s.insertSendTxn(tx sdk.Tx, height int64, index int64, t time.Time)
	case asset.RegisterMsg:
		return s.insertRegisterAssetMsg(tx sdk.Tx, height int64, index int64, t time.Time)

	}
	return nil
}

func (s *Store) insertSendTxn(tx sdk.Tx, height int64, index int64, t time.Time) error {
	return nil
}

// insert register asset msg 
func (s *Store) insertRegisterAssetMsg(tx asset.RegisterMsg, height int64, index int64, t time.Time) error {
	asset := types.Asset{
		ID: tx.ID,
		Height: height,
		Issuer: asset.Issuer.String(),
		Name: asset.Name,
		Quantity: asset.Quantity,
		Company: asset.Company,
		Email: asset.Email,
	}
	return s.assetC.Insert(asset)
}
