.PHONY: build-admission
run-admission:
	bazel run //cmd/admission:admission

.PHONY: build-admission
build-admission:
	bazel build //cmd/admission:admission

.PHONY: clean-admission
clean-app:
	bazel clean --expunge

.PHONY: build-probobuf
build-probobuf:
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

.PHONY: clean-probobuf
clean-probobuf:
	find . -name "*.pb.*" -type f -delete

.PHONY: deploy-tiller
deploy-tiller:
	helm init \
		--service-account tiller

.PHONY: undeploy-tiller
undeploy-tiller:
	helm reset \
		--force

.PHONY: deploy-istio
deploy-istio:
	helm init \
		--service-account tiller
  kubectl apply \
		-f https://github.com/istio/istio/install/kubernetes/helm/helm-service-account.yaml
  helm install https://github.com/istio/istio/install/kubernetes/helm/istio \
		--name istio \
		--namespace istio-system

.PHONY: undeploy-istio
undeploy-istio:
	kubectl delete \
		-f https://github.com/istio/istio/install/kubernetes/helm/helm-service-account.yaml
	kubectl \
		-n istio-system \
		delete job --all

.PHONY: deploy-app
deploy-app:
	helm install ./deployments/helm/unicampus \
		--name unicampus

.PHONY: undeploy-app
undeploy-app:
	helm delete \
		--purge unicampus
