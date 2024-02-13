protoc-all: protoc-basic protoc-service-acct-auth protoc-service-order-pay protoc-service-data-warehouse protoc-service-data-customer

protoc-basic:

	protoc -I . basic.proto --go_out=../../../.. --go-grpc_out=../../../..

protoc-service-acct-auth:

	protoc -I . service-acct-auth.proto --go_out=../../../.. --go-grpc_out=../../../..

protoc-service-order-pay:

	protoc -I . service-order-pay.proto --go_out=../../../.. --go-grpc_out=../../../..

protoc-service-data-warehouse:

	protoc -I . service-data-warehouse.proto --go_out=../../../.. --go-grpc_out=../../../..

protoc-service-data-customer:

	protoc -I . service-data-customer.proto --go_out=../../../.. --go-grpc_out=../../../..
