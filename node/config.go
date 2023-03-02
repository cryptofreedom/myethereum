package node

import (
	"myethereum/log"
	"os"
	"path/filepath"
)

const (
	datadirPrivateKey      = "nodekey"            // Path within the datadir to the node's private key
	datadirJWTKey          = "jwtsecret"          // Path within the datadir to the node's jwt secret
	datadirDefaultKeyStore = "keystore"           // Path within the datadir to the keystore
	datadirStaticNodes     = "static-nodes.json"  // Path within the datadir to the static node list
	datadirTrustedNodes    = "trusted-nodes.json" // Path within the datadir to the trusted node list
	datadirNodeDatabase    = "nodes"              // Path within the datadir to store the node infos
)

type Config struct {
	Name       string `toml:"-"`
	UserIndent string `toml:",omitempty"`
	Version    string `toml:"-"`
	DataDir    string
	//p2p config here
	KeyStoreDir           string `toml:",omitempty"`
	ExternalSigner        string `toml:",omitempty"`
	UseLightweightKDF     bool   `toml:",omitempty"`
	InsecureUnlockAllowed bool   `toml:",omitempty"`
	NoUSB                 bool   `toml:",omitempty"`
	USB                   bool   `toml:",omitempty"`
	SmartCardDaemonPath   string `toml:",omitempty"`
	IPCPath               string
	Logger                log.Logger `toml:",omitempty"`
}

func getKeyStoreDir(conf *Config) (string, bool, error) {
	keydir, err := conf.KeyDirConfig()
	if err != nil {
		return "", false, err
	}
	isEphemeral := false
	if keydir == "" {
		keydir, err = os.MkdirTemp(os.TempDir(), "go-ethereum-keystore")
		isEphemeral = true
	}
	if err != nil {
		return "", false, err
	}
	if err := os.MkdirAll(keydir, 0700); err != nil {
		return "", false, err
	}
	return keydir, isEphemeral, nil
}

func (c *Config) KeyDirConfig() (string, error) {
	var (
		keydir string
		err    error
	)
	switch {
	case filepath.IsAbs(c.KeyStoreDir):
		keydir = c.KeyStoreDir
	case c.DataDir != "":
		if c.KeyStoreDir == "" {
			keydir = filepath.Join(c.DataDir, datadirDefaultKeyStore)
		} else {
			keydir, err = filepath.Abs(c.KeyStoreDir)
		}
	case c.KeyStoreDir != "":
		keydir, err = filepath.Abs(c.KeyStoreDir)
	}
	return keydir, err
}
