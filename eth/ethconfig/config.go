package ethconfig

import (
	"myethereum/eth/gasprice"
	"myethereum/params"
	"time"
)

var FullNodeGPO = gasprice.Config{
	Blocks:           20,
	Percentile:       60,
	MaxHeaderHistory: 1024,
	MaxBlockHistory:  1024,
	MaxPrice:         gasprice.DefaultMaxPrice,
	IgnorePrice:      gasprice.DefaultIngorePrice,
}

var LightClientGPO = gasprice.Config{
	Blocks:           2,
	Percentile:       60,
	MaxHeaderHistory: 300,
	MaxBlockHistory:  5,
	MaxPrice:         gasprice.DefaultMaxPrice,
	IgnorePrice:      gasprice.DefaultIngorePrice,
}

type Config struct {
	NetworkId         uint64
	EthDiscoveryURLs  []string
	SnapDiscoveryURLs []string
	NoPruning         bool
	NoPrefetch        bool
	TxLookupLimit     uint64 `toml:",omitempty"`
	// Light client options
	LightServ               int      `toml:",omitempty"` // Maximum percentage of time allowed for serving LES requests
	LightIngress            int      `toml:",omitempty"` // Incoming bandwidth limit for light servers
	LightEgress             int      `toml:",omitempty"` // Outgoing bandwidth limit for light servers
	LightPeers              int      `toml:",omitempty"` // Maximum number of LES client peers
	LightNoPrune            bool     `toml:",omitempty"` // Whether to disable light chain pruning
	LightNoSyncServe        bool     `toml:",omitempty"` // Whether to serve light clients before syncing
	SyncFromCheckpoint      bool     `toml:",omitempty"` // Whether to sync the header chain from the configured checkpoint
	UltraLightServers       []string `toml:",omitempty"` // List of trusted ultra light servers
	UltraLightFraction      int      `toml:",omitempty"` // Percentage of trusted servers to accept an announcement
	UltraLightOnlyAnnounce  bool     `toml:",omitempty"` // Whether to only announce headers, or also serve them
	SkipBcVersionCheck      bool     `toml:"-"`
	DatabaseHandles         int      `toml:"-"`
	DatabaseCache           int
	DatabaseFreezer         string
	TrieCleanCache          int
	TrieCleanCacheJournal   string        `toml:",omitempty"` // Disk journal directory for trie cache to survive node restarts
	TrieCleanCacheRejournal time.Duration `toml:",omitempty"` // Time interval to regenerate the journal for clean cache
	TrieDirtyCache          int
	TrieTimeout             time.Duration
	SnapshotCache           int
	Preimages               bool
	FilterLogCacheSize      int
	// Enables tracking of SHA3 preimages in the VM
	EnablePreimageRecording bool

	// Miscellaneous options
	DocRoot string `toml:"-"`
	// RPCGasCap is the global gas cap for eth-call variants.
	RPCGasCap uint64
	// RPCEVMTimeout is the global timeout for eth-call.
	RPCEVMTimeout time.Duration
	// RPCTxFeeCap is the global transaction fee(price * gaslimit) cap for
	// send-transaction variants. The unit is ether.
	RPCTxFeeCap float64
	// Checkpoint is a hardcoded checkpoint which can be nil.
	Checkpoint *params.TrustedCheckpoint `toml:",omitempty"`

	// CheckpointOracle is the configuration for checkpoint oracle.
	CheckpointOracle *params.CheckpointOracleConfig `toml:",omitempty"`
	// OverrideShanghai (TODO: remove after the fork)
	OverrideShanghai *uint64 `toml:",omitempty"`
}

var Defaults = Config{

	NetworkId:               1,
	TxLookupLimit:           2350000,
	LightPeers:              100,
	UltraLightFraction:      75,
	DatabaseCache:           512,
	TrieCleanCache:          154,
	TrieCleanCacheJournal:   "triecache",
	TrieCleanCacheRejournal: 60 * time.Minute,
	TrieDirtyCache:          256,
	TrieTimeout:             60 * time.Minute,
	SnapshotCache:           102,
	FilterLogCacheSize:      32,

	RPCGasCap:     50000000,
	RPCEVMTimeout: 5 * time.Second,
	RPCTxFeeCap:   1, // 1 ether
}
