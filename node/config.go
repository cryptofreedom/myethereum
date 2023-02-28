package node

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
}
