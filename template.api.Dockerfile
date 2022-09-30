FROM golang:1.18 AS builder
ARG app

COPY . /src
WORKDIR /src

RUN cd api/$app && GOPROXY=https://goproxy.cn make build && cd bin && mv $app app

FROM debian:stable-20220912-slim
ARG app
ENV app $app

RUN apt-get update && apt-get install -y --no-install-recommends \
    ca-certificates  \
    netbase \
    && rm -rf /var/lib/apt/lists/ \
    && apt-get autoremove -y && apt-get autoclean -y

COPY --from=builder /src/api/$app/bin /app

WORKDIR /app

EXPOSE 8000
VOLUME /data/conf

CMD exec ./app -conf /data/conf/config.${app}.yaml