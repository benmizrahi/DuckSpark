package worker

import (
	"fmt"
	"log"
	"net"

	grpc "google.golang.org/grpc"
)

func StartWorker(port int, master string) {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	RegisterWorkerAPIServer(s, &RouteWorkerAPIServer{})

	w := NewWorker(port, master)
	go w.registerToMaster()

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
