proto:
	protoc --proto_path=. --go_out=./proto --go_opt=paths=source_relative  --go-grpc_out=./proto --go-grpc_opt=paths=source_relative team00_server.proto

clean_proto:
	rm ./proto/*.pb.go
.PHONY: proto