#!/bin/sh

cur_dir=`pwd`
bin_dir=${cur_dir}/bin
log_dir=${bin_dir}/log
pid_dir=${bin_dir}/pids
# 启动顺序无辜，不过还是建议rpc 服务先启，再起api 服务
app_dir="
    chat
    pushjob
    users
    apichat
    apiuser
    wspush
"

#function check() {
  # mac 需要额外安装 pidof 命令
#  if [[ $OSTYPE == 'darwin'* ]]; then
#    echo 'warn: macos need run `brew install pidof` install pidof command'
#  fi
#}

function stop_all(){
  echo "stop all"
  cd ${pid_dir}

  for app in ${app_dir}; do
    pid_file=${app}.pid
    pid=$(cat ${pid_file})
    pid_exist=$(ps aux | awk '{print $2}'| grep -w ${pid})

    if [ ${pid_exist} ];then
      kill ${pid}
      echo "kill ${app} ${pid}"
    fi
  done
}

function check_status(){
  echo 'check...'
  sleep 2
  cd ${pid_dir}

  for app in ${app_dir};do
    pid_file=${app}.pid
    pid=$(cat ${pid_file})
    pid_exist=$(ps aux | awk '{print $2}'| grep -w ${pid})

    if [ ! ${pid_exist} ];then
      echo ${app} ${pid} failed
    else
      echo ${app} ${pid} ok
    fi
  done
}

function start_all(){
  if [[ ! -d ${log_dir} ]];then
    mkdir -p ${log_dir}
  fi
  if [[ ! -d ${pid_dir} ]];then
    mkdir -p ${pid_dir}
  fi

  echo "server start..."
  cd ${bin_dir}

  for app in ${app_dir}; do
    config=config.${app}.yaml

    # 输出到日志
    # nohup 后台运行
    # >> 追加
    # 2>&1 把stderr重定向到stdout
    # & 后台运行nohup命令
    nohup ./${app} -conf=${config} >> ${log_dir}/${app}.log 2>&1 &

    # 记录pid
    # $!：Shell最后运行的后台Process的PID
    pid=$!
    #pid=$(pidof ${app})
    echo ${pid} > ${pid_dir}/${app}.pid

    echo "start ${app} ok"
  done
}

case $1 in
  stop)
    stop_all
    ;;
  start)
    stop_all
    start_all
    check_status
    ;;
  *)
    echo "restart.sh start|stop"
    echo "\t start: start all service"
    echo "\t stop: stop all service"
    ;;
esac

