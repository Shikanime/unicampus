all: refresh-deps refresh-app refresh-protocs

.PHONY: refresh-protocs
refresh-protocs:
	@protoc \
		-I/usr/local/include \
		-I. \
		-I$(GOPATH)/src \
		-I$(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--gofast_out=plugins=grpc:. \
		--js_out=import_style=commonjs:./website/src \
		--grpc-web_out=import_style=commonjs,mode=grpcwebtext:./website/src \
		./api/v1alpha1/*.proto

.PHONY: refresh-deps
refresh-deps:
	@go mod tidy
	bazel run //:gazelle -- update-repos -from_file=go.mod

.PHONY: refresh-app
refresh-app:
	@find . -name "./api/**/*.pb.*" -type f -delete
	bazel run //:gazelle

.PHONY: run-education
run-education:
	ibazel run //cmd/education:binary

.PHONY: build-education
build-education:
	bazel build //cmd/education:image

.PHONY: deploy-education
	helm install ./deployments/education --name unicampus-education
