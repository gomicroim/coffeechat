#!/bin/sh

app_dir="
    api/apichat/bin/apichat
    api/apiuser/bin/apiuser
    api/wspush/bin/wspush
    rpc/chat/bin/chat
    rpc/pushjob/bin/pushjob
    rpc/user/bin/user
"

function start_all(){
    echo "server start..."
    for app in ${app_dir}; do
        cd ${root_dir}

        # 输出到日志
        nohup ${server}  \
        -data-dir ${DATA_DIR} \
        –listen-client-urls http://0.0.0.0:2379 \
        -advertise-client-urls http://127.0.0.1:2379 \
        > ${ETCD_BIN}/log.log 2>&1 &

            pid=$(pidof etcd)
            echo ${pid} > ${ETCD_BIN}/etcd.pid


        echo "start ${app} ok"
    done
    echo "server start end!"
}

start_all