PHONY: generate
generate:
	mkdir -p pkg/grpc/user_v1
	protoc --go_out=pkg/grpc/user_v1 --go_opt=paths=source_relative \
			--go-grpc_out=pkg/grpc/user_v1 --go-grpc_opt=paths=source_relative \
			api/grpc/user_v1/service.proto