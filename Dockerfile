FROM golang:1.20-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod init tz-demo && \
    go mod tidy && \
    CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /go-bin/app main.go

FROM scratch
COPY --from=builder /go-bin/app /app
ENV TZ=UTC
ENTRYPOINT ["/app"]
