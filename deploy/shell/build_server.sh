#!/bin/sh

root_dir=`cd ../../ && pwd`
src_services="
    api/apichat
    api/apiuser
    api/wspush
    rpc/chat
    rpc/pushjob
    rpc/user
"

function build(){
    echo "build start..."
    for server in ${src_services}; do
        echo "build ${server} ok"
        cd ${root_dir}
        cd ${server} && make build
    done
    echo "build end!"
}

build