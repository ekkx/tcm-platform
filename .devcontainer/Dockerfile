FROM golang:1.24-bullseye

WORKDIR /tcmrsv-api

RUN curl https://get.volta.sh | bash

RUN go install github.com/air-verse/air@latest
RUN go install github.com/bufbuild/buf/cmd/buf@v1.54.0
RUN go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
RUN go install github.com/fullstorydev/grpcui/cmd/grpcui@latest

RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
