package node

import (
	"github.com/prometheus/tsdb/fileutil"
	"log"
	"sync"
)

type Node struct {
	config        *Config
	log           log.Logger
	keyDir        string
	keuyDirTemp   bool
	dirLock       fileutil.Releaser
	stop          chan struct{}
	startStopLock sync.Mutex
	state         int
	lock          sync.Mutex
	lifecycles    []Lifecycle
}
