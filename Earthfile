FROM golang:alpine

protogen:
	FROM +golang-base

    # copy the proto files to generate
    COPY --dir protos/ ./
    COPY buf.work.yaml buf.gen.yaml ./

    # generate the pbs
    RUN buf generate \
            --template buf.gen.yaml \
            --path protos/chief-of-state-protos/chief_of_state/v1

    # save artifact to
    SAVE ARTIFACT gen gen AS LOCAL gen

golang-base:
    WORKDIR /app

    # install gcc dependencies into alpine for CGO
    RUN apk add gcc musl-dev curl git openssh

    # install docker tools
    # https://docs.docker.com/engine/install/debian/
    RUN apk add --update --no-cache docker

    # install linter
    # binary will be $(go env GOPATH)/bin/golangci-lint
    RUN curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.42.0
    RUN ls -la $(which golangci-lint)

    # install the go generator plugins
    RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
    RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
    RUN export PATH="$PATH:$(go env GOPATH)/bin"

    # install vektra/mockery
    RUN go get github.com/vektra/mockery/v2/.../

    # install buf from source
    RUN GO111MODULE=on GOBIN=/usr/local/bin go get \
        github.com/bufbuild/buf/cmd/buf \
        github.com/bufbuild/buf/cmd/protoc-gen-buf-breaking \
        github.com/bufbuild/buf/cmd/protoc-gen-buf-lint

    # install linter
    # binary will be $(go env GOPATH)/bin/golangci-lint
    RUN curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.42.1
