package enode

import (
	"errors"
	"myethereum/p2p/enr"
)

var errMissingPrefix = errors.New("Missing 'enr:' prefix for base64-encoded record")

type ID [32]byte
type Node struct {
	r  enr.Record
	id ID
}

func New(validSchemes enr.IdentityScheme, r *enr.Record) (*Node, error) {
	if err := r
}
