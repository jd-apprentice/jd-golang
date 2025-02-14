FROM golang:1.23-alpine AS base
WORKDIR /build
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN GOARCH=amd64 go build -o ./app_name ./src/main.go

FROM alpine:latest AS runner

RUN addgroup -S app_user && adduser -S app_user -G app_user

WORKDIR /app
COPY --from=base /build/app_name .
USER app_user:app_user
ENTRYPOINT ["./app_name"]