package bbolt

import (
	"io/fs"
	"path/filepath"
	"time"

	"github.com/pubgo/x/merge"
	"github.com/pubgo/x/pathutil"
	"github.com/pubgo/xerror"
	bolt "go.etcd.io/bbolt"

	"github.com/pubgo/lug/config"
)

const Name = "bbolt"

var cfgMap = make(map[string]*Cfg)

type Cfg struct {
	FileMode        fs.FileMode       `json:"file_mode"`
	Timeout         time.Duration     `json:"timeout"`
	NoGrowSync      bool              `json:"no_grow_sync"`
	NoFreelistSync  bool              `json:"no_freelist_sync"`
	FreelistType    bolt.FreelistType `json:"freelist_type"`
	ReadOnly        bool              `json:"read_only"`
	MmapFlags       int               `json:"mmap_flags"`
	InitialMmapSize int               `json:"initial_mmap_size"`
	PageSize        int               `json:"page_size"`
	NoSync          bool              `json:"no_sync"`
	Path            string            `json:"path"`

	db *DB
}

func (t *Cfg) BuildOpts() *bolt.Options {
	var options = bolt.DefaultOptions
	options.Timeout = time.Second * 2
	xerror.Panic(merge.CopyStruct(options, t))
	return options
}

func (t *Cfg) Build() (gErr error) {
	defer xerror.RespErr(&gErr)

	var opts = t.BuildOpts()
	var path = filepath.Join(config.Home, t.Path)
	xerror.Panic(pathutil.IsNotExistMkDir(filepath.Dir(path)))

	db, err := bolt.Open(path, t.FileMode, opts)
	xerror.Panic(err)
	t.db = &DB{db: db}

	return
}

func DefaultCfg() *Cfg {
	return &Cfg{
		Path:     "./db/superloop",
		FileMode: 0666,
		Timeout:  time.Second,
	}
}
