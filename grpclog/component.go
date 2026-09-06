package grpclog

type componentData struct {
	name string
}

var cache = map[string]*componentData{}

func (c *componentData) InfoDepth(depth int, args ...any) { _ = "STUB: not implemented"; return }

func (c *componentData) WarningDepth(depth int, args ...any) { _ = "STUB: not implemented"; return }

func (c *componentData) ErrorDepth(depth int, args ...any) { _ = "STUB: not implemented"; return }

func (c *componentData) FatalDepth(depth int, args ...any) { _ = "STUB: not implemented"; return }

func (c *componentData) Info(args ...any) { _ = "STUB: not implemented"; return }

func (c *componentData) Warning(args ...any) { _ = "STUB: not implemented"; return }

func (c *componentData) Error(args ...any) { _ = "STUB: not implemented"; return }

func (c *componentData) Fatal(args ...any) { _ = "STUB: not implemented"; return }

func (c *componentData) Infof(format string, args ...any) { _ = "STUB: not implemented"; return }

func (c *componentData) Warningf(format string, args ...any) { _ = "STUB: not implemented"; return }

func (c *componentData) Errorf(format string, args ...any) { _ = "STUB: not implemented"; return }

func (c *componentData) Fatalf(format string, args ...any) { _ = "STUB: not implemented"; return }

func (c *componentData) Infoln(args ...any) { _ = "STUB: not implemented"; return }

func (c *componentData) Warningln(args ...any) { _ = "STUB: not implemented"; return }

func (c *componentData) Errorln(args ...any) { _ = "STUB: not implemented"; return }

func (c *componentData) Fatalln(args ...any) { _ = "STUB: not implemented"; return }

func (c *componentData) V(l int) bool { _ = "STUB: not implemented"; return false }

func Component(componentName string) DepthLoggerV2 {
	_ = "STUB: not implemented"
	return *new(DepthLoggerV2)
}
