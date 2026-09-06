package jwt

import (
	"errors"
	"time"
)

var (
	errTokenFileAccess = errors.New("token file access error")
	errJWTValidation   = errors.New("invalid JWT")
)

type jwtClaims struct {
	Exp int64 `json:"exp"`
}

type jwtFileReader struct {
	tokenFilePath string
}

func (r *jwtFileReader) readToken() (string, time.Time, error) {
	_ = "STUB: not implemented"
	return "", *new(time.Time), nil
}

const tokenDelim = "."

func extractClaimsRaw(s string) (string, bool) { _ = "STUB: not implemented"; return "", false }

func (r *jwtFileReader) extractExpiration(token string) (time.Time, error) {
	_ = "STUB: not implemented"
	return *new(time.Time), nil
}
