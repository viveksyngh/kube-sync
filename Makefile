
GIT_COMMIT=$(shell git rev-parse HEAD)
GIT_VERSION=$(shell git describe --tags 2>/dev/null || echo "$(GIT_COMMIT)")

.PHONY: build
build:
	CGO_ENABLED=0 GO111MODULE=on go build --ldflags "-s -w \
	   -X github.com/viveksyngh/kube-sync/pkg/version.GitCommit=${GIT_COMMIT} \
	   -X github.com/viveksyngh/kube-sync/pkg/version.Version=${GIT_VERSION}" \
	   -a -installsuffix cgo -o bin/kube-sync

.PHONY: build-linux-amd64
build-linux-amd64:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on go build --ldflags "-s -w \
	   -X github.com/viveksyngh/kube-sync/pkg/version.GitCommit=${GIT_COMMIT} \
	   -X github.com/viveksyngh/kube-sync/pkg/version.Version=${GIT_VERSION}" \
	   -a -installsuffix cgo -o bin/kube-sync-linux-amd64

.PHONY: build-linux-arm64
build-linux-arm64:
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 GO111MODULE=on go build --ldflags "-s -w \
	   -X github.com/viveksyngh/kube-sync/pkg/version.GitCommit=${GIT_COMMIT} \
	   -X github.com/viveksyngh/kube-sync/pkg/version.Version=${GIT_VERSION}" \
	   -a -installsuffix cgo -o bin/kube-sync-linux-arm64

.PHONY: build-windows-amd64
build-windows-amd64:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 GO111MODULE=on go build --ldflags "-s -w \
	   -X github.com/viveksyngh/kube-sync/pkg/version.GitCommit=${GIT_COMMIT} \
	   -X github.com/viveksyngh/kube-sync/pkg/version.Version=${GIT_VERSION}" \
	   -a -installsuffix cgo -o bin/kube-sync-windows-amd64

.PHONY: build-darwin-amd64
build-darwin-amd64:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 GO111MODULE=on go build --ldflags "-s -w \
	   -X github.com/viveksyngh/kube-sync/pkg/version.GitCommit=${GIT_COMMIT} \
	   -X github.com/viveksyngh/kube-sync/pkg/version.Version=${GIT_VERSION}" \
	   -a -installsuffix cgo -o bin/kube-sync-darwin-amd64

.PHONY: build-darwin-arm64
build-darwin-arm64:
	CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 GO111MODULE=on go build --ldflags "-s -w \
	   -X github.com/viveksyngh/kube-sync/pkg/version.GitCommit=${GIT_COMMIT} \
	   -X github.com/viveksyngh/kube-sync/pkg/version.Version=${GIT_VERSION}" \
	   -a -installsuffix cgo -o bin/kube-sync-darwin-arm64

.PHONY: build-all
build-all: build-linux-amd64 build-linux-arm64 build-windows-amd64 build-darwin-amd64 build-darwin-arm64

.PHONY: test-unit
test-unit:
	GO111MODULE=on go test ./... -cover