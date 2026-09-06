package stdout

import (
	"encoding/json"
	"log"
	"os"

	"google.golang.org/grpc/authz/audit"
	"google.golang.org/grpc/grpclog"
)

var grpcLogger = grpclog.Component("authz-audit")

const Name = "stdout_logger"

func init() {
	audit.RegisterLoggerBuilder(&loggerBuilder{
		goLogger: log.New(os.Stdout, "", 0),
	})
}

type event struct {
	FullMethodName string `json:"rpc_method"`
	Principal      string `json:"principal"`
	PolicyName     string `json:"policy_name"`
	MatchedRule    string `json:"matched_rule"`
	Authorized     bool   `json:"authorized"`
	Timestamp      string `json:"timestamp"`
}

type logger struct {
	goLogger *log.Logger
}

func (l *logger) Log(event *audit.Event) { _ = "STUB: not implemented"; return }

type loggerConfig struct {
	audit.LoggerConfig
}

type loggerBuilder struct {
	goLogger *log.Logger
}

func (loggerBuilder) Name() string { _ = "STUB: not implemented"; return "" }

func (lb *loggerBuilder) Build(audit.LoggerConfig) audit.Logger {
	_ = "STUB: not implemented"
	return *new(audit.Logger)
}

func (*loggerBuilder) ParseLoggerConfig(config json.RawMessage) (audit.LoggerConfig, error) {
	_ = "STUB: not implemented"
	return *new(audit.LoggerConfig), nil
}

func convertEvent(auditEvent *audit.Event) *event { _ = "STUB: not implemented"; return nil }
