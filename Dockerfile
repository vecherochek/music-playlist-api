FROM golang:1.23-alpine as builder

WORKDIR /app

RUN apk add --no-cache git make protobuf wget
RUN mkdir -p /usr/include/google/protobuf && \
    wget https://raw.githubusercontent.com/protocolbuffers/protobuf/main/src/google/protobuf/timestamp.proto -P /usr/include/google/protobuf && \
    wget https://raw.githubusercontent.com/protocolbuffers/protobuf/main/src/google/protobuf/duration.proto -P /usr/include/google/protobuf

COPY go.mod go.sum .env ./

RUN go mod download

COPY . .

RUN make install-deps
RUN make get-deps
RUN make generate

RUN go build -o playlist-service ./cmd/grpc_server

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/playlist-service .
COPY --from=builder /app/.env .

EXPOSE ${GRPC_PORT}

CMD ["./playlist-service"]
