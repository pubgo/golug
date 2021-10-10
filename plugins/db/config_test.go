package db

import (
	"fmt"
	"testing"

	"github.com/pubgo/lug/plugins/db/sqlite"
	"github.com/pubgo/xerror"
)

func TestConfig(t *testing.T) {
	defer xerror.RespTest(t)

	var cfg = DefaultCfg()
	cfg.Driver = sqlite.Name
	cfg.Source = "./sqlite.db"

	db := cfg.Build()
	fmt.Println(db.Query("select * from db"))
	fmt.Println(db.Query("select * from db where Field1=?", 1))
}
