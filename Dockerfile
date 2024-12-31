# Sử dụng image Golang chính thức để build ứng dụng
FROM golang:1.22 AS builder

# Thiết lập thư mục làm việc trong container
WORKDIR /app

# Copy file module và cài đặt dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy toàn bộ mã nguồn vào container
COPY . .

# Build ứng dụng Go với các biến môi trường để đảm bảo build cho Linux
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o promotion_service ./cmd/main.go

# Sử dụng một image nhỏ hơn để chạy ứng dụng
FROM alpine:latest

# Cài đặt các dependencies cần thiết và cleanup cache
RUN apk --no-cache add \
    ca-certificates \
    bash \
    mysql-client \
    && rm -rf /var/cache/apk/*

WORKDIR /app

# Copy file thực thi từ builder stage sang runtime stage
COPY --from=builder /app/promotion_service /app/
COPY config.yaml /app/
COPY wait-for-it.sh /app/

# Đảm bảo file thực thi có quyền thực thi
RUN chmod +x /app/promotion_service /app/wait-for-it.sh

# Thiết lập healthcheck
HEALTHCHECK --interval=30s --timeout=3s --retries=3 \
    CMD mysqladmin ping -h ${DB_HOST} -u ${DB_USER} -p${DB_PASSWORD} || exit 1

# Thiết lập entrypoint để chạy ứng dụng
ENTRYPOINT ["/app/wait-for-it.sh", "db:3306", "--timeout=60", "--", "/app/promotion_service"]
