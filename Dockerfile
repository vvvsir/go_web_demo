# 构建
FROM golang:1.16-alpine as builder
WORKDIR /app
ENV GOPROXY=https://goproxy.cn
COPY ./go/ ./
RUN go mod download && \
sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories && \
#安装tzdata安装包
apk add --no-cache tzdata
RUN CGO_ENABLED=0 go build -ldflags "-s -w" -o go_web_demo

# 打包
FROM alpine as runner
#设置时区
COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
COPY --from=builder /app/go_web_demo /app/
WORKDIR /app
CMD ["./go_web_demo"]
