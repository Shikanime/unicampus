SCHOOL_DIR := ./api/admission/v1alpha1
ACCOUNT_DIR := ./api/account/v1alpha1

gen-grpc:
	mkdir -p $(SCHOOL_DIR)
	protoc -I $(SCHOOL_DIR)/ \
		--gofast_out=plugins=grpc:$(SCHOOL_DIR) \
		$(SCHOOL_DIR)/generated.proto
	mkdir -p $(ACCOUNT_DIR)
	protoc -I $(ACCOUNT_DIR)/ \
		--gofast_out=plugins=grpc:$(ACCOUNT_DIR) \
		$(ACCOUNT_DIR)/generated.proto

gen: gen-grpc

all: gen

.PHONY: gen-grpc gen
