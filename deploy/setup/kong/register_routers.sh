#!/bin/sh

KONG_SRV_ADDR=http://localhost:8001/services
KONG_ROUTER_ADDR=http://localhost:8001/services

# 注册到kong服务的地址，可以改成本机IP，如果映射了 host.docker.internal ，则建议使用域名
SRV_API_HOST=http://host.docker.internal

# 列出所有服务
function list_srv(){
    curl -X GET ${KONG_SRV_ADDR}/services
}

# 列出所有路由
function list_routes(){
    curl -X GET ${KONG_SRV_ADDR}/routes
}

# 注册服务 $1：服务名 $2：端口
function register_srv(){
    srv_name=$1
    srv_port=$2
    srv_url=${SRV_API_HOST}':'${srv_port}

    # POST：仅创建
    # PUT：创建或者更新
    curl -i -s -X POST ${KONG_SRV_ADDR} \
        --data name=${srv_name} \
        --data url=${srv_url} > /dev/null 2>&1

    echo 'register service '${srv_name}' '${srv_url}    
}

# 注册路由 $1：服务名 $2：路径 $3：路由名称
function register_router(){
    srv_name=$1
    route_name=$2
    url_path=$3
    srv_url=${SRV_API_HOST}':'${srv_port}${url_path}

    # POST：仅创建
    # PUT：创建或者更新
    curl -i -X PUT ${KONG_SRV_ADDR}/${srv_name}/routes/${route_name} \
    -H 'Content-Type: application/json' \
    --data '{
        "name":"'${route_name}'",
        "paths": ["'${url_path}'"], 
        "methods":[
            "GET",
            "POST"
        ],
        "strip_path":false, 
        "headers": { "Content-Type":["application/json"]}
        }' > /dev/null 2>&1

    # --data 'paths[]='${url_path} \
    # --data 'strip_path=false' \
    # --data name=${route_name}
    # --data 'methods[]=GET&methods[]=POST' \

    echo 'register route '${srv_name}' '${route_name}' '${srv_url}
}

function api_user(){
    register_srv api_user 8100
    register_router api_user auth /auth
}

api_user
