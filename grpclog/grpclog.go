package grpclog

func init() {
	SetLoggerV2(newLoggerV2())
}

func V(l int) bool { _ = "STUB: not implemented"; return false }

func Info(args ...any) { _ = "STUB: not implemented"; return }

func Infof(format string, args ...any) { _ = "STUB: not implemented"; return }

func Infoln(args ...any) { _ = "STUB: not implemented"; return }

func Warning(args ...any) { _ = "STUB: not implemented"; return }

func Warningf(format string, args ...any) { _ = "STUB: not implemented"; return }

func Warningln(args ...any) { _ = "STUB: not implemented"; return }

func Error(args ...any) { _ = "STUB: not implemented"; return }

func Errorf(format string, args ...any) { _ = "STUB: not implemented"; return }

func Errorln(args ...any) { _ = "STUB: not implemented"; return }

func Fatal(args ...any) { _ = "STUB: not implemented"; return }

func Fatalf(format string, args ...any) { _ = "STUB: not implemented"; return }

func Fatalln(args ...any) { _ = "STUB: not implemented"; return }

func Print(args ...any) { _ = "STUB: not implemented"; return }

func Printf(format string, args ...any) { _ = "STUB: not implemented"; return }

func Println(args ...any) { _ = "STUB: not implemented"; return }

func InfoDepth(depth int, args ...any) { _ = "STUB: not implemented"; return }

func WarningDepth(depth int, args ...any) { _ = "STUB: not implemented"; return }

func ErrorDepth(depth int, args ...any) { _ = "STUB: not implemented"; return }

func FatalDepth(depth int, args ...any) { _ = "STUB: not implemented"; return }
