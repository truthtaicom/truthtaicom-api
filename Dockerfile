# Stage 1 (to create a "build" image, ~850MB)

FROM golang:1.11.0 AS builder
RUN go version

COPY . /go/src/tt/
WORKDIR /go/src/tt/

RUN set -x && \
    go get github.com/golang/dep/cmd/dep && \
    dep ensure -v

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o main .

# Stage 2 (to create a downsized "container executable", ~7MB)
FROM alpine:3.8
RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /go/src/tt/main .

EXPOSE 3000
ENTRYPOINT ["./main"]
