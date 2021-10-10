package main

import (
	"github.com/pubgo/x/q"
	"github.com/pubgo/xerror"

	"github.com/pubgo/lug/pkg/modutil"
)

func main() {
	defer xerror.RespExit()
	q.Q(modutil.GoModPath())
	q.Q(modutil.LoadMod())
}
