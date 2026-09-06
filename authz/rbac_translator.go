package authz

import (
	v3rbacpb "github.com/envoyproxy/go-control-plane/envoy/config/rbac/v3"
	v3routepb "github.com/envoyproxy/go-control-plane/envoy/config/route/v3"
	v3matcherpb "github.com/envoyproxy/go-control-plane/envoy/type/matcher/v3"
	"google.golang.org/protobuf/types/known/structpb"
)

const typeURLPrefix = "grpc.authz.audit_logging/"

type header struct {
	Key    string
	Values []string
}

type peer struct {
	Principals []string
}

type request struct {
	Paths   []string
	Headers []header
}

type rule struct {
	Name    string
	Source  peer
	Request request
}

type auditLogger struct {
	Name       string           `json:"name"`
	Config     *structpb.Struct `json:"config"`
	IsOptional bool             `json:"is_optional"`
}

type auditLoggingOptions struct {
	AuditCondition string         `json:"audit_condition"`
	AuditLoggers   []*auditLogger `json:"audit_loggers"`
}

type authorizationPolicy struct {
	Name                string
	DenyRules           []rule              `json:"deny_rules"`
	AllowRules          []rule              `json:"allow_rules"`
	AuditLoggingOptions auditLoggingOptions `json:"audit_logging_options"`
}

func principalOr(principals []*v3rbacpb.Principal) *v3rbacpb.Principal {
	_ = "STUB: not implemented"
	return nil
}

func permissionOr(permission []*v3rbacpb.Permission) *v3rbacpb.Permission {
	_ = "STUB: not implemented"
	return nil
}

func permissionAnd(permission []*v3rbacpb.Permission) *v3rbacpb.Permission {
	_ = "STUB: not implemented"
	return nil
}

func getStringMatcher(value string) *v3matcherpb.StringMatcher {
	_ = "STUB: not implemented"
	return nil
}

func getHeaderMatcher(key, value string) *v3routepb.HeaderMatcher {
	_ = "STUB: not implemented"
	return nil
}

func parsePrincipalNames(principalNames []string) []*v3rbacpb.Principal {
	_ = "STUB: not implemented"
	return nil
}

func parsePeer(source peer) *v3rbacpb.Principal { _ = "STUB: not implemented"; return nil }

func parsePaths(paths []string) []*v3rbacpb.Permission { _ = "STUB: not implemented"; return nil }

func parseHeaderValues(key string, values []string) []*v3rbacpb.Permission {
	_ = "STUB: not implemented"
	return nil
}

var unsupportedHeaders = map[string]bool{
	"host":                true,
	"connection":          true,
	"keep-alive":          true,
	"proxy-authenticate":  true,
	"proxy-authorization": true,
	"te":                  true,
	"trailer":             true,
	"transfer-encoding":   true,
	"upgrade":             true,
}

func unsupportedHeader(key string) bool { _ = "STUB: not implemented"; return false }

func parseHeaders(headers []header) ([]*v3rbacpb.Permission, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func parseRequest(request request) (*v3rbacpb.Permission, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func parseRules(rules []rule, prefixName string) (map[string]*v3rbacpb.Policy, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (options *auditLoggingOptions) toProtos() (allow *v3rbacpb.RBAC_AuditLoggingOptions, deny *v3rbacpb.RBAC_AuditLoggingOptions, err error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func toDenyCondition(condition v3rbacpb.RBAC_AuditLoggingOptions_AuditCondition) v3rbacpb.RBAC_AuditLoggingOptions_AuditCondition {
	_ = "STUB: not implemented"
	return *new(v3rbacpb.RBAC_AuditLoggingOptions_AuditCondition)
}

func translatePolicy(policyStr string) ([]*v3rbacpb.RBAC, string, error) {
	_ = "STUB: not implemented"
	return nil, "", nil
}
