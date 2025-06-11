PROTO_DIRS := proto/user proto/task
PROTOS     := $(wildcard $(addsuffix /*.proto,$(PROTO_DIRS)))

generate:
	protoc \
		--go_out=. --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		$(PROTOS)

clean:
	find . -name "*.pb.go" -delete
