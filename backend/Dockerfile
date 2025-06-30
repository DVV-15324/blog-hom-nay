FROM golang:1.23-alpine

# Cài thêm công cụ cần thiết (ví dụ: git để go mod dùng được)
RUN apk add --no-cache git

# Tạo thư mục làm việc
WORKDIR /app

# Copy module files trước để cache dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy toàn bộ source code
COPY . .

# Build binary
RUN go build -o app .

# Expose cổng dịch vụ
EXPOSE 3000

# Lệnh chạy container
CMD ["./app"]
