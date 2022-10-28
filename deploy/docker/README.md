# Docker Compose

## 全部跑在Docker中

无需特殊设置，直接从 compose 文件启动即可。

## 部分跑在Docker中

如果要Debug某个服务（宿主机Goland启动），依赖服务跑在Docker中，则需要配置特殊的 DNS 来完成Docker隔离网络到宿主机的相互通信。

参考这篇文章：
- [From inside of a Docker container, how do I connect to the localhost of the machine?](https://stackoverflow.com/questions/24319662/from-inside-of-a-docker-container-how-do-i-connect-to-the-localhost-of-the-mach)
- [Docker 多服务器网络配置（简单路由）](https://zhuanlan.zhihu.com/p/149830770)

### 宿主机增加DNS解析

```shell
$ vim /etc/hosts
```