FROM golang:1.18 AS builder

COPY . /src
WORKDIR /src

RUN cd api/apichat && GOPROXY=https://goproxy.cn make build

FROM debian:stable-slim

RUN apt-get update && apt-get install -y --no-install-recommends \
    ca-certificates  \
    netbase \
    && rm -rf /var/lib/apt/lists/ \
    && apt-get autoremove -y && apt-get autoclean -y

COPY --from=builder /src/api/apichat/bin /app

WORKDIR /app

EXPOSE 8000
VOLUME /data/conf

CMD ["./apichat", "-conf", "/data/conf/config.apichat.yaml"]