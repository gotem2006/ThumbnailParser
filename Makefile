.PHONY:
generate:
	protoc --go_out . \
	--go-grpc_out pkg/thumbnail --go-grpc_opt paths=source_relative \
	--grpc-gateway_out  pkg/thumbnail --grpc-gateway_opt paths=source_relative \
	proto/thumbnail.proto 
	mv pkg/thumbnail/proto/thumbnail.pb.gw.go pkg/thumbnail 
	mv pkg/thumbnail/proto/thumbnail_grpc.pb.go pkg/thumbnail




.PHONY:
run:
	go install client/main.go
	go run cmd/app/main.go


