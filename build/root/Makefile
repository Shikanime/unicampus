SCHOOL_DIR := ./api/admission/v1alpha1
ACCOUNT_DIR := ./api/account/v1alpha1

build-admission:
	bazel build //cmd/admission:admission

build-admission-image:
	bazel build --platforms=@io_bazel_rules_go//go/toolchain:linux_amd64 //cmd/admission:admission_image

build: build-admission

gen-grpc:
	mkdir -p $(SCHOOL_DIR)
	protoc -I/usr/local/include -I. \
		-I$(GOPATH)/src \
		-I$(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--gofast_out=plugins=grpc:. \
		$(SCHOOL_DIR)/generated.proto
	mkdir -p $(ACCOUNT_DIR)
	protoc -I/usr/local/include -I. \
		-I$(GOPATH)/src \
		-I$(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--gofast_out=plugins=grpc:. \
		$(ACCOUNT_DIR)/generated.proto

gen: gen-grpc

all: gen

.PHONY: gen-grpc gen build-admission build
