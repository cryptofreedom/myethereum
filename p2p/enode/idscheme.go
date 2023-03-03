package enode

import (
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rlp"
	"golang.org/x/crypto/sha3"
	"io"
	"myethereum/common/math"
	"myethereum/p2p/enr"
)

var ValidSchemes = enr.SchemeMap{
	"v4": V4ID{},
}

var ValidSchemesForTesting = enr.SchemeMap{
	"v4":   V4ID{},
	"null": NullID{},
}

type V4ID struct{}

func SignV4(r *enr.Record, privkey *ecdsa.PrivateKey) error {
	cpy := *r
	cpy.Set(enr.ID("v4"))
	cpy.Set(Secp256k1(privkey.PublicKey))
	h := sha3.NewLegacyKeccak256()
	rlp.Encode(h, cpy.AppendElements(nil))
	sig, err := crypto.Sign(h.Sum(nil), privkey)
	if err != nil {
		return err
	}
	sig = sig[:len(sig)-1]
	if err = cpy.SetSig(V4ID{}, sig); err == nil {
		*r = cpy
	}
	return err
}
func (V4ID) Verify(r *enr.Record, sig []byte) error {
	var entry s256raw
	if err := r.Load(&entry); err != nil {
		return err
	} else if len(entry) != 33 {
		return fmt.Errorf("invalid public key")
	}
	h := sha3.NewLegacyKeccak256()
	rlp.Encode(h, r.AppendElements(nil))
	if !crypto.VerifySignature(entry, h.Sum(nil), sig) {
		return enr.ErrInvalidSig
	}
	return nil
}

func (V4ID) NodeAddr(r *enr.Record) []byte {
	var pubKey Secp256k1
	err := r.Load(&pubKey)
	if err != nil {
		return nil
	}
	buf := make([]byte, 64)
	math.ReadBits(pubKey.X, buf[:32])
	math.ReadBits(pubKey.Y, buf[32:])
	return crypto.Keccak256(buf)
}

type Secp256k1 ecdsa.PublicKey

func (v Secp256k1) ENRKey() string { return "secp256k1" }

// EncodeRLP implements rlp.Encoder.
func (v Secp256k1) EncodeRLP(w io.Writer) error {
	return rlp.Encode(w, crypto.CompressPubkey((*ecdsa.PublicKey)(&v)))
}

// DecodeRLP implements rlp.Decoder.
func (v *Secp256k1) DecodeRLP(s *rlp.Stream) error {
	buf, err := s.Bytes()
	if err != nil {
		return err
	}
	pk, err := crypto.DecompressPubkey(buf)
	if err != nil {
		return err
	}
	*v = (Secp256k1)(*pk)
	return nil
}

type s256raw []byte

func (s256raw) ENRKey() string { return "secp256k1" }

type v4CompatID struct {
	V4ID
}

type NullID struct {
}

func (NullID) Verify(r *enr.Record, sig []byte) error {
	return nil
}

func (NullID) NodeAddr(r *enr.Record) []byte {
	var id ID
	r.Load(enr.WithEntry("nulladdr", &id))
	return id[:]
}

func SignNull(r *enr.Record, id ID) *Node {
	r.Set(enr.ID("null"))
	r.Set(enr.WithEntry("nulladdr", id))
	if err := r.SetSig(NullID{}, nil); err != nil {
		panic(err)
	}
	return &Node{r: *r, id: id}
}
