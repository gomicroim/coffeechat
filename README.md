# gomicroim

opensource distributed microservice im server write by golang（开源分布式微服务IM服务端）.

特性：

- 分布式微服务架构。基于 kratos v2框架实现
- 同时支持读写扩散。针对 web 使用读扩散模式，针对 app 使用写扩散
- 离线消息同步使用 time-line模型，配合mongo实现高性能离线消息同步
- 在线消息漫游使用 mysql 存储，满足审计和任意时间回溯消息的需求
- BFF 层设计 + kong api网关设计，更贴合实战
- 简化私有协议设计，除了消息推送(websocket协议)之外，发消息、拉聊天列表等等都使用HTTP+JSON方式实现，简化私有协议开发成本
- websocket gateway 网关之间通信使用 kafka assign模式，解决到吞吐量下的消息跨服路由性能问题和可用性问题
- 支持docker compose 和 k8s 部署
- 支持百万级并发用户在线。建议使用 k8s 部署，配合动态扩容，实现高峰期的资源自适应分配和调整

## Client

- [Android](https://github.com/gomicroim/client-android): 基于android 7.0 + java实现，目前正在开发中（2022年8月），适用于 `V2版本` 服务端。

## Preview

to do.

## Features

See: [FEATURE.md](FEATURE.MD)

## Quick Start

> PS：请切换到**master**分支，编译和运行！

启动Server（要求安装docker desktop >= 4.0.1）:

```bash
$ git clone https://github.com/gomicroim/gomicroim.git
$ cd gomicroim
# 启动 redis kafka mongo mysql etcd等依赖
$ docker-compose up -d -f docker-compose.dep.yml
# 启动 kong api 网关相关依赖
$ docker-compose up -d -f docker-compose.kong.yml
# 启动 gomicroim 所有的服务
$ docker-compose up -d -f docker-compose.yml
```

停止：

```bash
$ cd gomicroim
$ docker-compose down -v
```

## Build

see: [build.md](build.md)

### Thinks

- [mattermost](https://github.com/mattermost/mattermost-server)：主要学习其go工程实践方面的一些技巧，目前还处在研究阶段。
- [Open-IM-Server](https://github.com/OpenIMSDK/Open-IM-Server)：通过分析它的架构和代码，理解了收件箱机制和im 微服务(go)的划分实践。
- [goim](https://github.com/Terry-Mao/goim)：学习到百万级架构下kafka是如何应用在聊天室场景的。
- [Terry-Ye/im](https://github.com/Terry-Ye/im)：结合goim，理解了所谓的job含义，看懂了goim的架构。
- [gim](https://github.com/alberliu/gim)：一个简单的写扩散项目，可以更深入理解写扩散的架构和原理。

更多开源im，请移步：[史上最全开源IM盘点](https://blog.csdn.net/xmcy001122/article/details/110679978)

## Contact

email：xmcy0011@sina.com  
微信交流：xuyc1992（请备注：im）  

喜欢的话，关注下公众号吧😊  
《Go和分布式IM》👈👈  
![qrcode](./docs/images/qrcode.png)

## LICENSE

gomicroim is provided under the [mit license](https://github.com/gomicroim/gomicroim/blob/master/LICENSE).