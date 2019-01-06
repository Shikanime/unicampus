PROTOBUFS_DIR := ./protos

SCHOOL_DIR := ./pkg/school

gen-grpc:
	mkdir -p $(SCHOOL_DIR)
	protoc -I $(PROTOBUFS_DIR)/ \
		--go_out=plugins=grpc:$(SCHOOL_DIR) \
		$(PROTOBUFS_DIR)/school.proto

gen: gen-grpc

all: gen

.PHONY: gen-grpc gen
