.PHONY: dep
dep:
	GO111MODULE=on go mod tidy
	GO111MODULE=on go mod vendor
	bazel run //:gazelle -- update-repos -from_file=go.mod

.PHONY: build
build:
	bazel build //server //client

.PHONY: test
test:
	bazel test

.PHONY: gazelle
gazelle:
	bazel run //:gazelle

.PHONY: buildifier
buildifier:
	bazel run //:buildifier

.PHONY: server
server:
	bazel run //server

.PHONY: client
client:
	bazel run //client

