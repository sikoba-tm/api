FROM golang:1.19.0-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o sikoba-api

FROM alpine:latest
COPY --from=builder /app .
CMD [ "./sikoba-api" ]