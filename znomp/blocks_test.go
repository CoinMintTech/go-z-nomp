package znomp_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/CoinMintTech/go-z-nomp/znomp"
)

const (
	address        = "abcdefghijklmnopqrstuvwxyz123456789"
	address_suffix = "asics1"
	block          = uint(123456)
	coin           = "zclassic"
	timestamp      = 1535160554891
)

var (
	blocksResponse = []byte(fmt.Sprintf(
		`{"%s-%d":"073e8c1e63274117a761bfbc690e7987073e8c1e63274117a761bfbc690e7987:cc04db7ab9de4b24843a3fe0f16e47fecc04db7ab9de4b24843a3fe0f16e47fe:123456:%s.%s:%d"}`,
		coin,
		block,
		address,
		address_suffix,
		timestamp,
	))
)

func TestBlocksResponse(t *testing.T) {
	var (
		err error
		res = make(znomp.BlocksResponse, 0)
	)

	err = json.Unmarshal(blocksResponse, &res)
	if assert.NotNil(t, res) && assert.Nil(t, err) {
		if assert.Equal(t, 1, len(res)) {
			b := []znomp.Block(res)[0]
			assert.Equal(t, coin, b.Coin)
			assert.Equal(t, fmt.Sprintf("%s.%s", address, address_suffix), b.Miner)
		}
	}
}
