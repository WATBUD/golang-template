# 使用官方 Go 鏡像作為基礎鏡像
FROM golang:1.16

# 設置工作目錄
WORKDIR /app

# 複製 go.mod 和 go.sum 文件
COPY go.mod go.sum ./

# 下載所有依賴
RUN go mod download

# 複製源代碼
COPY . .

# 編譯 Go 應用程序
RUN go build -o main .

# 暴露服務器端口
EXPOSE 8080

# 運行可執行文件
CMD ["./main"]
