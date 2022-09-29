#!/bin/sh

root_dir=`cd ../../ && pwd`
out_bin_dir=${root_dir}/bin/
src_services="
  api/apichat
  api/apiuser
  api/wspush
  rpc/chat
  rpc/pushjob
  rpc/users
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

function copyApp() {
  if [[  ! -d ${out_bin_dir} ]];then
    mkdir ${out_bin_dir}
  fi

  cd ${root_dir}
  for server in ${src_services}; do
    cp -R ${server}/bin/ ${out_bin_dir}
    echo "copy ${server}/bin/ => ${out_bin_dir} ok"
  done
}

function copyConfig() {
  cd ${root_dir}
  for server in ${src_services}; do
    # 取得app名
    fileName=config.${server##*/}.yaml
    src=${server}/configs/config.example.yaml
    dst=${out_bin_dir}/${fileName}

    cp ${src} ${dst}
    echo "copy  ${src} => ${dst} ok"
  done
}

case $1 in
  "all")
    build
    copyApp
    copyConfig
    ;;

  "app")
    build
    copyApp
    ;;

  *)
    echo "build.sh all|app"
    echo "\t all: build source and copy config"
    echo "\t app: only build source"
    ;;
esac

