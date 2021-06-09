package ethclient

import (
	"github.com/ansermino/go-mock-demo/ethclient/mock"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"math/big"
	"testing"
)

func TestChainInteractor_CheckBlocks(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	client := mock_ethclient.NewMockEthClient(ctrl)
	db := mock_ethclient.NewMockDb(ctrl)

	chain := &ChainReader{client: client, db: db}

	client.EXPECT().BlockByNumber(gomock.Any(), big.NewInt(0)).Return(&types.Block{}, nil)
	db.EXPECT().StoreBlock(gomock.Any())

	client.EXPECT().BlockByNumber(gomock.Any(), big.NewInt(1)).Return(&types.Block{}, nil)
	db.EXPECT().StoreBlock(gomock.Any())

	client.EXPECT().BlockByNumber(gomock.Any(), big.NewInt(2)).Return(&types.Block{}, nil)
	db.EXPECT().StoreBlock(gomock.Any())


	err := chain.CheckBlocks(0, 3)
	assert.Nil(t, err)
}

func TestCheckBlocks(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	client := mock_ethclient.NewMockEthClient(ctrl)
	db := mock_ethclient.NewMockDb(ctrl)

	client.EXPECT().BlockByNumber(gomock.Any(), big.NewInt(0)).Return(&types.Block{}, nil)
	db.EXPECT().StoreBlock(gomock.Any())

	client.EXPECT().BlockByNumber(gomock.Any(), big.NewInt(1)).Return(&types.Block{}, nil)
	db.EXPECT().StoreBlock(gomock.Any())

	client.EXPECT().BlockByNumber(gomock.Any(), big.NewInt(2)).Return(&types.Block{}, nil)
	db.EXPECT().StoreBlock(gomock.Any())

	err := CheckBlocks(0, 3, client, db)
	assert.Nil(t, err)
}