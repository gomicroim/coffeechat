# 通过传参数生成 Docker Image, SEE: https://segmentfault.com/q/1010000041543680

######## build stage ########
FROM golang:1.18 AS builder

# 程序目录
# 通过 Docker Compose args 动态传参
ARG app

COPY . /src
WORKDIR /src

# build，并且把所有程序都统一命名为app 方便下面的启动
RUN cd rpc/$app && GOPROXY=https://goproxy.cn make build \
    && cd bin && mv $app app

FROM debian:stable-20220912-slim
ARG app
ENV app $app

# 为了加快速度，不在进行额外的依赖安装
#RUN apt-get update && apt-get install -y --no-install-recommends \
#    ca-certificates  \
#    netbase \
#    && rm -rf /var/lib/apt/lists/ \
#    && apt-get autoremove -y && apt-get autoclean -y

COPY --from=builder /src/rpc/$app/bin /app

WORKDIR /app

EXPOSE 9000
VOLUME /data/conf

# ${app} 是环境变量，dockerfile 中 args 只在build time生效，CMD是 runtime
# see: https://github.com/moby/moby/issues/34772
CMD exec ./app -conf /data/conf/config.${app}.yaml