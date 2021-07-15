WORKDIR=`pwd`
VersionBase=github.com/pubgo/lug
Domain=lugo
Version=$(shell git tag --sort=committerdate | tail -n 1)
BuildTime=$(shell date "+%F %T")
CommitID=$(shell git rev-parse --short=6 HEAD)
GOPATH=$(shell go env GOPATH )
TAG=$(shell git describe --abbrev=0 --tags)
LDFLAGS=-ldflags " \
-X '${VersionBase}/version.BuildTime=${BuildTime}' \
-X '${VersionBase}/version.CommitID=${CommitID}' \
-X '${VersionBase}/version.Version=${Version}' \
-X '${VersionBase}/version.Domain=${Domain}' \
-X '${VersionBase}/version.Data=hello' \
"

.PHONY: build
build:
	@go build ${LDFLAGS} -mod vendor -v -o main cmd/lug/main.go

build_hello_test:
	@go build ${LDFLAGS} -mod vendor -v -o main  example/hello/main.go

.PHONY: test./ma
test:
	@go test -short -race -v ./... -cover

ci:
	@golangci-lint run -v --timeout=5m


gen-proto:
	rm -rf example/proto/hello/*.go
	rm -rf example/proto/hello/*.json
	rm -rf example/proto/login/*.go
	rm -rf example/proto/login/*.json
	flerken protoc ls
	flerken protoc gen

proto:
	rm -rf example/proto/hello/*.go
	rm -rf example/proto/hello/*.json
	rm -rf example/proto/login/*.go
	rm -rf example/proto/login/*.json
	flerken protoc vendor-rm
	flerken protoc vendor
	flerken protoc ls
	flerken protoc gen


.PHONY: gen
gen:
	cd protoc-gen-lug && go install .

.PHONY: example
example:
	#go build ${LDFLAGS} -mod vendor -v -o main example/*.go
	CGO_CFLAGS=-Wno-undef-prefix go build ${LDFLAGS} -v -o main example/*.go

docker:
	docker build -t lug .

build-all:
	go build -tags "kcp quic" ./...

deps:
	go list -f '{{ join .Deps  "\n"}}' ./... |grep "/" | grep -v "github.com/smallnest/rpcx"| grep "\." | sort |uniq

cover:
	gocov test -tags "kcp quic" ./... | gocov-html > cover.html
	open cover.html

check-libs:
	GIT_TERMINAL_PROMPT=1 GO111MODULE=on go list -m -u all | column -t

update-libs:
	GIT_TERMINAL_PROMPT=1 GO111MODULE=on go get -u -v ./...

mod-tidy:
	GIT_TERMINAL_PROMPT=1 GO111MODULE=on go mod tidy

mac:
	GOOS=darwin go build -ldflags="-s -w" -ldflags="-X 'main.BuildTime=$(version)'" -o goctl-darwin goctl.go
	$(if $(shell command -v upx), upx goctl-darwin)
win:
	GOOS=windows go build -ldflags="-s -w" -ldflags="-X 'main.BuildTime=$(version)'" -o goctl.exe goctl.go
	$(if $(shell command -v upx), upx goctl.exe)
linux:
	GOOS=linux go build -ldflags="-s -w" -ldflags="-X 'main.BuildTime=$(version)'" -o goctl-linux goctl.go
	$(if $(shell command -v upx), upx goctl-linux)

changelog:
	docker run --rm \
		--interactive \
		--tty \
		-e "CHANGELOG_GITHUB_TOKEN=${CHANGELOG_GITHUB_TOKEN}" \
		-v "$(PWD):/usr/local/src/your-app" \
		ferrarimarco/github-changelog-generator:1.14.3 \
				-u grpc-ecosystem \
				-p grpc-gateway \
				--author \
				--compare-link \
				--github-site=https://github.com \
				--unreleased-label "**Next release**" \
				--release-branch=master \
				--future-release=v2.3.0

vet:
	@go vet ./...

generate:
	@go generate ./...

.PHONY: install
install:
	@go install -v github.com/tinylib/msgp
	@go install -v github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2