package enode

import (
	"errors"
	"github.com/ethereum/go-ethereum/p2p/enr"
)

var errMissingPrefix = errors.New("Missing 'enr:' prefix for base64-encoded record")

type Node struct {
	r enr.Record
}
