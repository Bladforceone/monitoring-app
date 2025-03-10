FROM golang:1.24 AS builder

WORKDIR /app
COPY . .

RUN go mod tidy && go build -o monitoring-app cmd/main.go

FROM debian:latest
WORKDIR /app
COPY --from=builder /app/monitoring-app /app/monitoring-app

CMD ["/app/monitoring-app"]
