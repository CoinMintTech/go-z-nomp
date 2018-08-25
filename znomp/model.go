package znomp

import (
	"time"
)

// Block representation of z block.
type Block struct {
	Coin   string
	Block  uint
	Field1 string
	Field2 string
	Field3 uint
	// Miner name of the miner.
	Miner string
	// Timestamp point in time when block was mined.
	Timestamp time.Time
}
