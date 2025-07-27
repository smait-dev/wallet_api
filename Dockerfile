FROM golang:1.23-alpine AS build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main ./cmd/server/main.go

FROM alpine:latest
WORKDIR /app
COPY --from=build /app/main .
COPY configs ./configs
COPY migrations ./migrations
COPY db_seeds ./db_seeds
CMD ["./main"]