package znomp_test

import (
	"fmt"
	"testing"
)

const (
	address        = "abcdefghijklmnopqrstuvwxyz123456789"
	address_suffix = "asics1"
	block          = uint(123456)
	coin           = "zclassic"
	timestamp      = 1535160554891
)

var (
	blocksResponse = fmt.Sprintf(
		`{"%s-%d":"073e8c1e63274117a761bfbc690e7987073e8c1e63274117a761bfbc690e7987:cc04db7ab9de4b24843a3fe0f16e47fecc04db7ab9de4b24843a3fe0f16e47fe:123456:%s.%s:%d"}`,
		coin,
		block,
		address,
		address_suffix,
		timestamp,
	)
)

func TestBlocksResponse(t *testing.T) {

}
