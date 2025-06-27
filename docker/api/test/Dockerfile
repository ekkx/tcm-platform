FROM golang:1.24-bullseye

WORKDIR /api

RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

CMD ["sleep", "infinity"]
