package ethclient

import (
	"context"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
)

type EthClient interface {
	BlockByNumber(context.Context, *big.Int) (*types.Block, error)
}

type Db interface {
	StoreBlock(block *types.Block)
}

type ChainReader struct {
	//client *ethclient.Client
	client EthClient
	db Db
}

func (c *ChainReader) CheckBlocks(start, n int64) error {
	for curr := start; curr < start + n; curr++ {
		// Fetch the current block
		block, err := c.client.BlockByNumber(context.Background(), big.NewInt(curr))
		if err != nil {
			return nil
		}
		// Save the block
		c.db.StoreBlock(block)
	}
	return nil
}

////////////////////////
// REFACTORED VERSION //
////////////////////////

func CheckBlocks(start, n int64, client EthClient, db Db) error {
	for curr := start; curr < start + n; curr++ {
		// Fetch the current block
		block, err := client.BlockByNumber(context.Background(), big.NewInt(curr))
		if err != nil {
			return nil
		}
		// Save the block
		db.StoreBlock(block)
	}
	return nil
}