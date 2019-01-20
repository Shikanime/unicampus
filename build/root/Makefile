build-admission:
	bazel build //cmd/admission:admission

build: build-admission

gen-grpc:
	protoc -I/usr/local/include -I. \
		-I$(GOPATH)/src \
		-I$(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--gofast_out=plugins=grpc:. \
		./api/admission/v1alpha1/generated.proto
	protoc -I/usr/local/include -I. \
		-I$(GOPATH)/src \
		-I$(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--gofast_out=plugins=grpc:. \
		./api/account/v1alpha1/generated.proto

gen: gen-grpc

# TODO: Move do deployments
install:
	helm init \
		--service-account tiller

deploy:
	helm install ./deployments/helm/unicampus \
		--name unicampus

clean:
	helm delete \
		--purge unicampus

.PHONY: gen-grpc gen build-admission build
