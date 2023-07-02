package definition

// This log interface is not NFC required.
// But we have used this to log some errors when authorization request failed.

type LogLevel int8

const (
	LogLevelFalat LogLevel = iota
	LogLevelError
	LogLevelInfo
	LogLevelTrace
)

// Logger interface
type Logger interface {
	Falatf(format string, args ...any)
	Errorf(format string, args ...any)
	Infof(format string, args ...any)
}

// TraceableLogger interface
type TraceableLogger interface {
	Tracef(format string, args ...any)
}
