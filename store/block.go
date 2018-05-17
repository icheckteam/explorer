package store

import (
	"github.com/icheckteam/explorer/types"
	"gopkg.in/mgo.v2/bson"
)

func (s *Store) GetBlockHash(hash string) (*types.Block, error) {
	b := &types.Block{}
	err := s.blockC.Find(bson.M{"hash": hash}).One(b)
	return b, err
}

func (s *Store) GetBlockHeight(height int64) (*types.Block, error) {
	b := &types.Block{}
	err := s.blockC.Find(bson.M{"height": height}).One(b)
	return b, err
}

func (s *Store) GetBlocks(limit, skip int) ([]types.Block, error) {
	blocks := []types.Block{}
	err := s.blockC.Find(bson.M{}).Limit(limit).Skip(skip).Sort("-height").All(&blocks)
	return blocks, err
}
