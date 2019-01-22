.PHONY: build-admission
build-admission:
	bazel build //cmd/admission:admission

.PHONY: clean-admission
clean:
	bazel clean --expunge

.PHONY: install
install:
	helm install \
		--name unicampus \
		./installments/helm/unicampus

.PHONY: uninstall
uninstall:
	helm delete \
		--purge unicampus
