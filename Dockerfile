# 使用 Golang 官方镜像作为构建环境
FROM golang:1.19 AS builder

# 复制源代码到工作目录
COPY . /src
WORKDIR /src

# 设置 Go 代理并构建应用。你可以根据自己的构建命令进行相应的替换。
# 此处示例使用 GOPROXY 环境变量设置为国内代理，并使用 make 命令进行应用构建。
# RUN GOPROXY=https://proxy.golang.org,direct make build
RUN GOPROXY=https://goproxy.cn,direct make build

# 使用 Debian slim 镜像作为最终镜像
FROM debian:stable-slim

# 安装必要依赖
RUN apt-get update && apt-get install -y --no-install-recommends \
        ca-certificates \
        netbase && \
    rm -rf /var/lib/apt/lists/* && \
    apt-get clean && \
    apt-get autoremove -y

# 从构建器阶段复制构建产物到最终镜像的工作目录
COPY --from=builder /src/bin /app
# 如果configs目录在/src下，记住一定要加上配置！！！
COPY --from=builder /src/configs /app/configs

# 指定工作目录
WORKDIR /app

# 暴露运行服务的端口
EXPOSE 8000
EXPOSE 9000

# 设置启动命令
CMD ["./server", "-conf", "/app/configs"]