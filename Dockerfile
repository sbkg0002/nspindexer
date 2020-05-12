
FROM golang:alpine AS builder
RUN apk update && apk add --no-cache git
WORKDIR $GOPATH/src/nspindexer/
COPY . .
RUN go build -o /go/bin/nspindexer main.go 


FROM scratch
COPY --from=builder /go/bin/nspindexer /go/bin/nspindexer
WORKDIR /nsp
ENTRYPOINT ["/go/bin/nspindexer"]