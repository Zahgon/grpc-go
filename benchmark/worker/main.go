package main

import (
	"context"
	"flag"
	"net"
	"net/http"
	_ "net/http/pprof"
	"runtime"
	"strconv"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"

	testgrpc "google.golang.org/grpc/interop/grpc_testing"
	testpb "google.golang.org/grpc/interop/grpc_testing"
)

var (
	driverPort    = flag.Int("driver_port", 10000, "port for communication with driver")
	serverPort    = flag.Int("server_port", 0, "port for benchmark server if not specified by server config message")
	pprofPort     = flag.Int("pprof_port", -1, "Port for pprof debug server to listen on. Pprof server doesn't start if unset")
	blockProfRate = flag.Int("block_prof_rate", 0, "fraction of goroutine blocking events to report in blocking profile")

	logger = grpclog.Component("benchmark")
)

type byteBufCodec struct {
}

func (byteBufCodec) Marshal(v any) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (byteBufCodec) Unmarshal(data []byte, v any) error { _ = "STUB: not implemented"; return nil }

func (byteBufCodec) String() string { _ = "STUB: not implemented"; return "" }

type workerServer struct {
	testgrpc.UnimplementedWorkerServiceServer
	stop       chan<- bool
	serverPort int
}

func (s *workerServer) RunServer(stream testgrpc.WorkerService_RunServerServer) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *workerServer) RunClient(stream testgrpc.WorkerService_RunClientServer) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *workerServer) CoreCount(context.Context, *testpb.CoreRequest) (*testpb.CoreResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *workerServer) QuitWorker(context.Context, *testpb.Void) (*testpb.Void, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func main() {
	grpc.EnableTracing = false

	flag.Parse()
	lis, err := net.Listen("tcp", ":"+strconv.Itoa(*driverPort))
	if err != nil {
		logger.Fatalf("failed to listen: %v", err)
	}
	logger.Infof("worker listening at port %v", *driverPort)

	s := grpc.NewServer()
	stop := make(chan bool)
	testgrpc.RegisterWorkerServiceServer(s, &workerServer{
		stop:       stop,
		serverPort: *serverPort,
	})

	go func() {
		<-stop

		time.Sleep(time.Second)
		s.Stop()
	}()

	runtime.SetBlockProfileRate(*blockProfRate)

	if *pprofPort >= 0 {
		go func() {
			logger.Infoln("Starting pprof server on port " + strconv.Itoa(*pprofPort))
			logger.Infoln(http.ListenAndServe("localhost:"+strconv.Itoa(*pprofPort), nil))
		}()
	}

	s.Serve(lis)
}
