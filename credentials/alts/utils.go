package alts

import (
	"context"

	"google.golang.org/grpc/peer"
)

func AuthInfoFromContext(ctx context.Context) (AuthInfo, error) {
	_ = "STUB: not implemented"
	return *new(AuthInfo), nil
}

func AuthInfoFromPeer(p *peer.Peer) (AuthInfo, error) {
	_ = "STUB: not implemented"
	return *new(AuthInfo), nil
}

func ClientAuthorizationCheck(ctx context.Context, expectedServiceAccounts []string) error {
	_ = "STUB: not implemented"
	return nil
}
