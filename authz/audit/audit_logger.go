package audit

import (
	"encoding/json"
	"sync"
)

type loggerBuilderRegistry struct {
	mu       sync.Mutex
	builders map[string]LoggerBuilder
}

var (
	registry = loggerBuilderRegistry{
		builders: make(map[string]LoggerBuilder),
	}
)

func RegisterLoggerBuilder(b LoggerBuilder) { _ = "STUB: not implemented"; return }

func GetLoggerBuilder(name string) LoggerBuilder {
	_ = "STUB: not implemented"
	return *new(LoggerBuilder)
}

type Event struct {
	FullMethodName string

	Principal string

	PolicyName string

	MatchedRule string

	Authorized bool
}

type LoggerConfig interface {
	loggerConfig()
}

type Logger interface {
	Log(*Event)
}

type LoggerBuilder interface {
	ParseLoggerConfig(config json.RawMessage) (LoggerConfig, error)

	Build(LoggerConfig) Logger

	Name() string
}
