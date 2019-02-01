.PHONY: build-education
build-education:
	cd packages/go; bazel build //cmd/education:education_binary

.PHONY: clean-education
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
