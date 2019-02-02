.PHONY: install
install:
	protoc \
		-I/usr/local/include \
		-I. \
		-I$(GOPATH)/src \
		-I$(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--gofast_out=plugins=grpc:. \
		./api/v1alpha1/education_search.proto

.PHONY: refresh-deps
refresh-deps:
	go mod tidy
	bazel run //:gazelle -- update-repos -from_file=go.mod

.PHONY: refresh-import
refresh-import:
	find . -name "./api/**/*.pb.*" -type f -delete
	bazel run //:gazelle

.PHONY: refresh
refresh: refresh-deps refresh-import

.PHONY: build-exe
build-exe:
	bazel build //cmd/service:binary

.PHONY: build-image
build-image:
	bazel build //cmd/service:image

.PHONY: clean-service
clean:
	bazel clean --expunge
