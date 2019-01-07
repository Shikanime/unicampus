PROTOBUFS_DIR := ./protos

ACCOUNT_DIR := ./pkg/account
SCHOOL_DIR := ./pkg/unicampus_admission

gen-grpc:
	mkdir -p $(SCHOOL_DIR)
	protoc -I $(PROTOBUFS_DIR)/ \
		--gofast_out=plugins=grpc:$(SCHOOL_DIR) \
		$(PROTOBUFS_DIR)/admission.proto
	mkdir -p $(ACCOUNT_DIR)
	protoc -I $(PROTOBUFS_DIR)/ \
		--gofast_out=plugins=grpc:$(ACCOUNT_DIR) \
		$(PROTOBUFS_DIR)/account.proto

gen: gen-grpc

all: gen

.PHONY: gen-grpc gen
