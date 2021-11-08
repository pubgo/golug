package logz

import (
	"github.com/pubgo/lava/logger"
	"github.com/pubgo/xerror"
	"go.uber.org/zap"
)

func New(name string) *Logger {
	if name == "" {
		panic("[name] should not be null")
	}
	return &Logger{name: name}
}

type Logger struct {
	name string
}

func (t *Logger) With(args ...zap.Field) *zap.Logger {
	return getName(t.name).With(args...)
}

func (t *Logger) StepAndThrow(msg string, fn func() error, fields ...zap.Field) {
	var log = t.Depth(1)
	log = log.With(fields...)

	log.Info(msg)

	var err error
	xerror.TryWith(&err, func() { err = fn() })

	if err == nil {
		log.Info(msg + " ok")
		return
	}

	log.Error(msg+" error", logger.WithErr(err)...)
	panic(err)
}

func (t *Logger) Step(msg string, fn func() error, fields ...zap.Field) {
	var log = t.Depth(1)
	log = log.With(fields...)

	log.Info(msg)

	var err error
	xerror.TryWith(&err, func() { err = fn() })

	if err == nil {
		log.Info(msg + " ok")
		return
	}

	log.Error(msg+" error", logger.WithErr(err)...)
}

func (t *Logger) Logs(msg string, fn func() error, fields ...zap.Field) {
	var log = t.Depth(1)
	var err error
	xerror.TryWith(&err, func() { err = fn() })

	if err == nil {
		log.Info(msg, fields...)
		return
	}

	log.Error(msg, logger.WithErr(err, fields...)...)
}

func (t *Logger) LogAndThrow(msg string, fn func() error, fields ...zap.Field) {
	var log = t.Depth(1)
	var err error
	xerror.TryWith(&err, func() { err = fn() })

	if err == nil {
		log.Info(msg, fields...)
		return
	}

	log.Error(msg, logger.WithErr(err, fields...)...)
	panic(err)
}

func (t *Logger) WithErr(err error, fields ...zap.Field) *zap.Logger {
	if err == nil {
		return Discard
	}

	return t.With(logger.WithErr(err, fields...)...)
}

func (t *Logger) Depth(depth ...int) *zap.Logger {
	if len(depth) > 0 {
		return getName(t.name).WithOptions(zap.AddCallerSkip(depth[0]))
	}
	return getName(t.name)
}

func (t *Logger) DepthS(depth ...int) *zap.SugaredLogger {
	return t.Depth(depth...).Sugar()
}

func (t *Logger) Infof(template string, args ...interface{}) {
	t.DepthS(1).Infof(template, args...)
}

func (t *Logger) Info(args ...interface{}) {
	t.DepthS(1).Info(args...)
}

func (t *Logger) Infow(msg string, keysAndValues ...interface{}) {
	t.DepthS(1).Infow(msg, keysAndValues...)
}

func (t *Logger) Errorf(template string, args ...interface{}) {
	t.DepthS(1).Errorf(template, args...)
}

func (t *Logger) Error(args ...interface{}) {
	t.DepthS(1).Error(args...)
}

func (t *Logger) Errorw(msg string, keysAndValues ...interface{}) {
	t.DepthS(1).Errorw(msg, keysAndValues...)
}

func (t *Logger) Warnf(template string, args ...interface{}) {
	t.DepthS(1).Warnf(template, args...)
}

func (t *Logger) Warn(args ...interface{}) {
	t.DepthS(1).Warn(args...)
}

func (t *Logger) Warnw(msg string, keysAndValues ...interface{}) {
	t.DepthS(1).Warnw(msg, keysAndValues...)
}

func (t *Logger) TryWith(fn func()) *zap.SugaredLogger {
	var err error
	xerror.TryWith(&err, fn)
	if err == nil {
		return Discard.Sugar()
	}

	return debugLog.Named(t.name).With(logger.WithErr(err, logger.FuncStack(fn))...).Sugar()
}

func getName(name string) *zap.Logger {
	if val, ok := loggerMap.Load(name); ok {
		return val.(*zap.Logger)
	}

	var l = debugLog.Named(name)
	loggerMap.LoadOrStore(name, l)
	return l
}
