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

# Cài đặt các dependencies cần thiết (nếu có)
RUN apk --no-cache add ca-certificates bash

# Copy file thực thi từ builder stage sang runtime stage
COPY --from=builder /app/promotion_service /promotion_service
COPY config.yaml /config.yaml
# COPY wait-for-it.sh /wait-for-it.sh

# Đảm bảo file thực thi có quyền thực thi
RUN chmod +x /promotion_service
# RUN chmod +x /promotion_service /wait-for-it.sh

# Thiết lập entrypoint để chạy ứng dụng
# ENTRYPOINT ["/wait-for-it.sh", "db:3306", "--", "/promotion_service"]
ENTRYPOINT ["/promotion_service"]