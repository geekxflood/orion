FROM golang:1.22-alpine AS builder

# Build dependencies
WORKDIR /go/src/
COPY . .
RUN apk update && apk add make git
RUN go build -o build/orion .

# Second stage
FROM alpine:3

COPY --from=builder /go/src/build/orion /usr/local/bin/orion
CMD ["/usr/local/bin/orion"]
EXPOSE 9981
