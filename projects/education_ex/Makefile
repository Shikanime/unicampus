.PHONY: install
install:
	protoc \
		-I/usr/local/include \
		-I. \
		-I$(GOPATH)/src \
		-I$(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--elixir_out=plugins=grpc:./lib/education \
		./api/v1alpha1/education.proto

.PHONY: build-exe
build-exe:
	mix release

.PHONY: build-image
build-image:
	mix release
