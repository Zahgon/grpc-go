package main

import (
	"flag"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/examples/data"
	pb "google.golang.org/grpc/examples/route_guide/routeguide"
)

var (
	tls                = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	caFile             = flag.String("ca_file", "", "The file containing the CA root cert file")
	serverAddr         = flag.String("addr", "localhost:50051", "The server address in the format of host:port")
	serverHostOverride = flag.String("server_host_override", "x.test.example.com", "The server name used to verify the hostname returned by the TLS handshake")
)

func printFeature(client pb.RouteGuideClient, point *pb.Point) { _ = "STUB: not implemented"; return }

func printFeatures(client pb.RouteGuideClient, rect *pb.Rectangle) {
	_ = "STUB: not implemented"
	return
}

func runRecordRoute(client pb.RouteGuideClient) { _ = "STUB: not implemented"; return }

func runRouteChat(client pb.RouteGuideClient) { _ = "STUB: not implemented"; return }

func randomPoint() *pb.Point { _ = "STUB: not implemented"; return nil }

func main() {
	flag.Parse()
	var opts []grpc.DialOption
	if *tls {
		if *caFile == "" {
			*caFile = data.Path("x509/ca_cert.pem")
		}
		creds, err := credentials.NewClientTLSFromFile(*caFile, *serverHostOverride)
		if err != nil {
			log.Fatalf("Failed to create TLS credentials: %v", err)
		}
		opts = append(opts, grpc.WithTransportCredentials(creds))
	} else {
		opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	}

	conn, err := grpc.NewClient(*serverAddr, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewRouteGuideClient(conn)

	printFeature(client, &pb.Point{Latitude: 409146138, Longitude: -746188906})

	printFeature(client, &pb.Point{Latitude: 0, Longitude: 0})

	printFeatures(client, &pb.Rectangle{
		Lo: &pb.Point{Latitude: 400000000, Longitude: -750000000},
		Hi: &pb.Point{Latitude: 420000000, Longitude: -730000000},
	})

	runRecordRoute(client)

	runRouteChat(client)
}
