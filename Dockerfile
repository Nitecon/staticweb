# Build the manager binary
FROM golang:alpine as builder

RUN apk --update add ca-certificates

WORKDIR /workspace
# Copy the Go Modules manifests
COPY . /workspace
# cache deps before building and copying source so that we don't need to re-download as much
# and so that source changes don't invalidate our downloaded layer
RUN go mod download

# Build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=auto go build -ldflags "-s -w" -a -o app main.go

FROM scratch

WORKDIR /
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /workspace/app .
COPY data /data
USER nobody

ENV BIND_PORT 80
ENV WEB_DIR /data

ENTRYPOINT ["/app"]
EXPOSE 80