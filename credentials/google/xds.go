package google

import (
	"context"
	"net"

	"google.golang.org/grpc/credentials"
)

const cfeClusterNamePrefix = "google_cfe_"
const cfeClusterResourceNamePrefix = "/envoy.config.cluster.v3.Cluster/google_cfe_"
const cfeClusterAuthorityName = "traffic-director-c2p.xds.googleapis.com"

type clusterTransportCreds struct {
	tls  credentials.TransportCredentials
	alts credentials.TransportCredentials
}

func newClusterTransportCreds(tls, alts credentials.TransportCredentials) *clusterTransportCreds {
	_ = "STUB: not implemented"
	return nil
}

func clusterName(ctx context.Context) string { _ = "STUB: not implemented"; return "" }

func isDirectPathCluster(ctx context.Context) bool { _ = "STUB: not implemented"; return false }

func (c *clusterTransportCreds) ClientHandshake(ctx context.Context, authority string, rawConn net.Conn) (net.Conn, credentials.AuthInfo, error) {
	_ = "STUB: not implemented"
	return *new(net.Conn), *new(credentials.AuthInfo), nil
}

func (c *clusterTransportCreds) ServerHandshake(conn net.Conn) (net.Conn, credentials.AuthInfo, error) {
	_ = "STUB: not implemented"
	return *new(net.Conn), *new(credentials.AuthInfo), nil
}

func (c *clusterTransportCreds) Info() credentials.ProtocolInfo {
	_ = "STUB: not implemented"
	return *new(credentials.ProtocolInfo)
}

func (c *clusterTransportCreds) Clone() credentials.TransportCredentials {
	_ = "STUB: not implemented"
	return *new(credentials.TransportCredentials)
}

func (c *clusterTransportCreds) OverrideServerName(s string) error {
	_ = "STUB: not implemented"
	return nil
}
