package core

import (
	"myethereum/params"
)

type Genesis struct {
	Config *params.ChainConfig `json:"config"`
	Nonce  uint64              `json:"nonce"`
}
