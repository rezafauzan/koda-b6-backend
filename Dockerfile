FROM golang:1.25.0-alpine AS builder

WORKDIR /workspace

COPY . .

RUN go mod tidy

RUN go build -o Coffeeshop-Backend ./cmd/main.go

RUN chmod +x Coffeeshop-Backend

FROM alpine:latest

WORKDIR /app

COPY --from=builder /workspace/Coffeeshop-Backend .

EXPOSE 8888

ENTRYPOINT [ "/app/Coffeeshop-Backend" ]