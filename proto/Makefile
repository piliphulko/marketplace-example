protoc-all: protoc-basic protoc-service-acct-auth protoc-service-data-warehouse

protoc-basic:

	protoc -I . basic.proto --go_out=../../../.. --go-grpc_out=../../../..

protoc-service-acct-auth:

	protoc -I . service-acct-auth.proto --go_out=../../../.. --go-grpc_out=../../../..

protoc-service-data-warehouse:

	protoc -I . service-data-warehouse.proto --go_out=../../../.. --go-grpc_out=../../../..