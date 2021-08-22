FROM golang:1.16 as builder

ARG GIT_COMMIT
ARG VERSION

WORKDIR /workspace
# Copy the Go Modules manifests
COPY go.mod go.mod
COPY go.sum go.sum

# Copy the go controller-manager go source
COPY main.go main.go
COPY cmd/ cmd/
COPY pkg/ pkg/

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on go build  --ldflags "-s -w \
       -X github.com/viveksyngh/kube-sync/pkg/version.GitCommit=${GIT_COMMIT} \
       -X github.com/viveksyngh/kube-sync/pkg/version.Version=${VERSION}" \
        -a -installsuffix cgo -o kube-sync

# Release stage
FROM alpine:3.13

RUN apk --no-cache add ca-certificates git

RUN addgroup -S app \
       && adduser -S -g app app \
       && apk add --no-cache ca-certificates

WORKDIR /home/app

COPY --from=builder  /workspace/kube-sync .

RUN chown -R app:app ./

ENV PATH=$PATH:/home/app/

USER app

CMD ["kube-sync"]