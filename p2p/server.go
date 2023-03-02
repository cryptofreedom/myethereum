package p2p

import (
	"crypto/ecdsa"
	"errors"
	"github.com/ethereum/go-ethereum/p2p/enode"
	"time"
)

const (
	defaultDialTimeout     = 15 * time.Second
	discmixTimeout         = 5 * time.Second
	defaultMaxPendingPeers = 50
	defaultDialRatio       = 3
	// This time limits inbound connection attempts per source IP.
	inboundThrottleTime = 30 * time.Second
	frameReadTimeout    = 30 * time.Second
	frameWriteTimeout   = 20 * time.Second
)

var errServerStopped = errors.New("server stopped")

type Config struct {
	PrivateKey      *ecdsa.PrivateKey `toml:"-"`
	MaxPeers        int
	MaxPendingPeers int `toml:",omitempty"`
	DialRatio       int `toml:",omitempty"`
	NoDiscovery     bool
	DiscoveryV5     bool   `toml:",omitempty"`
	Name            string `toml:"-"`
	BootstrapNodes  []*enode.Node
	Boot
}
