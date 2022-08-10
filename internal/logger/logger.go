package logger

import (
	"github.com/go-logr/logr"
)

// Logger is the default logger object.
var Logger AdaptedLogger = NewDefault()

// New create a new AdaptedLogger
//
// The cfg is optional, we use cfg.Level only.
// NOTE: there are 4 ways to control logger.Level(with priority decend):
// 1. value of "DefaultLogLevel"
// 2. value of "os.Getenv(CABIN_LOG_LEVEL)" eg: "set CABIN_LOG_LEVEL=0"
// 3. value of Config.Level
// 4. default value DebugLevel(-1)
func New(cfg *LogConfig, options ...Option) AdaptedLogger {
	return newPackZap(cfg, options...)
}

// NewDefault create a defult Logger
func NewDefault() AdaptedLogger {
	return New(nil)
}

// NewFromLogr new Logger from logr.Logger
func NewFromLogr(log logr.Logger) AdaptedLogger {
	return newPackLogr(log)
}

// AdaptedLogger is the interface that adapt zap.logger
type AdaptedLogger interface {
	// WithValues returns a new Logger with additional key/value pairs.
	WithValues(keysAndValues ...interface{}) AdaptedLogger

	// WithName returns a new Logger with the specified name appended.
	WithName(name string) AdaptedLogger

	// WithLevel returns a new Logger with the specified level filter.
	WithLevel(level Level) AdaptedLogger

	// WithOptions clones the current Logger, applies the supplied Options, and
	// returns the resulting Logger. It's safe to use concurrently.
	WithOptions(opts ...Option) AdaptedLogger

	// PutError write log with error
	PutError(err error, msg string, keysAndValues ...interface{})

	// Debug uses fmt.Sprint to construct and log a message.
	Debug(args ...interface{})

	// Info uses fmt.Sprint to construct and log a message.
	Info(args ...interface{})

	// Warn uses fmt.Sprint to construct and log a message.
	Warn(args ...interface{})

	// Error uses fmt.Sprint to construct and log a message.
	Error(args ...interface{})

	// DPanic uses fmt.Sprint to construct and log a message. In development, the
	// logger then panics. (See DPanicLevel for details.)
	DPanic(args ...interface{})

	// Panic uses fmt.Sprint to construct and log a message, then panicl.
	Panic(args ...interface{})

	// Fatal uses fmt.Sprint to construct and log a message, then calls ol.Exit.
	Fatal(args ...interface{})

	// Debugf uses fmt.Sprintf to log a templated message.
	Debugf(template string, args ...interface{})

	// Infof uses fmt.Sprintf to log a templated message.
	Infof(template string, args ...interface{})

	// Warnf uses fmt.Sprintf to log a templated message.
	Warnf(template string, args ...interface{})

	// Errorf uses fmt.Sprintf to log a templated message.
	Errorf(template string, args ...interface{})

	// DPanicf uses fmt.Sprintf to log a templated message. In development, the
	// logger then panics. (See DPanicLevel for details.)
	DPanicf(template string, args ...interface{})

	// Panicf uses fmt.Sprintf to log a templated message, then panicl.
	Panicf(template string, args ...interface{})

	// Fatalf uses fmt.Sprintf to log a templated message, then calls ol.Exit.
	Fatalf(template string, args ...interface{})

	// Debugw logs a message with some additional context. The variadic key-value
	// pairs are treated as they are in With.
	//
	// When debug-level logging is disabled, this is much faster than
	//  l.With(keysAndValues).Debug(msg)
	Debugw(msg string, keysAndValues ...interface{})

	// Infow logs a message with some additional context. The variadic key-value
	// pairs are treated as they are in With.
	Infow(msg string, keysAndValues ...interface{})

	// Warnw logs a message with some additional context. The variadic key-value
	// pairs are treated as they are in With.
	Warnw(msg string, keysAndValues ...interface{})

	// Errorw logs a message with some additional context. The variadic key-value
	// pairs are treated as they are in With.
	Errorw(msg string, keysAndValues ...interface{})

	// DPanicw logs a message with some additional context. In development, the
	// logger then panics. (See DPanicLevel for details.) The variadic key-value
	// pairs are treated as they are in With.
	DPanicw(msg string, keysAndValues ...interface{})

	// Panicw logs a message with some additional context, then panicl. The
	// variadic key-value pairs are treated as they are in With.
	Panicw(msg string, keysAndValues ...interface{})

	// Fatalw logs a message with some additional context, then calls ol.Exit. The
	// variadic key-value pairs are treated as they are in With.
	Fatalw(msg string, keysAndValues ...interface{})

	// Sync flushes any buffered log entries.
	Sync() error
}
