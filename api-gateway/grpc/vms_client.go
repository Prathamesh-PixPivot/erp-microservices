package grpc

import (
	"api-gateway/grpc/vms_pb"
	"log"

	"google.golang.org/grpc"
)

var VMSClient vms_pb.VendorServiceClient

func InitVMSClient() {
	conn, err := grpc.Dial("localhost:50058", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to VMS gRPC service: %v", err)
	}
	VMSClient = vms_pb.NewVendorServiceClient(conn)
}
