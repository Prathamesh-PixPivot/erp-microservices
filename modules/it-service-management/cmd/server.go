package main

import (
	"log"
	"net"

	pb "itsm/internal/transport/grpc/proto"

	"google.golang.org/grpc"
)

func StartITSMServer(address string, itsmServer pb.ITSMServiceServer) error {
	lis, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}
	grpcServer := grpc.NewServer()
	pb.RegisterITSMServiceServer(grpcServer, itsmServer)
	log.Printf("ITSM gRPC server listening on %s", address)
	return grpcServer.Serve(lis)
}
