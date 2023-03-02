package node

import (
	"fmt"
	"github.com/prometheus/tsdb/fileutil"
	"myethereum/log"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

type Node struct {
	config        *Config
	log           log.Logger
	keyDir     string
	keyDirTemp bool
	dirLock    fileutil.Releaser
	stop          chan struct{}
	startStopLock sync.Mutex
	state         int
	lock          sync.Mutex
	lifecycles    []Lifecycle
}

const (
	initializingState = iota
	runningState
	closedState
)

func New(conf *Config) (*Node, error) {
	confCopy := *conf
	conf = &confCopy
	if conf.DataDir != "" {
		absdatadir, err := filepath.Abs(conf.DataDir)
		if err != nil {
			return nil, err
		}
		conf.DataDir = absdatadir
	}
	if conf.Logger == nil {
		conf.Logger = log.New()
	}
	if strings.ContainsAny(conf.Name, `/\`) {
		return nil, fmt.Errorf("node name cannot contain path separators")
	}
	if conf.Name == datadirDefaultKeyStore {
		return nil, fmt.Errorf("node name cannot be 'keystore'")
	}
	if strings.HasSuffix(conf.Name, ".ipc") {
		return nil, fmt.Errorf("node name cannot end in '.ipc'")
	}
	node := &Node{
		config: conf,
		log:    conf.Logger,
		stop:   make(chan struct{}),
	}

	//TODO register Api

	if err := node.openDataDir(); err != nil {
		return nil, err
	}
	keyDir, isEpem, err := getKeyStoreDir(conf)
	if err != nil {
		return nil, err
	}
	node.keyDir=keyDir
	node.keyDirTemp =isEpem
	//TODO set the Account Manager
	node.
}

func (n *Node) openDataDir() error {
	if n.config.DataDir == "" {
		return nil
	}
	instDir := filepath.Join(n.config.DataDir, n.config.Name)
	if err := os.MkdirAll(instDir, 0700); err != nil {
		return err
	}
	release, _, err := fileutil.Flock(filepath.Join(instDir, "LOCK"))
	if err != nil {
		return convertFileLockError(err)
	}
	n.dirLock = release
	return nil
}
