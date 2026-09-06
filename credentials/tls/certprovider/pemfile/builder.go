package pemfile

import (
	"encoding/json"
	"time"

	"google.golang.org/grpc/credentials/tls/certprovider"
)

const (
	PluginName             = "file_watcher"
	defaultRefreshInterval = 10 * time.Minute
)

func init() {
	certprovider.Register(&pluginBuilder{})
}

type pluginBuilder struct{}

func (p *pluginBuilder) ParseConfig(c any) (*certprovider.BuildableConfig, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *pluginBuilder) Name() string { _ = "STUB: not implemented"; return "" }

func pluginConfigFromJSON(jd json.RawMessage) (Options, error) {
	_ = "STUB: not implemented"
	return *new(Options), nil
}
