package store

import (
	"github.com/icheckteam/explorer/types"
	ttypes "github.com/tendermint/tendermint/types"
	"gopkg.in/mgo.v2/bson"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var txInfo struct {
	Hash string
	Height int64
	Tx sdk.Tx
	Time time.Time
	Index uint64
}

// insertBlock....
func (s *Store) insertBlock(block *ttypes.Block) error {
	var err error
	b := &types.Block{
		Height:    block.Header.Height,
		Hash:      block.Hash().String(),
		NumTxs:    block.Header.NumTxs,
		Time:      block.Header.Time,
		PrevBlock: block.Header.LastBlockID,
	}
	// insert new block ...
	if err = s.blockC.Insert(b); err != nil {
		return err
	}

	// update next block
	if b.Height > 1 {
		return s.updateNextBlockHash(hash, b.Height-1)
	}

	if b.NumTxs == 0 {
		return nil
	}

	// insert txs 
	for i, tx := block.Data.Txs {
		txn, err:= parseTx(s.cdc, txBytes)
		if err != nil {
			return err
		}
		info := txInfo{
			Hash: tx.Hash().String(),
			Height: b.Height,
			Time: b.Time,
			Tx: txn,
			Index: i,
		}
		if err = s.insertTxn(info); err != nil {
			return err
		}
	}
	return nil
}

// updateNextBlockHash ...
func (s *Store) updateNextBlockHash(hash string, height uint64) error {
	return s.blockC.Update(bson.M{
		"height": height,
	}, bson.M{"$set": bson.M{
		"next_block": hash,
	}})
}


func parseTx(cdc *wire.Codec, txBytes []byte) (sdk.Tx, error) {
	var tx sdk.StdTx
	err := cdc.UnmarshalBinary(txBytes, &tx)
	if err != nil {
		return nil, err
	}
	return tx, nil
}





