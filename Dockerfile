# 构建阶段
FROM golang:1.24 AS builder
WORKDIR /app
COPY . .
RUN go mod tidy
RUN go build -o main main.go

# 运行阶段
FROM debian:bullseye-slim
WORKDIR /app
COPY --from=builder /app/main .
EXPOSE 8080
CMD ["./main"]
