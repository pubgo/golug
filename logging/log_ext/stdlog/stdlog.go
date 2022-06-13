package stdlog

import (
	"io"
	"log"

	"github.com/pubgo/dix"
	"github.com/pubgo/x/byteutil"
	"go.uber.org/zap"

	"github.com/pubgo/lava/logging"
)

// 替换std默认log
func init() {
	dix.Register(func(logger *logging.Logger) *logging.ExtLog {
		var stdLog = log.Default()
		// 接管系统默认log
		*stdLog = *zap.NewStdLog(logging.Component("std").L())
		return new(logging.ExtLog)
	})
}

var _ io.Writer = (*std)(nil)

type std struct {
	l *zap.Logger
}

func (s std) Write(p []byte) (n int, err error) {
	s.l.Info(byteutil.ToStr(p))
	return len(p), err
}
