FROM golang:1.21-alpine AS builder

ARG ARCH=amd64

ENV GOROOT /usr/local/go
ENV GOPATH /go
ENV PATH $GOPATH/bin:$GOROOT/bin:$PATH
ENV GO_VERSION 1.21
ENV GO111MODULE on
ENV CGO_ENABLED=0

# Build dependencies
WORKDIR /go/src/
COPY . .
RUN apk update && apk add make git
RUN go build -a -gcflags=all="-l -B" -ldflags="-w -s" -o build/orion .

# Second stage
FROM alpine:3.19

COPY --from=builder /go/src/build/orion /usr/local/bin/orion
CMD ["/usr/local/bin/orion"]
EXPOSE 9981
