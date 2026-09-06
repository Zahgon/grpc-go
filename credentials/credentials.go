package credentials

import (
	"context"
	"errors"
	"net"

	"google.golang.org/grpc/attributes"
	"google.golang.org/protobuf/proto"
)

type PerRPCCredentials interface {
	GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error)

	RequireTransportSecurity() bool
}

type SecurityLevel int

const (
	InvalidSecurityLevel SecurityLevel = iota

	NoSecurity

	IntegrityOnly

	PrivacyAndIntegrity
)

func (s SecurityLevel) String() string { _ = "STUB: not implemented"; return "" }

type CommonAuthInfo struct {
	SecurityLevel SecurityLevel
}

func (c CommonAuthInfo) GetCommonAuthInfo() CommonAuthInfo {
	_ = "STUB: not implemented"
	return *new(CommonAuthInfo)
}

type ProtocolInfo struct {
	ProtocolVersion string

	SecurityProtocol string

	SecurityVersion string

	ServerName string
}

type AuthInfo interface {
	AuthType() string
}

type AuthorityValidator interface {
	ValidateAuthority(authority string) error
}

var ErrConnDispatched = errors.New("credentials: rawConn is dispatched out of gRPC")

type TransportCredentials interface {
	ClientHandshake(context.Context, string, net.Conn) (net.Conn, AuthInfo, error)

	ServerHandshake(net.Conn) (net.Conn, AuthInfo, error)

	Info() ProtocolInfo

	Clone() TransportCredentials

	OverrideServerName(string) error
}

type Bundle interface {
	TransportCredentials() TransportCredentials

	PerRPCCredentials() PerRPCCredentials

	NewWithMode(mode string) (Bundle, error)
}

type RequestInfo struct {
	Method string

	AuthInfo AuthInfo
}

type requestInfoKey struct{}

func RequestInfoFromContext(ctx context.Context) (ri RequestInfo, ok bool) {
	_ = "STUB: not implemented"
	return *new(RequestInfo), false
}

func NewContextWithRequestInfo(ctx context.Context, ri RequestInfo) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

type ClientHandshakeInfo struct {
	Attributes *attributes.Attributes
}

func ClientHandshakeInfoFromContext(ctx context.Context) ClientHandshakeInfo {
	_ = "STUB: not implemented"
	return *new(ClientHandshakeInfo)
}

func CheckSecurityLevel(ai AuthInfo, level SecurityLevel) error {
	_ = "STUB: not implemented"
	return nil
}

type ChannelzSecurityInfo interface {
	GetSecurityValue() ChannelzSecurityValue
}

type ChannelzSecurityValue interface {
	isChannelzSecurityValue()
}

type OtherChannelzSecurityValue struct {
	ChannelzSecurityValue
	Name  string
	Value proto.Message
}
