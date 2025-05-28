FROM golang:1.23-bullseye

WORKDIR /api

RUN apt-get update && apt-get install -y \
  protobuf-compiler \
  && apt-get clean

RUN go install github.com/air-verse/air@latest
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

ENV PATH="$PATH:/go/bin"

CMD ["sh", "-c", "cd ./server && air -c .air.toml"]
