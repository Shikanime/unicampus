PROTOBUF_DIR := ./api/protobuf-spec

ACCOUNT_DIR := ./pkg/unicampus_api_account_v1
SCHOOL_DIR := ./pkg/unicampus_api_admission_v1

gen-grpc:
	mkdir -p $(SCHOOL_DIR)
	protoc -I $(PROTOBUF_DIR)/ \
		--gofast_out=plugins=grpc:$(SCHOOL_DIR) \
		$(PROTOBUF_DIR)/unicampus_api_admission_v1.proto
	mkdir -p $(ACCOUNT_DIR)
	protoc -I $(PROTOBUF_DIR)/ \
		--gofast_out=plugins=grpc:$(ACCOUNT_DIR) \
		$(PROTOBUF_DIR)/unicampus_api_account_v1.proto

gen: gen-grpc

all: gen

.PHONY: gen-grpc gen
