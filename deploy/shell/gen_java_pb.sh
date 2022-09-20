#!/bin/sh

CUR_DIR=`cd ../../ && pwd`
API_USER_DIR=${CUR_DIR}/api/apiuser/api/user/v1/
API_CHAT_DIR=${CUR_DIR}/api/apichat/api/chat/v1/
API_CHAT_CONSTANT_DIR=${CUR_DIR}/rpc/chat/api/chat/constant.proto
THIRD_PARTY_DIR=${CUR_DIR}/third_party

PROTOC_DIRS="
${API_USER_DIR}
${API_CHAT_DIR}
${API_CHAT_CONSTANT_DIR}
"

echo "generate start..."
for dir in ${PROTOC_DIRS}; do
    if [[ ${dir} == */ ]];then
      for file in ${dir}*.proto; do
          # go
          # protoc --proto_path=${dir} --proto_path=${CUR_DIR} --proto_path=${THIRD_PARTY_DIR} --go-grpc_out=${CUR_DIR} --go_out=${CUR_DIR} ${file}
          if [[ ! ${file} =~ "api_" ]]; then
            # java
            protoc --proto_path=${dir} --proto_path=${CUR_DIR} --proto_path=${THIRD_PARTY_DIR} --java_out=${CUR_DIR} ${file}
            echo "gen ${file} ok"
          fi
      done
    else
      protoc --proto_path=${CUR_DIR} --proto_path=${THIRD_PARTY_DIR} --java_out=${CUR_DIR} ${dir}
      echo "gen ${dir} ok"
    fi
done
echo "generate end!"