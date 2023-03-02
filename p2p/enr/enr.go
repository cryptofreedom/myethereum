package enr

import (
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/rlp"
	"sort"
)

const SizeLimit = 300

var (
	ErrInvalidSig     = errors.New("invalid signature on node record")
	errNotSorted      = errors.New("record key/value pairs are not sorted by key")
	errDuplicateKey   = errors.New("record contains duplicate key")
	errIncompletePair = errors.New("record contains incomplete k/v pair")
	errIncompleteList = errors.New("record contains less than two list elements")
	errTooBig         = fmt.Errorf("record bigger than %d bytes", SizeLimit)
	errEncodeUnsigned = errors.New("can't encode unsigned record")
	errNotFound       = errors.New("no such key in record")
)

type IdentityScheme interface {
	Verify(r *Record, sig []byte) error
	NodeAddr(r *Record) []byte
}
type SchemeMap map[string]IdentityScheme

func (m SchemeMap) Verify(r *Record, sig []byte) error {
	s := m[r.IdentityScheme()]
	if s == nil {
		return ErrInvalidSig
	}
	return s.Verify(r, sig)
}

type Record struct {
	seq       uint64
	signature []byte
	raw       []byte
	pairs     []pair
}

type pair struct {
	k string
	v rlp.RawValue
}

func (r *Record) Size() uint64 {
	if r.raw != nil {
		return uint64(len(r.raw))
	}
	return computeSize(r)
}

func (r *Record) IdentityScheme() string {
	var id ID
	r.Load(&id)
	return string(id)
}

func (r *Record) Load(e Entry) error {
	i := sort.Search(len(r.pairs), func(i int) bool {
		return r.pairs[i].k >= e.ENRKey()
	})
	if i < len(r.pairs) && r.pairs[i].k == e.ENRKey() {
		if err := rlp.DecodeBytes(r.pairs[i].v, e); err != nil {
			return &KeyError{Key: e.ENRKey(), Err: err}
		}
		return nil
	}
	return &KeyError{Key: e.ENRKey(), Err: errNotFound}
}

func computeSize(r *Record) uint64 {
	size := uint64(rlp.IntSize(r.seq))
	size += rlp.BytesSize(r.signature)
	for _, p := range r.pairs {
		size += rlp.StringSize(p.k)
		size += uint64(len(p.v))
	}
	return rlp.ListSize(size)
}
