package keiheila

import (
	"fmt"
	"net"
	"time"

	"github.com/lonelyevil/khl"
	"go.uber.org/zap"
)

type loggerLevel int

const (
	traceLevel = loggerLevel(iota)
	debugLevel = loggerLevel(iota)
	infoLevel
	warnLevel
	errorLevel
	fatalLevel
)

type LoggerAdapter struct {
	l *zap.Logger
}

func NewLogger(l *zap.Logger) *LoggerAdapter {
	return &LoggerAdapter{l: l}
}

func (l LoggerAdapter) Trace() khl.Entry {
	return &EntryAdapter{l.l, traceLevel}
}

func (l LoggerAdapter) Debug() khl.Entry {
	return &EntryAdapter{l.l, debugLevel}
}

func (l LoggerAdapter) Info() khl.Entry {
	return &EntryAdapter{l.l, infoLevel}
}

func (l LoggerAdapter) Warn() khl.Entry {
	return &EntryAdapter{l.l, warnLevel}
}

func (l LoggerAdapter) Error() khl.Entry {
	return &EntryAdapter{l.l, errorLevel}
}

func (l LoggerAdapter) Fatal() khl.Entry {
	return &EntryAdapter{l.l, fatalLevel}
}

type EntryAdapter struct {
	e *zap.Logger
	t loggerLevel
}

func (e *EntryAdapter) Bool(key string, b bool) khl.Entry {
	e.e = e.e.With(zap.Bool(key, b))
	return e
}

func (e *EntryAdapter) Bytes(key string, val []byte) khl.Entry {
	e.e = e.e.With(zap.Binary(key, val))
	return e
}

func (e *EntryAdapter) Caller(depth int) khl.Entry {
	e.e = e.e.WithOptions(zap.AddCallerSkip(depth))
	return e
}

func (e *EntryAdapter) Dur(key string, d time.Duration) khl.Entry {
	e.e = e.e.With(zap.Duration(key, d))
	return e
}

func (e *EntryAdapter) Err(key string, err error) khl.Entry {
	e.e = e.e.With(zap.NamedError(key, err))
	return e
}

func (e *EntryAdapter) Float64(key string, f float64) khl.Entry {
	e.e = e.e.With(zap.Float64(key, f))
	return e
}

func (e *EntryAdapter) IPAddr(key string, ip net.IP) khl.Entry {
	e.e = e.e.With(zap.String(key, ip.String()))
	return e
}

func (e *EntryAdapter) Int(key string, i int) khl.Entry {
	e.e = e.e.With(zap.Int(key, i))
	return e
}

func (e *EntryAdapter) Int64(key string, i int64) khl.Entry {
	e.e = e.e.With(zap.Int64(key, i))
	return e
}

func (e *EntryAdapter) Interface(key string, i interface{}) khl.Entry {
	e.e = e.e.With(zap.Any(key, i))
	return e
}

func (e *EntryAdapter) Msg(msg string) {
	switch e.t {
	case traceLevel, debugLevel:
		e.e.Debug(msg)
	case infoLevel:
		e.e.Info(msg)
	case warnLevel:
		e.e.Warn(msg)
	case errorLevel:
		e.e.Error(msg)
	case fatalLevel:
		e.e.Fatal(msg)
	}
}

func (e *EntryAdapter) Msgf(f string, i ...interface{}) {
	e.Msg(fmt.Sprintf(f, i...))
}

func (e *EntryAdapter) Str(key string, s string) khl.Entry {
	e.e = e.e.With(zap.String(key, s))
	return e
}

func (e *EntryAdapter) Strs(key string, s []string) khl.Entry {
	e.e = e.e.With(zap.Strings(key, s))
	return e
}

func (e *EntryAdapter) Time(key string, t time.Time) khl.Entry {
	e.e = e.e.With(zap.Time(key, t))
	return e
}
