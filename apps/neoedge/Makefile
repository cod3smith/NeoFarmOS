PROTO_DIR=internal/api/proto
PROTO_FILES=$(wildcard $(PROTO_DIR)/*.proto)

.PHONY: proto
proto:
	protoc --go_out=$(PROTO_DIR) --go-grpc_out=$(PROTO_DIR) $(PROTO_FILES) 