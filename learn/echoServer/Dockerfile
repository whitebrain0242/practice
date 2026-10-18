FROM golang:1.25-alpine AS builder 
# 基础镜像,ap是极简linux,这个阶段叫builder
WORKDIR /app
COPY go.mod ./
COPY main.go ./
COPY main_test.go ./
RUN go build -o echoserver main.go # 编译阶段

# u运行阶段

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/echoserver ./ 
# 把 编译好的j二进制文件复制过来
EXPOSE 8000
CMD ["./echoserver"] 
# 容器启动时的命令


