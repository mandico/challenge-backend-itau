FROM golang:1.22.2-alpine AS builder
WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o challenge cmd/main.go

FROM alpine:latest
WORKDIR /apps

COPY --from=builder /build/challenge .
EXPOSE 8888

USER 1000
CMD ["./challenge"]