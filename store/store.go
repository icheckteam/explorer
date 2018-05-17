package store

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/icheckteam/explorer/types"
	"github.com/icheckteam/ichain/x/asset"
	ttypes "github.com/tendermint/tendermint/types"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Service ...
type Store struct {
	blockC *mgo.Collection
	txnC   *mgo.Collection
	assetC *mgo.Collection
}

// NewStore ...
func NewStore(db *mgo.Database) *Store {
	return &Store{
		blockC: db.C("blocks"),
		txnC:   db.C("transactions"),
		assetC: db.C("assets"),
	}
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

// GetCurrentHeight ....
func (s *Store) UpdateNextBlockHash(hash string, height int64) error {
	return s.blockC.Update(bson.M{
		"height": height,
	}, bson.M{"$set": bson.M{"next_block": hash}})
}

func (s *Store) InsertTxn(info types.TxInfo) error {
	msg := info.Tx.GetMsg()

	switch msg := msg.(type) {
	case asset.RegisterMsg:
		return s.insertRegisterAssetTxn(msg, info)
	default:
		errMsg := fmt.Sprintf("Unrecognized trace Msg type: %v", reflect.TypeOf(msg).Name())
		return errors.New(errMsg)
	}
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
