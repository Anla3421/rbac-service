ARG GO_VERSION="1.23"
ARG ALPINE_VERSION="3.21"

#
## build stage
FROM golang:${GO_VERSION}-alpine${ALPINE_VERSION} AS builder

# 設置工作目錄
WORKDIR /app

RUN apk update  && \
    apk add --no-cache git gcc musl-dev
COPY go.mod go.sum ./
RUN go mod download

# 複製整個專案
COPY . .

# 編譯應用程序
RUN CGO_ENABLED=1 GOOS=linux go build -a -installsuffix cgo -o rbac-service .

#
## Release Stage
FROM alpine:${ALPINE_VERSION} 

# 安裝必要的系統依賴
RUN apk update  && \
    apk add --no-cache tzdata

# 設置工作目錄
WORKDIR /root/

# 從 builder 階段複製編譯好的二進制文件(可執行文件)
COPY --from=builder /app/rbac-service .

# 複製配置文件目錄
COPY --from=builder /app/configs ./configs

# 設置時區
ENV TZ=Asia/Taipei

# 暴露服務端口
EXPOSE 5002

# 設置健康檢查
HEALTHCHECK --interval=30s --timeout=10s --start-period=5s --retries=3 \
  CMD wget --quiet --tries=1 --spider http://localhost:5002/health || exit 1

# 運行應用程序
CMD ["./rbac-service"]