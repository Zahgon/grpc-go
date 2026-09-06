package glogger

import (
	"google.golang.org/grpc/grpclog"
)

const d = 2

func init() {
	grpclog.SetLoggerV2(&glogger{})
}

type glogger struct{}

func (g *glogger) Info(args ...any) { _ = "STUB: not implemented"; return }

func (g *glogger) Infoln(args ...any) { _ = "STUB: not implemented"; return }

func (g *glogger) Infof(format string, args ...any) { _ = "STUB: not implemented"; return }

func (g *glogger) InfoDepth(depth int, args ...any) { _ = "STUB: not implemented"; return }

func (g *glogger) Warning(args ...any) { _ = "STUB: not implemented"; return }

func (g *glogger) Warningln(args ...any) { _ = "STUB: not implemented"; return }

func (g *glogger) Warningf(format string, args ...any) { _ = "STUB: not implemented"; return }

func (g *glogger) WarningDepth(depth int, args ...any) { _ = "STUB: not implemented"; return }

func (g *glogger) Error(args ...any) { _ = "STUB: not implemented"; return }

func (g *glogger) Errorln(args ...any) { _ = "STUB: not implemented"; return }

func (g *glogger) Errorf(format string, args ...any) { _ = "STUB: not implemented"; return }

func (g *glogger) ErrorDepth(depth int, args ...any) { _ = "STUB: not implemented"; return }

func (g *glogger) Fatal(args ...any) { _ = "STUB: not implemented"; return }

func (g *glogger) Fatalln(args ...any) { _ = "STUB: not implemented"; return }

func (g *glogger) Fatalf(format string, args ...any) { _ = "STUB: not implemented"; return }

func (g *glogger) FatalDepth(depth int, args ...any) { _ = "STUB: not implemented"; return }

func (g *glogger) V(l int) bool { _ = "STUB: not implemented"; return false }
