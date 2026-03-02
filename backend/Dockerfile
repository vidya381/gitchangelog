FROM golang:1.23-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o gitchangelog ./cmd/server

FROM scratch
COPY --from=builder /app/gitchangelog /gitchangelog
EXPOSE 8080
ENTRYPOINT ["/gitchangelog"]
