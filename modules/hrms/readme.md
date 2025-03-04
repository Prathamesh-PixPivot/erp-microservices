
<!-- install -->
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest


<!-- gen *_pb.go from protoc -->
protoc --go_out=internal/transport/grpc/proto --go-grpc_out=internal/transport/grpc/proto --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative proto/*.proto


<!--  -->