# kong 常用api

详见官网文档：[Services and Routes](https://docs.konghq.com/gateway/latest/get-started/services-and-routes/)

`!!!!WARN`: please use envoy api gateway.

## service

### 创建服务

```bash
curl -i -s -X POST http://localhost:8001/services \
  --data name=example_service \
  --data url='http://mockbin.org'
```

### 删除服务

```bash
curl -i -s -X DELETE http://localhost:8001/services/{service name or id}
```

### 查看服务详情

```bash
curl -X GET http://localhost:8001/services/example_service
```

### 更改服务设置

```bash
curl --request PATCH \
  --url localhost:8001/services/example_service \
  --data retries=6
```

### 查看所有服务列表

```bash
curl -X GET http://localhost:8001/services
```

## router

### 创建路由

```bash
curl -i -X POST http://localhost:8001/services/example_service/routes \
  --data 'paths[]=/mock' \
  --data name=example_route
```

此时，所有发送给 http://localhost:8001/mock 的请求都会被转发到example_service。

包括子请求，比如：

- /mock/test
- /mock/test/1

### 查看路由详情

```bash
curl -X GET http://localhost:8001/services/example_service/routes/example_route
```

可以通过 路由ID 或者 名称访问：

- /services/{service name or id}/routes/{route name or id}
- /routes/{route name or id}

### 更新路由

```bash
curl --request PATCH \
  --url localhost:8001/services/example_service/routes/example_route \
  --data tags="tutorial"
```

### 列出所有路由

```bash
curl http://localhost:8001/routes
```

### 删除路由

```bash
curl -i -s -X DELETE http://localhost:8001/routes/{route name or id}
```

## 请求代理

此时，kong 就代理了 `https://mockbin.org/` ， 它的本地访问地址是：`http://localhost:8000/mock`。
