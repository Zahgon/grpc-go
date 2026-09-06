package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/tls/certprovider"
	"google.golang.org/grpc/security/advancedtls"
)

const credRefreshInterval = 1 * time.Minute
const serverAddr = "localhost"
const goodServerPort string = "50051"
const revokedServerPort string = "50053"
const insecurePort string = "50054"
const message string = "Hello"

func makeRootProvider(credsDirectory string) certprovider.Provider {
	_ = "STUB: not implemented"
	return *new(certprovider.Provider)
}

func makeIdentityProvider(revoked bool, credsDirectory string) certprovider.Provider {
	_ = "STUB: not implemented"
	return *new(certprovider.Provider)
}

func runClientWithProviders(rootProvider certprovider.Provider, identityProvider certprovider.Provider, crlProvider advancedtls.CRLProvider, port string, shouldFail bool) {
	_ = "STUB: not implemented"
	return
}

func tlsWithCRLsToGoodServer(credsDirectory string) { _ = "STUB: not implemented"; return }

func tlsWithCRLsToRevokedServer(credsDirectory string) { _ = "STUB: not implemented"; return }

func tlsWithCRLs(credsDirectory string) { _ = "STUB: not implemented"; return }

func makeCRLProvider(crlDirectory string) *advancedtls.FileWatcherCRLProvider {
	_ = "STUB: not implemented"
	return nil
}

func customVerificationSucceed(info *advancedtls.HandshakeVerificationInfo) (*advancedtls.PostHandshakeVerificationResults, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func customVerificationFail(info *advancedtls.HandshakeVerificationInfo) (*advancedtls.PostHandshakeVerificationResults, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func customVerification(credsDirectory string) { _ = "STUB: not implemented"; return }

func runClientWithCustomVerification(credsDirectory string, port string) {
	_ = "STUB: not implemented"
	return
}

func credentialsNewTLSExample(credsDirectory string) { _ = "STUB: not implemented"; return }

func insecureCredentialsExample() { _ = "STUB: not implemented"; return }

func runWithCredentials(creds credentials.TransportCredentials, fullServerAddr string, shouldSucceed bool) {
	_ = "STUB: not implemented"
	return
}

func main() {
	credsDirectory := flag.String("credentials_directory", "", "Path to the creds directory of this example repo")
	flag.Parse()

	if *credsDirectory == "" {
		fmt.Println("Must set credentials_directory argument to this repo's creds directory")
		os.Exit(1)
	}
	tlsWithCRLs(*credsDirectory)
	customVerification(*credsDirectory)
	credentialsNewTLSExample(*credsDirectory)
	insecureCredentialsExample()
}
