FROM --platform=$BUILDPLATFORM golang:1.24-alpine AS builder

RUN apk add --no-cache ca-certificates git

ARG TARGETARCH

WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN CGO_ENABLED=0 GOARCH=$TARGETARCH go build -o /bin/api ./cmd/api
RUN CGO_ENABLED=0 GOARCH=$TARGETARCH go build -o /bin/scheduler ./cmd/tcmscheduler

FROM alpine:3.21

RUN apk add --no-cache ca-certificates tzdata
ENV TZ=Asia/Tokyo

COPY --from=builder /bin/api /usr/local/bin/api
COPY --from=builder /bin/scheduler /usr/local/bin/scheduler
COPY migrations /migrations

ENTRYPOINT ["api"]
