package advancedtls

import (
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/asn1"

	"google.golang.org/grpc/grpclog"
)

var grpclogLogger = grpclog.Component("advancedtls")

type RevocationOptions struct {
	DenyUndetermined bool

	CRLProvider CRLProvider
}

type revocationStatus int

const (
	RevocationUndetermined revocationStatus = iota

	RevocationUnrevoked

	RevocationRevoked
)

type CRL struct {
	certList *x509.RevocationList

	authorityKeyID []byte
	rawIssuer      []byte
}

func NewCRL(b []byte) (*CRL, error) { _ = "STUB: not implemented"; return nil, nil }

func ReadCRLFile(path string) (*CRL, error) { _ = "STUB: not implemented"; return nil, nil }

const tagDirectoryName = 4

var (
	oidDeltaCRLIndicator = asn1.ObjectIdentifier{2, 5, 29, 27}

	oidIssuingDistributionPoint = asn1.ObjectIdentifier{2, 5, 29, 28}

	oidCertificateIssuer = asn1.ObjectIdentifier{2, 5, 29, 29}

	oidAuthorityKeyIdentifier = asn1.ObjectIdentifier{2, 5, 29, 35}
)

func checkChainRevocation(verifiedChains [][]*x509.Certificate, cfg RevocationOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func checkChain(chain []*x509.Certificate, cfg RevocationOptions) revocationStatus {
	_ = "STUB: not implemented"
	return *new(revocationStatus)
}

func fetchCRL(c *x509.Certificate, crlVerifyCrt []*x509.Certificate, cfg RevocationOptions) (*CRL, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func checkCert(c *x509.Certificate, crlVerifyCrt []*x509.Certificate, cfg RevocationOptions) revocationStatus {
	_ = "STUB: not implemented"
	return *new(revocationStatus)
}

func checkCertRevocation(c *x509.Certificate, crl *CRL) (revocationStatus, error) {
	_ = "STUB: not implemented"
	return *new(revocationStatus), nil
}

func parseCertIssuerExt(ext pkix.Extension) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type authKeyID struct {
	ID []byte `asn1:"optional,tag:0"`
}

type issuingDistributionPoint struct {
	DistributionPoint          asn1.RawValue  `asn1:"optional,tag:0"`
	OnlyContainsUserCerts      bool           `asn1:"optional,tag:1"`
	OnlyContainsCACerts        bool           `asn1:"optional,tag:2"`
	OnlySomeReasons            asn1.BitString `asn1:"optional,tag:3"`
	IndirectCRL                bool           `asn1:"optional,tag:4"`
	OnlyContainsAttributeCerts bool           `asn1:"optional,tag:5"`
}

func parseCRLExtensions(c *x509.RevocationList) (*CRL, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func verifyCRL(crl *CRL, chain []*x509.Certificate) error { _ = "STUB: not implemented"; return nil }

const pemType string = "X509 CRL"

var crlPemPrefix = []byte("-----BEGIN X509 CRL")

func crlPemToDer(crlBytes []byte) []byte { _ = "STUB: not implemented"; return nil }

func extractCRLIssuer(crlBytes []byte) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func parseRevocationList(crlBytes []byte) (*x509.RevocationList, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
