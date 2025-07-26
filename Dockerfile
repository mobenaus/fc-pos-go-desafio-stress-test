FROM golang:1.22 AS builder
WORKDIR /app
COPY . .
RUN GOOS=linux CGO_ENABLED=0 go build -o stress_test main.go

FROM alpine:latest
COPY --from=builder /app/stress_test .
ENTRYPOINT ["./stress_test"]