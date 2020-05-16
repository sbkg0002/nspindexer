# stage 0
FROM golang:latest as builder
WORKDIR /go/src/github.com/PierreZ/goStatic
COPY . .
RUN mkdir ./bin && \
    CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -tags netgo -installsuffix netgo -o ./bin/goStatic && \
    mkdir ./bin/etc && \
    ID=$(shuf -i 100-9999 -n 1) && \
    echo $ID && \
    echo "appuser:x:$ID:$ID::/sbin/nologin:/bin/false" > ./bin/etc/passwd && \
    echo "appgroup:x:$ID:appuser" > ./bin/etc/group

# stage 1
FROM scratch
WORKDIR /
COPY --from=builder /go/src/github.com/PierreZ/goStatic/bin/ .
USER appuser
ENTRYPOINT ["/goStatic"]
 