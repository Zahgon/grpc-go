package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"time"

	pb "google.golang.org/grpc/examples/features/proto/echo"

	"google.golang.org/grpc/credentials/tls/certprovider"
	"google.golang.org/grpc/security/advancedtls"
)

type server struct {
	pb.UnimplementedEchoServer
	name string
}

const credRefreshInterval = 1 * time.Minute
const goodServerWithCRLPort int = 50051
const revokedServerWithCRLPort int = 50053
const insecurePort int = 50054

func (s *server) UnaryEcho(_ context.Context, req *pb.EchoRequest) (*pb.EchoResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func insecureServer() { _ = "STUB: not implemented"; return }

func createAndRunInsecureServer(port int) { _ = "STUB: not implemented"; return }

func createAndRunTLSServer(credsDirectory string, useRevokedCert bool, port int) {
	_ = "STUB: not implemented"
	return
}

func makeRootProvider(credsDirectory string) certprovider.Provider {
	_ = "STUB: not implemented"
	return *new(certprovider.Provider)
}

func makeIdentityProvider(useRevokedCert bool, credsDirectory string) certprovider.Provider {
	_ = "STUB: not implemented"
	return *new(certprovider.Provider)
}

func makeCRLProvider(crlDirectory string) *advancedtls.FileWatcherCRLProvider {
	_ = "STUB: not implemented"
	return nil
}

func main() {
	credentialsDirectory := flag.String("credentials_directory", "", "Path to the creds directory of this repo")
	flag.Parse()
	if *credentialsDirectory == "" {
		fmt.Println("Must set credentials_directory argument")
		os.Exit(1)
	}
	go createAndRunTLSServer(*credentialsDirectory, false, goodServerWithCRLPort)
	go createAndRunTLSServer(*credentialsDirectory, true, revokedServerWithCRLPort)
	insecureServer()
}
