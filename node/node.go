package node

import (
	"github.com/prometheus/tsdb/fileutil"
	"log"
	"path/filepath"
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

const (
	initializingState = iota
	runningState
	closedState
)

func New(conf *Config) (*Node, error) {
	confCopy := *conf
	conf = &confCopy
	if conf.DataDir != ""{
		absdatadir,err := filepath.Abs(conf.DataDir)
		if err!=nil{
			return nil,err
		}
		conf.DataDir=absdatadir
	}
	if conf.
}
