.PHONY: build-app
build-app:
	bazel build //cmd/admission:admission

.PHONY: build-app
build-app:
	bazel clean --expunge

.PHONY: build-grpc
build-grpc:
	protoc \
		-I/usr/local/include \
		-I. \
		-I$(GOPATH)/src \
		-I$(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--gofast_out=plugins=grpc:. \
		./api/admission/v1alpha1/generated.proto
	protoc \
		-I/usr/local/include \
		-I. \
		-I$(GOPATH)/src \
		-I$(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--gofast_out=plugins=grpc:. \
		./api/account/v1alpha1/generated.proto

.PHONY: build-grpc
build-grpc:
	find . -name "*.pb.*" -type f -delete

.PHONY: install-tiller
install-tiller:
	helm init \
		--service-account tiller

.PHONY: clean-tiller
clean-tiller:
	helm reset \
		--force

.PHONY: install-istio
install-istio:
	helm init \
		--service-account tiller
  kubectl apply \
		-f https://github.com/istio/istio/install/kubernetes/helm/helm-service-account.yaml
  helm install https://github.com/istio/istio/install/kubernetes/helm/istio \
		--name istio \
		--namespace istio-system

.PHONY: clean-istio
clean-istio:
  kubectl delete \
		-f https://github.com/istio/istio/install/kubernetes/helm/helm-service-account.yaml
  kubectl \
		-n istio-system \
		delete job --all

.PHONY: install-app
install-app:
	helm install ./deployments/helm/unicampus \
		--name unicampus

.PHONY: clean-app
clean-app:
	helm delete \
		--purge unicampus
