FROM golang:1.18 AS builder

COPY . /src
WORKDIR /src

RUN cd rpc/chat && GOPROXY=https://goproxy.cn make build

FROM debian:stable-slim

RUN apt-get update && apt-get install -y --no-install-recommends \
    ca-certificates  \
    netbase \
    && rm -rf /var/lib/apt/lists/ \
    && apt-get autoremove -y && apt-get autoclean -y

COPY --from=builder /src/rpc/chat/bin /app

WORKDIR /app

EXPOSE 9000
VOLUME /data/conf

CMD ["./chat", "-conf", "/data/conf/config.chat.yaml"]