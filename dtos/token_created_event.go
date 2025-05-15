package dtos

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type TokenCreatedEvent struct {
	Owner    common.Address
	TokenID  *big.Int
	TokenURI string
}
