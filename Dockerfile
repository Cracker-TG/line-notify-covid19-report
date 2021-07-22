# build stage
FROM golang:1.16.5-alpine3.14 AS builder
WORKDIR /app
COPY . .
RUN apk add --no-cache git
RUN GOARCH=amd64 GOOS=linux go build -v -o go-covid19-line-notify .

# final stage
FROM alpine:latest
WORKDIR /root
RUN apk --no-cache add ca-certificates
COPY --from=builder /app/go-covid19-line-notify .
ENTRYPOINT ./go-covid19-line-notify